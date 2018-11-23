package yunsjs

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dB *sqlx.DB

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

type Img struct {
	ID    int    `json:"id" form:"id"`
	Imgs  string `json:"imgs" form:"imgs"`
	Hrefs string `json:"hrefs" form:"hrefs"`
}

func SleLunImg() ([]Img, error) {
	mod := make([]Img, 0)
	err := dB.Select(&mod, `select * from lunboimg`)
	return mod, err
}
