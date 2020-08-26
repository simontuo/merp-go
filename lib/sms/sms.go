package sms

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/simontuo/merp-go/config"

	"github.com/simontuo/merp-go/helper/helper"

	uuid "github.com/satori/go.uuid"
)

type HuaWeyYunSms struct {
	Body struct {
		From          string
		To            string
		TemplateId    string
		TemplateParas []string
	}
}

type HuaWeiYunRsp struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Result      []struct {
		OriginTo   string `json:"originTo"`
		CreateTime string `json:"createTime"`
		From       string `json:"from"`
		SmsMsgId   string `json:"smsMsgId"`
		Status     string `json:"status"`
	}
}

func (s *HuaWeyYunSms) Send() (*HuaWeiYunRsp, error) {
	// api地址
	uri := config.G_SMS.Host + config.G_SMS.SentSmsUri
	// post 参数
	data := url.Values{}
	templateParas, err := json.Marshal(s.Body.TemplateParas)
	if err != nil {
		return nil, err
	}
	s.Body.From = config.G_SMS.SignatureCommunication
	data.Set("from", s.Body.From)
	data.Set("to", s.Body.To)
	data.Set("templateId", s.Body.TemplateId)
	data.Set("templateParas", string(templateParas))
	data.Set("signature", config.G_SMS.SignatureName)
	// 构造请求
	req, err := http.NewRequest("POST", uri, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	// 设置头部参数
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", `WSSE realm="SDP",profile="UsernameToken",type="Appkey"`)
	req.Header.Set("X-WSSE", buildWsseHeader())
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	res, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	var smsRsp HuaWeiYunRsp
	if err := json.Unmarshal(res, &smsRsp); err != nil {
		return nil, err
	}

	return &smsRsp, nil
}

func buildWsseHeader() string {
	created := helper.Time2UtcString(time.Now())
	nonce := uuid.Must(uuid.NewV4(), nil).String()

	hashString := nonce + created + config.G_SMS.AppSecret
	passwordDigest := base64.StdEncoding.EncodeToString(GetSHA256HashCode([]byte(hashString)))

	return fmt.Sprintf(`UsernameToken Username="%s",PasswordDigest="%s",Nonce="%s",Created="%s"`, config.G_SMS.AppKey, passwordDigest, nonce, created)
}

func GetSHA256HashCode(message []byte) []byte {
	hash := sha256.New()
	hash.Write(message)
	return hash.Sum(nil)
}
