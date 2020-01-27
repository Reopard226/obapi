package apiaccess

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"

	"oceanbolt.com/obapi/internal/echoapi/api/apiaccess"
	"oceanbolt.com/obapi/internal/echoapi/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc apiaccess.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc apiaccess.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/apiaccess")

	ur.GET("/keys", h.listKey)
	ur.POST("/keys", h.createKey)
	ur.DELETE("/keys/:apikey_id", h.deleteKey)
}

func (h *HTTP) listKey(c echo.Context) error {
	result, err := h.svc.ListKey(c)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) createKey(c echo.Context) error {
	type Params struct {
		Tag string `json:"tag"`
		Exp int64  `json:"exp"`
	}
	var payload Params
	err := jsoniter.NewDecoder(c.Request().Body).Decode(&payload)
	if err != nil {
		return c.String(http.StatusBadRequest, model.ErrParseCreatKeyParams)
	}
	if payload.Tag == "" {
		return c.String(http.StatusBadRequest, model.ErrMissingTag)
	}

	result, err := h.svc.CreateKey(c, payload.Tag, payload.Exp)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) deleteKey(c echo.Context) error {
	apiKeyID := c.Param("apikey_id")
	result, err := h.svc.DeleteKey(c, apiKeyID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
