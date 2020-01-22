package transport

import model "oceanbolt.com/iamservice/pkg/utl/model"

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*model.Port
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []model.Port `json:"ports"`
	}
}
