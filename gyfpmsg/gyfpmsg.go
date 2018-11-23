package gyfpmsg

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dB *sqlx.DB

//连接数据库
func init() {
	var err error
	dB, err = sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/nmy?charsetutf8&parseTime=true`)
	if err != nil {
		fmt.Println("数据库请求失败")
		os.Exit(1)
	}
	if err = dB.Ping(); err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(1)
	}

	// fmt.Println("数据库连接成功")
}

type Gyfpmsg struct {
	ID        int    `json:"id" form:"id"`
	Imgs      string `json:"imgs" form:"imgs"`
	Messtitle string `json:"messtitle" form:"messtitle"`
	Msg       string `json:"msg" form:"msg"`
}

func GyfpMsg() (Gyfpmsg, error) {
	mod := Gyfpmsg{}
	err := dB.Get(&mod, `select * from gyfpmsg`)
	return mod, err
}

type Gyfpcountry struct {
	ID              int    `json:"id" form:"id"`
	Countrysidename string `json:"countrysidename" form:"countrysidename"`
	Links           string `json:"links" form:"links"`
}

func GyCountry() ([]Gyfpcountry, error) {
	mod := make([]Gyfpcountry, 0, 3)
	err := dB.Select(&mod, `select * from gyfpcountry`)
	// fmt.Println(err)
	return mod, err
}

type TextMsg struct {
	ID      int    `json:"id,omitempty"`
	Imgs    string `json:"imgs,omitempty"`
	Gytitle string `json:"gytitle,omitempty"`
	Gymsg   string `json:"gymsg,omitempty"`
}

func GyTextmsg() ([]TextMsg, error) {
	mod := make([]TextMsg, 0)
	err := dB.Select(&mod, `select * from gyfpstory`)
	return mod, err
}
