package config

import (
	"encoding/json"
	"io/ioutil"
)

type DB struct {
	DBHost        string `json:"dbHost"`
	DBPort        string `json:"dbPort"`
	DBUser        string `json:"dbUser"`
	DBPassword    string `json:"dbPassword"`
	DBDatabase    string `json:"dbDatabase"`
	DBDebug       bool   `json:"dbDebug"`
	DevDBHost     string `json:"devdbHost"`
	DevDBPort     string `json:"devdbPort"`
	DevDBUser     string `json:"devdbUser"`
	DevDBPassword string `json:"devdbPassword"`
	DevDBDatabase string `json:"devdbDatabase"`
	DevDBDebug    bool   `json:"devdbDebug"`
}

var G_DB *DB

func InitDB(fileName string, dataBase string) error {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	var conf DB

	if err := json.Unmarshal(content, &conf); err != nil {
		return err
	}

	G_DB = &conf
	G_DB.DBDatabase = dataBase
	if G_Config.Dev {
		G_DB.DBHost = conf.DevDBHost
		G_DB.DBPort = conf.DevDBPort
		G_DB.DBUser = conf.DevDBUser
		G_DB.DBPassword = conf.DevDBPassword
		G_DB.DBDebug = conf.DevDBDebug
	}
	return nil
}
