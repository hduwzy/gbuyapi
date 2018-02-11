package results

import "gbuyapi/models"

type GetUserAddressListResult struct {
	Total int `json:"total"`
	List []models.UserAddress `json:"用户地址"`
}
