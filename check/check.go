package check

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
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

//定义用户注册的信息

type UserRegister struct {
	ID        int    `json:"id,omitempty" form:"id"`
	Name      string `json:"name,omitempty"`
	Idcard    string `json:"idcard,omitempty"`
	Email     string `json:"email" from:"email"`
	Address   string `json:"address,omitempty"`
	Tel       string `json:"tel,omitempty"`
	Newname   string `json:"newname,omitempty"`
	Pass      string `json:"pass" form:"pass"`
	Againpass string `json:"againpass,omitempty"`
}

//创建用户数据库

//往数据库中插入用户注册的信息
func Useradd(name string, idcard string, email string, address string, tel string, newname string, pass string, againpass string) (UserRegister, error) {
	s := UserRegister{}
	sm, err := dB.Exec(`insert into register(name, idcard, email, address, tel, newname, pass, againpass) values(?,?,?,?,?,?,?,?)`, name, idcard, email, address, tel, newname, pass, againpass)
	sm.LastInsertId()
	return s, err
}

func CheckEmail(email string) ([]UserRegister, error) {
	mod := make([]UserRegister, 0)
	err := dB.Select(&mod, `select * from register`)
	return mod, err
}

//在数据库中进行查找name和pass
func CheckUser(name string) (UserRegister, error) {

	mods := UserRegister{}
	err := dB.Get(&mods, `select * from register where name=?`, name)
	return mods, err
}

//忘记密码进行的操作
func Lostpass(email string) (UserRegister, error) {
	dp := UserRegister{}
	err := dB.Get(&dp, `select * from register where email=?`, email)
	return dp, err
}

func UpDataPass(pass string, againpass string, email string) (UserRegister, error) {
	udp := UserRegister{}
	_, err := dB.Exec(`update register set pass=?, againpass=? where email=?`, pass, againpass, email)
	return udp, err
}

// Jwt json web token
type Jwt struct {
	Id   int
	Name string
	jwt.StandardClaims
}
