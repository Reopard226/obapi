package transport

import (
	"net/http"
	"strings"

	"oceanbolt.com/obapi/internal/echoapi/api/tonnage"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc tonnage.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc tonnage.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/tonnage")

	ur.GET("/basin", h.basinBalance)
	ur.GET("/drydock", h.dryDockTimeseries)
	ur.GET("/drydock/summary", h.dryDockSummaryStats)
}

func (h *HTTP) basinBalance(c echo.Context) error {
	var Params struct {
		Basin   string `json:"basin"`
		Segment string `json:"segment"`
	}
	c.Bind(&Params)
	if !model.CheckValidSegment(Params.Segment) {
		return c.String(http.StatusBadRequest, model.ErrInvalidSegment)
	}

	s := strings.Title(strings.ToLower(Params.Segment))
	if s == "Total" {
		s = ""
	}
	result, err := h.svc.BasinBalance(c, strings.ToLower(Params.Basin), s)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) dryDockTimeseries(c echo.Context) error {
	var Params struct {
		Segment  string `json:"segment"`
		Metric   string `json:"metric"`
		Absolute bool   `json:"absolute"`
	}
	c.Bind(&Params)

	if !model.CheckValidSegment(Params.Segment) {
		return c.String(http.StatusBadRequest, model.ErrInvalidSegment)
	}
	if !model.CheckValidMetric(Params.Metric) {
		return c.String(http.StatusBadRequest, model.ErrInvalidMetric)
	}

	s := strings.Title(strings.ToLower(Params.Segment))
	if s == "Total" {
		s = ""
	}
	result, err := h.svc.DryDockTimeseries(c, s, Params.Metric, Params.Absolute)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) dryDockSummaryStats(c echo.Context) error {
	var Params struct {
		Segment string `json:"segment"`
	}
	c.Bind(&Params)
	if !model.CheckValidSegment(Params.Segment) {
		return c.String(http.StatusBadRequest, model.ErrInvalidSegment)
	}

	s := strings.Title(strings.ToLower(Params.Segment))
	if s == "Total" {
		s = ""
	}
	result, err := h.svc.DryDockSummaryStats(c, s)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
