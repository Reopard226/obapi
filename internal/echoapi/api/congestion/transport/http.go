package transport

import (
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"oceanbolt.com/obapi/internal/echoapi/api/congestion"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc congestion.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc congestion.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/congestion")

	ur.GET("/port", h.congestionPort)
	ur.GET("/region", h.congestionRegion)
	ur.GET("/port/list", h.listPort)
	ur.GET("/region/list", h.listRegion)
}

func (h *HTTP) congestionPort(c echo.Context) error {
	var Params struct {
		Port_ID string `json:"region_id"`
		Segment string `json:"segment"`
	}

	c.Bind(&Params)
	_, err := primitive.ObjectIDFromHex(Params.Port_ID)
	if err != nil {
		return c.String(http.StatusBadRequest,
			"bad valud for query parameter 'port_id' - only mongo ID type is accepted - you sent: "+Params.Port_ID)
	}
	if !model.CheckValidSegment(Params.Segment) {
		return c.String(http.StatusBadRequest, model.ErrInvalidSegment)
	}

	s := strings.Title(strings.ToLower(Params.Segment))
	result, err := h.svc.CongestionPort(c, Params.Port_ID, s)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) congestionRegion(c echo.Context) error {
	var Params struct {
		Region_ID string `json:"region_id"`
		Segment   string `json:"segment"`
	}
	c.Bind(&Params)

	regID := strings.ToUpper(Params.Region_ID)
	if regID == "" {
		return c.String(http.StatusBadRequest, model.ErrMissingRegionID)
	}
	if !model.CheckValidSegment(Params.Segment) {
		return c.String(http.StatusBadRequest, model.ErrInvalidSegment)
	}

	s := strings.Title(strings.ToLower(Params.Segment))
	result, err := h.svc.CongestionRegion(c, regID, s)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) listPort(c echo.Context) error {
	result, err := h.svc.ListPort(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) listRegion(c echo.Context) error {
	result, err := h.svc.ListRegion(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
