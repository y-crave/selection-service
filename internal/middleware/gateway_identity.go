package middleware

import (
	"encoding/base64"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"selection-service/internal/authidentity"
	"selection-service/internal/config"
)

const (
	headerUserID      = "X-User-Id"
	headerUserEmail   = "X-User-Email"
	headerDisplayName = "X-User-Name"
)

type jwtClaims struct {
	Sub               string `json:"sub"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
}

func claimsFromBearerJWTUnsafe(authzHeader string) (jwtClaims, error) {
	token := strings.TrimPrefix(strings.TrimSpace(authzHeader), "Bearer ")
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return jwtClaims{}, http.ErrNoCookie
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return jwtClaims{}, err
	}
	var c jwtClaims
	if err := json.Unmarshal(payload, &c); err != nil {
		return jwtClaims{}, err
	}
	return c, nil
}

// GatewayIdentity attaches trusted upstream identity from Istio-forwarded headers.
// JWT validation does not happen here — RequestAuthentication must run on the mesh ingress.
//
// Local development without Istio: set DEV_TRUST_HEADERS=true and send Authorization: Bearer <jwt>;
// the middleware extracts sub/email/name from the payload without verifying the signature (never enable in production).
func GatewayIdentity(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawID := strings.TrimSpace(r.Header.Get(headerUserID))
			email := strings.TrimSpace(r.Header.Get(headerUserEmail))
			display := strings.TrimSpace(r.Header.Get(headerDisplayName))

			if rawID == "" && cfg.DevTrustHeaders {
				authz := r.Header.Get("Authorization")
				if claims, err := claimsFromBearerJWTUnsafe(authz); err == nil && strings.TrimSpace(claims.Sub) != "" {
					rawID = strings.TrimSpace(claims.Sub)
					if email == "" {
						email = strings.TrimSpace(claims.Email)
					}
					if display == "" {
						display = strings.TrimSpace(claims.Name)
						if display == "" {
							display = strings.TrimSpace(claims.PreferredUsername)
						}
					}
				}
			}

			if rawID == "" {
				slog.Warn("gateway identity missing", "header", headerUserID)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
				return
			}

			userID, err := uuid.Parse(rawID)
			if err != nil {
				slog.Warn("gateway identity invalid user id")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
				return
			}

			ctx := authidentity.WithContext(r.Context(), authidentity.Identity{
				UserID:      userID,
				Email:       email,
				DisplayName: display,
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
