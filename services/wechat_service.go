package services

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/oauth"
	"github.com/silenceper/wechat/cache"
	"errors"
	"github.com/rs/xid"
	"gbuyapi/models"
	"time"
	"encoding/json"
	"github.com/astaxie/beego/context"
)

type Session map[string]interface{}

//
//func (s Session) LoadSession(sessionId string) Session {
//
//}

var (
	EmptyString            = ""
	EmptyInt               = 0

	ErrSessionNotExists    = errors.New("session not exists")
	ErrSessionKeyNotExists = errors.New("session key not exists")
)

func NewSession() (Session, string) {
	sid := xid.New().String()
	s := make(Session)
	sess[sid] = s
	return s, sid
}

func DestroySession(sid string) {

}

var sess map[string]Session

func LoadSession(sid string) Session {
	if nil == sess {
		sess = make(map[string]Session)
		return nil
	}

	if _, ok := sess[sid]; !ok {
		return nil
	}

	return sess[sid]
}

func (s Session) GetInt(key string) (int, error) {

	if res, ok := s[key]; ok {
		if result, assOk := res.(int); assOk {
			return result, nil
		}
	}

	return EmptyInt, ErrSessionKeyNotExists
}

func (s Session) GetString(key string) (string, error) {

	if res, ok := s[key]; ok {
		if result, assOk := res.(string); assOk {
			return result, nil
		}
	}

	return EmptyString, ErrSessionKeyNotExists
}

func (s Session) GetBool(key string, b bool) (bool, error) {

	if res, ok := s[key]; ok {
		if result, assOk := res.(bool); assOk {
			return result, nil
		}
	}

	return false, ErrSessionKeyNotExists
}

func (s Session) Set(key string, val interface{}) {
	s[key] = val
}

type WechatService struct {
	BaseService
	wechat *wechat.Wechat
	oauth  *oauth.Oauth
}

func NewWechatService() *WechatService {
	cfg := wechat.Config{
		AppID:     "wx33235b232bc0b664",
		AppSecret: "510313fd901ca47137ab46f7ea13a713",
		Cache:     cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379"}),
	}

	w := wechat.NewWechat(&cfg)
	s := WechatService{}
	s.wechat = w
	s.oauth = w.GetOauth()
	return &s
}

func (w *WechatService) login(code string, fn func(token, openId string) string) string {
	res, e := w.oauth.GetUserAccessToken(code)
	if e != nil {
		panic(e)
	}

	return fn(res.AccessToken, res.OpenID)
}

func (w *WechatService) LoginSubject(code string, avatUrl string, nick string) (sessionId string) {
	return w.login(code, func(token, openId string) string {
		qs := w.Orm().QueryTable("users")
		u := models.Users{}
		err := qs.Filter("open_id__eq", openId).One(&u)
		if err != nil {
			u.Nickname = nick
			u.AvaUrl = avatUrl
			u.OpenId = openId
			u.AddTime = time.Now()
			u.UpdateTime = time.Now()
			w.Orm().Insert(&u)
		}
		sess, sid := NewSession()
		userInfo, _ := json.Marshal(u)
		sess.Set("user_info", string(userInfo))
		return sid
	})
}


func (w *WechatService) Logout() {
	
}

const SidKey = "__SID__"

func Sess(c *context.Context) Session {
	sid := c.Request.Header.Get(SidKey)
	if len(sid) == 0 {
		return nil
	}

	return LoadSession(sid)
}