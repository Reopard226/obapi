package apiaccess

import (
	"time"

	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/api/apiaccess"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
	"oceanbolt.com/obapi/rpc/iam"
)

// New creates new apiaccess logging service
func New(svc apiaccess.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents apiaccess logging service
type LogService struct {
	apiaccess.Service
	logger model.Logger
}

const name = "apiaccess"

// ReqCreateKey is ReqCreateKey request params type
type ReqCreateKey struct {
	Tag string
	exp int64
}

// ReqDelete is ReqDelete request params type
type ReqDelete struct {
	ApikeyID string
}

// ListKey logging
func (ls *LogService) ListKey(c echo.Context) (resp *iam.UserKeys, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "ListKey request", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ListKey(c)
}

// CreateKey logging
func (ls *LogService) CreateKey(c echo.Context, tag string, exp int64) (resp *iam.UserKeyWithSecret, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "CreateKey request", err,
			map[string]interface{}{
				"req":  ReqCreateKey{tag, exp},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.CreateKey(c, tag, exp)
}

// DeleteKey logging
func (ls *LogService) DeleteKey(c echo.Context, apikeyID string) (resp *iam.KeyDeletedResponse, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "DeleteKey request", err,
			map[string]interface{}{
				"req":  ReqDelete{apikeyID},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.DeleteKey(c, apikeyID)
}
