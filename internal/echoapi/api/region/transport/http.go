package transport

import (
	"net/http"

	"oceanbolt.com/obapi/internal/echoapi/api/region"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc region.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc region.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/congestion")

	// swagger:operation GET /v1/users users listUsers
	// ---
	// summary: Returns list of users.
	// description: Returns list of users. Depending on the user role requesting it, it may return all users for SuperAdmin/Admin users, all company/location users for Company/Location admins, and an error for non-admin users.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/userListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/region/list", h.list)

	// swagger:operation GET /v1/users/{id} users getUser
	// ---
	// summary: Returns a single user.
	// description: Returns a single user by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of user
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/userResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/region", h.view)
}

func (h *HTTP) view(c echo.Context) error {
	var Params struct {
		Region_ID string `form:"region_id"`
		Segment   string
	}
	c.Bind(&Params)
	result, err := h.svc.View(c, Params.Region_ID, Params.Segment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (h *HTTP) list(c echo.Context) error {
	var Params struct {
		Region_ID string `json:"region_id"`
		Segment   string `json:"segment"`
	}
	c.Bind(&Params)
	result, err := h.svc.List(c, Params.Region_ID, Params.Segment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
