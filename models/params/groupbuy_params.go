package params

type BindGroupbuyGoodsRequest struct {
	Gid         int   `json:"gid"`
	GoodsIdList []int `json:"goods_id_list"`
}

type AddGroupbuyParam struct {
	Gbuy    AddGroupbuyGbuyParam `json:"gbuy"`
	Details []AddGroupbuyDetailParam `json:"details"`
}

type AddGroupbuyGbuyParam struct {
	Title       string `json:"title"`
	Detail      string `json:"detail"`
	ImgUrl      string `json:"img_url"`
	ReceiveType int8   `json:"receive_type"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

type AddGroupbuyDetailParam struct {
	GoodsId int `json:"goods_id"`
	SkuId   int `json:"sku_id"`
	Price   int `json:"price"`
	Amount  int `json:"amount"`
}
