package model

import (
	"errors"

	"github.com/labstack/echo"
)

var (
	// ErrGeneric is used for testing purposes and for errors handled later in the callstack
	ErrGeneric = errors.New("generic error")

	// ErrBadRequest (400) is returned for bad request (validation)
	ErrBadRequest = echo.NewHTTPError(400)

	// ErrUnauthorized (401) is returned when user is not authorized
	ErrUnauthorized = echo.ErrUnauthorized

	// ErrInvalidSegment is invalid segment error within request
	ErrInvalidSegment = "Error in query parameter: invalid value for parameter 'segment'"

	// ErrInvalidMetric is invalid metric error within request
	ErrInvalidMetric = "Error in query parameter: invalid value for parameter 'metric'"

	// ErrMissingRegionID is region_id missing error within request
	ErrMissingRegionID = "Error parameter 'region_id' is missing with no default"

	// ErrParseCreatKeyParams is parse error of create key request params
	ErrParseCreatKeyParams = "Could not parse request payload - 'tag' entry should be string, 'exp' entry should be int64 and value should be seconds since epoch"

	// ErrMissingTag is tag missig error within request
	ErrMissingTag = "Missing 'tag' entry in request payload"
)
