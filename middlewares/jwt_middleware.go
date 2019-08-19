package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/context"
)
func JwtHandler() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("dKe580xiDDtIhttpClnsBXigizsR9xys"), nil
		},
		//自定义错误响应
		ErrorHandler:ErrorResp,
		//自定义获取token的方式
		//Extractor: FromDataDescription(),
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})


}


func ErrorResp(ctx context.Context, err string) {
	ctx.Text("error:" + err)
}


//func FromDataDescription() jwtmiddleware.TokenExtractor {
//	return func(ctx context.Context) (string, error) {
//
//		dataDescription :=ctx.FormValue("DataDescription")
//		data:=new(responses.DataDescription)
//		err:=json.Unmarshal([]byte(dataDescription),data)
//		if err !=nil {
//			return "",nil
//		}
//		fmt.Println(data.Token)
//		return data.Token, nil
//	}
//}