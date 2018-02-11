package params

type BindGroupbuyGoodsRequest struct {
	Gid int `json:"gid"`
	GoodsIdList []int `json:"goods_id_list"`
}
