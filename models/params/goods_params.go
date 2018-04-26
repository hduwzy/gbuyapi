package params

type AddGoodsParams struct {
	Title  string `json:"title"`
	CateId int    `json:"cate_id"`
	Detail string `json:"detail"`

	Skus []struct {
		SkuTitle string `json:"sku_title"`
		Price    int    `json:"price"`
		Unit     string `json:"unit"`
		Stock    int    `json:"stock"`
		ImgUrl   string `json:"img_url"`
	} `json:"skus"`
}
