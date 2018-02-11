package results

import "gbuyapi/models"

type GetOrderListResult struct {
	Total int	`json:"total"`
	List []models.Orders `json:"list"`
}