package auth

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/nutsdo/go-next/session"
	"net/url"
)

func Authorize(ctx iris.Context) {

	s := session.Sess.Start(ctx)
	sm := s.GetAll()
	if _,ok := sm["user"];ok {
		ctx.Next()
	}else {
		wxurl:="https://open.weixin.qq.com/connect/qrconnect?%s#wechat_redirect"
		params := url.Values{}
		params.Add("appid","wxdd6335bb645b19fc")
		params.Add("redirect_uri","http://www.judazhe.com")
		params.Add("response_type","code")
		params.Add("scope","snsapi_login")
		params.Add("state","12345")
		s:= params.Encode()
		//跳转到微信授权登录
		ctx.Redirect(fmt.Sprintf(wxurl,s))
	}

}

//func GenerateRandomString(len uint32) string{
//	//rand.Intn()
//}