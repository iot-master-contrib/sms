package tencent

import (
	"log"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type Sender struct {
	client *sms.Client
	AppId  string
}

func (s *Sender) Connect(id, key string) (err error) {
	credential := common.NewCredential(id, key)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"
	s.client, err = sms.NewClient(credential, "ap-guangzhou", cpf)
	return
}

func (s *Sender) Send(phone []string, sign, code string, params []string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = common.StringPtr(s.AppId)
	req.SignName = common.StringPtr(sign)
	req.TemplateId = common.StringPtr(code)
	req.TemplateParamSet = common.StringPtrs(params)
	req.PhoneNumberSet = common.StringPtrs(phone)
	req.SessionContext = common.StringPtr("")
	resp, err := s.client.SendSms(req)
	if err != nil {
		log.Println(resp)
	}
	return err
}
