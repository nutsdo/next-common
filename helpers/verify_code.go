package helpers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/nutsdo/go-next/database"
	"math/rand"
	"strings"
	"time"
)

type VerifyCode struct {
	Phone string
	VerifyCodeType string //register login forgot password
	Code int32
	Platform string
}

func New(phone, tp string) *VerifyCode {
	verifyCode := &VerifyCode{
		Phone:phone,
		VerifyCodeType:tp,
	}
	return verifyCode
}


func (v *VerifyCode) VerifyCodeGenerate() string {

	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)
	v.Code = r
	v.StoreCache(time.Minute)

	return fmt.Sprintf("%04v",r)
}


func (v *VerifyCode) StoreCache(expired time.Duration)  {
	set,err := database.Rdb.Set(v.generateKey(),v,expired).Result()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(set)
}

func (v *VerifyCode) generateKey() string{
	var key strings.Builder
	key.WriteString(v.Phone)
	key.WriteString(string(v.Code))
	key.WriteString(v.Platform)
	key.WriteString(v.VerifyCodeType)
	return fmt.Sprintf("%x",md5.Sum([]byte(key.String())))
}

func (v *VerifyCode) Find() (string,error) {
	val,err:=database.Rdb.Get(v.generateKey()).Result()
	if err !=nil {
		return "",err
	}
	return val,nil
}

func (v *VerifyCode) MarshalBinary() ([]byte, error) {
	return json.Marshal(v)
}