package abouts

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

// type Aboutmsg struct {
// 	ID      int    `json:"id,omitempty"`
// 	Msg1    string `json:"msg_1,omitempty"`
// 	Msg2    string `json:"msg_2,omitempty"`
// 	Litemsg string `json:"litemsg,omitempty"`
// }

// func Aboutmymsg() (Aboutmsg, error) {
// 	mod := Aboutmsg{}
// 	err := dB.Get(&mod, `select * from aboutmsg where id =1`)
// 	return mod, err
// }

// func Aboutmy() ([]Aboutmsg, error) {
// 	mod := make([]Aboutmsg, 0)
// 	err := dB.Select(&mod, `select * from aboutmsg`)
// 	return mod, err
// }

// type Friends struct {
// 	ID   int    `json:"id,omitempty"`
// 	Imgs string `json:"imgs,omitempty"`
// 	Name string `json:"name,omitempty"`
// 	Mess string `json:"mess,omitempty"`
// }

// func Aboutfr() ([]Friends, error) {
// 	mod := make([]Friends, 0)
// 	err := dB.Select(&mod, `select * from aboutfr`)
// 	return mod, err
// }
