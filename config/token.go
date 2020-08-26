package config

import (
	"encoding/json"
	"io/ioutil"
)

type Token struct {
	Key           string            `json:"key"`
	ExpiresAt     int               `json:"expiresAt"`
	PassEndpoints map[string]string `json:"passEndpoints"`
	PassPath      []string          `json:"passPath"`
}

var G_Token *Token

func InitToken(fileName string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var conf Token

	if err := json.Unmarshal(content, &conf); err != nil {
		return err
	}

	G_Token = &conf

	return nil
}
