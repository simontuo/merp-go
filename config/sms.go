package config

import (
	"encoding/json"
	"io/ioutil"
)

type SMS struct {
	AppKey                  string `json:"appKey"`
	AppSecret               string `json:"appSecret"`
	Host                    string `json:"host"`
	SignatureCommunication  string `json:"signatureCommunication"`
	SignatureName           string `json:"signatureName"`
	SentSmsUri              string `json:"sentSmsUri"`
	LoginTemplateId         string `json:"loginTemplateId"`
	ResetPasswordTemplateId string `json:"resetPasswordTemplateId"`
}

var G_SMS *SMS

func InitSMS(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var conf SMS

	if err := json.Unmarshal(content, &conf); err != nil {
		return err
	}

	G_SMS = &conf

	return nil
}
