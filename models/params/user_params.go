package params


type UserLoginRequest struct {
	Code string `json:"code"`
	Nickname string `json:"nick_name"`
	AvatUrl string `json:"ava_url"`
}
