-- idempotent migration: search_filters table per app_user_id
CREATE TABLE IF NOT EXISTS search_filters (
    app_user_id   UUID        NOT NULL PRIMARY KEY,
    genders       TEXT[]      NOT NULL DEFAULT ARRAY['Все'],
    age_min       SMALLINT    NOT NULL DEFAULT 18,
    age_max       SMALLINT    NOT NULL DEFAULT 65,
    distance_km   SMALLINT    NOT NULL DEFAULT 50,
    height_min    SMALLINT,
    height_max    SMALLINT,
    religions     TEXT[]      NOT NULL DEFAULT '{}',
    wants_children BOOLEAN,
    values        TEXT[]      NOT NULL DEFAULT '{}',
    only_verified BOOLEAN     NOT NULL DEFAULT FALSE,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
