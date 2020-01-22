package port

import (
	"time"

	"github.com/labstack/echo"
	"oceanbolt.com/obapi/internal/echoapi/api/port"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"
)

// New creates new port logging service
func New(svc port.Service, logger model.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents port logging service
type LogService struct {
	port.Service
	logger model.Logger
}

const name = "port"

// Req is request params type
type Req struct {
	PortID  string
	Segment string
}

// List logging
func (ls *LogService) List(c echo.Context, portID, segment string) (resp []model.Port, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List port request", err,
			map[string]interface{}{
				"req":  Req{portID, segment},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, portID, segment)
}

// View logging
func (ls *LogService) View(c echo.Context, portID, segment string) (resp *model.Port, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View port request", err,
			map[string]interface{}{
				"req":  Req{portID, segment},
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, portID, segment)
}
