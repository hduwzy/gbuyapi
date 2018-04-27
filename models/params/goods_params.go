package params

type AddGoodsParams struct {
	Title  string `json:"title"`
	CateId int   `json:"cate_id"`
	Detail string `json:"detail"`

	Skus []AddGoodsSkuParam `json:"skus"`
}

type AddGoodsSkuParam struct {
	SkuTitle string `json:"sku_title"`
	Price    int    `json:"price"`
	Unit     string `json:"unit"`
	Stock    int    `json:"stock"`
	ImgUrl   string `json:"img_url"`
}

type UpdateGoodsParams struct {
	Id     int    `json:"goods_id"`
	Title  string `json:"title"`
	CateId int    `json:"cate_id"`
	Detail string `json:"detail"`
	Skus []UpdateGoodsSkuParams `json:"skus"`
}

type UpdateGoodsSkuParams struct {
	SkuId    int    `json:"sku_id"`
	SkuTitle string `json:"sku_title"`
	Price    int    `json:"price"`
	Unit     string `json:"unit"`
	Stock    int    `json:"stock"`
	ImgUrl   string `json:"img_url"`
}
