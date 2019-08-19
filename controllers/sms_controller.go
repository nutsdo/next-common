package controllers

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/kataras/iris"
	"github.com/nutsdo/go-next/helpers"
	"github.com/nutsdo/go-next/responses"
	"time"
)

func SMSHandler(ctx iris.Context)  {

	// 获取短信服务类型
	smsType := ctx.URLParam("smsType")
	phone := ctx.FormValueDefault("phone","")
	fmt.Println("输入的手机号为：")
	fmt.Println(phone)
	if phone =="" {
		ctx.JSON(responses.ErrResponse(-1,201001,"手机号格式不正确"))
		return
	}

	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou",
		"LTAIjF0HDbPFUm4c", "Os2FnBSWD6ZEO2ySSC1qkn1LUbk2Iy")

	if err!=nil {
		ctx.JSON(responses.ErrResponse(-1,201002,err.Error()))
	}
	client.SetConnectTimeout(time.Second*10)
	switch smsType {
	case "send":
		start := time.Now().Unix()
		ctx.JSON(sendSMS(client,ctx))
		end := time.Now().Unix()
		fmt.Printf("调用发送短信函数消耗的时间为:%v秒\n", end - start)

	case "query":
		ctx.JSON(smsQuerySendDetails(client))
	case "sendBatch":
		ctx.JSON(sendBatchSMS(client))
	default:
		ctx.JSON(responses.ErrResponse(-1,200001,"发送短信类型错误"))
	}

}

func smsQuerySendDetails(client *dysmsapi.Client) *dysmsapi.QuerySendDetailsResponse {

	request := dysmsapi.CreateQuerySendDetailsRequest()
	request.Scheme = "https"

	response, err := client.QuerySendDetails(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return response
}

func sendSMS(client *dysmsapi.Client,ctx iris.Context) *dysmsapi.SendSmsResponse{

	phone := ctx.FormValueDefault("phone","")
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = "开门得宝"
	request.TemplateCode = "SMS_171076357"

	vc:= &helpers.VerifyCode{
		Phone:phone,
		VerifyCodeType: "register",
	}
	request.TemplateParam = fmt.Sprintf("{\"code\":%s}",vc.VerifyCodeGenerate())
	val,_:=vc.Find()
	fmt.Printf("获取生成的验证码缓存: %#v\n",val)
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	return response
}

func sendBatchSMS(client *dysmsapi.Client) *dysmsapi.SendBatchSmsResponse {
	request := dysmsapi.CreateSendBatchSmsRequest()
	request.Scheme = "https"

	response, err := client.SendBatchSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	return response
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}