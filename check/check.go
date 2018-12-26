package check

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//用户注册信息
type Regis struct {
	ID      int    `json:"id" xml:"id"`
	Num  string `json:"num" xml:"num"`
	Tel     string `json:"tel" xml:"tel"`
	Pass    string `json:"pass" xml:"pass"`
	Tagname string `json:"tagname" xml:"tagname"`
	Name    string `json:"name" xml:"name"`
	Sex     string `json:"sex" xml:"sex"`
	Idcard  string `json:"idcard" xml:"idcard"`
	Email   string `json:"email" xml:"email"`
	Imgs  string `json:"imgs" xml:"imgs"`
}



//用户注册生成得到账号
func InsertUserNum(nums string, tel string) error {
	_, err := dB.Exec(`update regis set num=? where tel=?`, nums, tel)
	return err
}

//显示用户昵称
func ShowUsermsg(id int) (Regis, error) {
	mod := Regis{}
	err := dB.Get(&mod, `select imgs, tagname from regis where id=?`, id)
	return mod, err
}

//显示用户所有个人信息
func ShowUsermsgall(id int) (Regis, error) {
	mod := Regis{}
	err := dB.Get(&mod, `select tagname, name, sex, idcard, email from regis where id=?`, id)
	return mod, err
}

//用户注册操作
func UserRegis(tel, pass, tagname, name, sex, idcard, email, imgs string) error   {
	
	_, err := dB.Exec(`insert into regis(tel, pass,tagname, name, sex, idcard, email, imgs) values(?,?,?,?,?,?,?,?)`, tel, pass, tagname, name, sex, idcard, email,imgs)
	return err
}

//上传用户头像
func UpUserimg(imgs string, id int) error {
	_, err := dB.Exec(`update regis set imgs=? where id=?`, imgs, id)
	return err
}

//用户添加个人信息
func InsertMsg(id int, tagname, name, sex, idcard, email string) error {
	_, err := dB.Exec(`update regis set tagname=?,name=?,sex=?, idcard=?, email=? where id=?`, tagname, name, sex, idcard, email, id)
	return err
}

// //个人中心显示用户信息
func UserName(id int) (Regis, error) {
	mod := Regis{}
	err := dB.Get(&mod, `select name from regis where id=?`, id)
	return mod, err
}

// func CheckEmail(email string) ([]Regis, error) {
// 	mod := make([]Regis, 0)
// 	err := dB.Select(&mod, `select * from register`)
// 	return mod, err
// }

//在数据库中进行查找tel和pass
func CheckUser(num string) (Regis, error) {

	mods := Regis{}
	err := dB.Get(&mods, `select * from regis where num=?`, num)
	return mods, err
}

//忘记密码进行的操作
func Lostpass(tel string) (Regis, error) {
	dp := Regis{}
	err := dB.Get(&dp, `select * from regis where tel=?`, tel)
	return dp, err
}

func UpDataPass(pass string, tel string) (Regis, error) {
	udp := Regis{}
	_, err := dB.Exec(`update regis set pass=? where tel=?`, pass, tel)
	return udp, err
}

//显示用户头像
func ShowUserims(id int) (Regis, error) {
	mod := Regis{}
	err := dB.Get(&mod, `select imgs,tagname from regis where id=?`, id)
	return mod, err
}

// Jwt json web token
type Jwt struct {
	Id  int
	Tel string
	jwt.StandardClaims
}