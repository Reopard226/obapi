package transport

import model "oceanbolt.com/obapi/internal/echoapi/utl/model"

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*model.Region
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Retions []model.Region `json:"regions"`
	}
}
