package results

import "gbuyapi/models"

type GetGoodsRsponse struct {
	Goods models.Goods `json:"goods"`
	Skus []GetGoodsSkuRsponse
}

type GetGoodsSkuRsponse struct {
	Sku models.Sku `json:"sku"`
	Img string `json:"img"`
}