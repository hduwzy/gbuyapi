package results

import "gbuyapi/models"

type GetGroupbuyResponse struct {
	Gbuy models.Groupbuy `json:"gbuy"`
	GbuyDetail []models.GroupbuyDetail `json:"gbuy_detail"`
}

type GetGroupbuyListRespongse struct {
	Gbuys []models.Groupbuy `json:"gbuys"`
}
