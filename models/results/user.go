package results

import "gbuyapi/models"

type GetUserAddressListResult struct {
	Total int                  `json:"total"`
	List  []models.UserAddress `json:"list"`
}

type LoginResult struct {
	IsSuccess bool   `json:"is_success"`
	Sid       string `json:"sid"`
}
