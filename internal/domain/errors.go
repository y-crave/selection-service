package domain

import "errors"

var (
	// Ошибки валидации
	ErrInvalidSexValue    = errors.New("invalid sex value")
	ErrInvalidAgeRange    = errors.New("age_from must be <= age_to")
	ErrAgeOutOfRange      = errors.New("age must be between min and max")
	ErrInvalidHeightRange = errors.New("height_from must be <= height_to")
	ErrHeightOutOfRange   = errors.New("height must be between min and max")

	// Ошибки парсинга / маппинга

	ErrInvalidFilterID     = errors.New("invalid filter ID")
	ErrInvalidUserID       = errors.New("invalid user ID")
	ErrInvalidSearchTypeID = errors.New("invalid search type ID")
	ErrInvalidUseTargetID  = errors.New("invalid use target ID")
	ErrInvalidTagID        = errors.New("invalid tag ID")

	// Общие ошибки
	ErrInternal            = errors.New("internal error")
	ErrFilterAlreadyExists = errors.New("filter already exists")
)
