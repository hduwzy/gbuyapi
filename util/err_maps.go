package util

const SYS_PARAMS_ERR = 21901 // 参数解析错误
const SYS_ERR = 21999 // 参数解析错误

const GOODS_NOT_EXISTS = 21001 // 商品不存在
const GOODS_UNVALID_INFO = 21002 // 商品信息不全

var err_msp = map[int]string{
	GOODS_NOT_EXISTS: "商品不存在",
}

func GetErrorMsg(code int) string {
	if msg, ok := err_msp[code]; ok {
		return msg
	} else {
		return "未知错误"
	}
}
