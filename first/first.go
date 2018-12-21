package first

import (
	"fmt"
	"os"

	"math/rand"

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

	fmt.Println("数据库连接成功")
}

type Infor struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Links string `json:"links" form:"links"`
	Times string `json:"times" form:"times"`
}

func SleInfor1() ([]Infor, error) {
	mod := make([]Infor, 0)
	err := dB.Select(&mod, `select * from informations`)
	return mod, err
}

func SleInfor2() ([]Infor, error) {
	mod := make([]Infor, 0)
	err := dB.Select(&mod, `select * from information2`)
	return mod, err
}

func SleInfor3() ([]Infor, error) {
	mod := make([]Infor, 0)
	err := dB.Select(&mod, `select * from information3`)
	return mod, err
}

type Products struct {
	ID      int     `json:"id,omitempty" form:"id"`
	Imgs    string  `json:"imgs,omitempty" form:"imgs"`
	Message string  `json:"message,omitempty" form:"message"`
	Price   float64 `json:"price,omitempty" form:"price"`
	Hrefs   string  `json:"hrefs,omitempty"`
}

//首页产品
func SleProdu() ([]Products, error) {
	mod := make([]Products, 0)
	err := dB.Select(&mod, `select * from products`)
	return mod, err
}

type IndexMessages struct {
	ID      int    `json:"id" form:"id"`
	Imgs    string `json:"imgs" form:"imgs"`
	Mess    string `json:"mess" form:"mess"`
	Content string `json:"content" form:"content"`
}

func SleIndexMess() ([]IndexMessages, error) {
	mod := make([]IndexMessages, 0)
	err := dB.Select(&mod, `select * from indexmessage`)
	return mod, err
}

type IndexNav struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Hrefs string `json:"hrefs" form:"hrefs"`
}

func SleNavs() ([]IndexNav, error) {
	mod := make([]IndexNav, 0)
	err := dB.Select(&mod, `select * from navs`)
	return mod, err
}

type Videos struct {
	ID    int    `json:"id" form:"id"`
	Video string `json:"video" form:"video"`
	Times string `json:"times" form:"times"`
}

func SleVideo() ([]Videos, error) {
	mod := make([]Videos, 0)
	err := dB.Select(&mod, `select * from videos`)
	return mod, err
}

//轮播图

type Img struct {
	ID   int    `json:"id" form:"id"`
	Imgs string `json:"imgs" form:"imgs"`
	Hre  string `json:"hre" form:"hre"`
}

func SleLunImg() ([]Img, error) {
	mod := make([]Img, 0)
	err := dB.Select(&mod, `select * from lunboimg`)
	return mod, err
}

//云上集市优质产品
type YSJSprod struct {
	ID     int     `json:"id" form:"id"`
	Imgs   string  `json:"imgs" form:"imgs"`
	Mess   string  `json:"mess" form:"mess"`
	Momess string  `json:"momess" form:"momess"`
	Price  float64 `json:"price" form:"price"`
	Hrefs  string  `json:"hrefs" form:"hrefs"`
}

//云上集市优质产品
func SleYSJS() ([]YSJSprod, error) {
	mod := make([]YSJSprod, 0)
	err := dB.Select(&mod, `select * from ysjsgood`)
	return mod, err
}

//云上集市相关产品
type YSJSpordu struct {
	ID       int     `json:"id" form:"id"`
	Imgs     string  `json:"imgs" form:"imgs"`
	Mess     string  `json:"mess" form:"mess"`
	Moremess string  `json:"moremess" form:"moremess"`
	Number   string  `json:"number" form:"number"`
	Price    float64 `json:"price" form:"price"`
	Hrefs    string  `json:"hrefs" form:"hrefs"`
}

//云上集市罗列产品
func Sleprodu() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	err := dB.Select(&mod, `select * from yunshop1 limit 6`)
	return mod, err
}

func Sleprodu1() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	err := dB.Select(&mod, `select * from yunshop1 limit 6,6`)
	return mod, err
}

//友情链接
type Friends struct {
	ID    int    `json:"id" form:"id"`
	Names string `json:"names" form:"names"`
	Hrefs string `json:"hrefs" form:"hrefs"`
}

func SelFriends() ([]Friends, error) {
	mod := make([]Friends, 0)
	err := dB.Select(&mod, `select * from friendlink`)
	return mod, err
}

//用户注册信息
type UserRegister struct {
	ID      int    `json:"id" xml:"id"`
	Tel     string `json:"tel" xml:"tel"`
	Tagname string `json:"tagname" xml:"tagname"`
	Email   string `json:"email" xml:"email"`
	Imgs    string `json:"imgs" xml:"imgs"`
}

//购物车信息
type GWC struct {
	ID      int    `json:"id" form:"id"`
	Goodsid int    `json:"goodsid" form:"goodsid"`
	Userid  int    `json:"userid" form:"userid"`
	Num     int    `json:"num" form:"num"`
	Addr    string `json:"addr" form:"addr"`
	Imgs    string `json:"imgs" form:"imgs"`
	Name    string `json:"name" form:"name"`
	Msgs    string `json:"msgs" form:"msgs"`
	Price   string `json:"price" form:"price"`
}

//在购物车中显示用户信息
func ShowUserMsg(id int) (UserRegister, error) {
	mod := UserRegister{}
	err := dB.Get(&mod, `select imgs, tagname, tel, email from regis where id=?`, id)
	return mod, err
}

//购物车
func InsertGWC(goodsid, userid, num int, addr, imgs, name, msgs, price string) bool {
	// mod := GWC{}
	mod, _ := dB.Exec(`insert into gwc(goodsid,userid, num, addr,imgs,name,msgs,price) values(?,?,?,?,?,?,?,?)`, goodsid, userid, num, addr, imgs, name, msgs, price)
	bef, _ := mod.LastInsertId()
	if bef == 1 {
		return true
	}
	return false
}

// 在购物车中展示用户的地址信息
// func ShowAddre() ([]GWC, error) {
// 	mod := make([]GWC, 0)
// 	err := dB.Select(&mod, `select * from gwc where userid=2`)
// 	return mod, err
// }

//显示在购物车上的信息
func Showgwc(userid int) ([]GWC, error) {
	mod := make([]GWC, 0, 2)
	err := dB.Select(&mod, `select * from gwc where userid=?`, userid)
	// err := dB.Select(&mod, `select * from yunshop1`)
	// fmt.Println(4654)
	return mod, err
}

//显示在购物车上的购物数量
// func Showgwcnum() (GWC, error) {
// 	mod := GWC{}
// 	err := dB.Get(&mod, `select addr, num from gwc where userid=2`)
// 	// err := dB.Select(&mod, `select * from yunshop1`)
// 	// fmt.Println(4654)
// 	return mod, err
// }

//从购物车中移除相应商品
func RemoveGod(id int) bool {
	mod, _ := dB.Exec(`delete from gwc where id=?`, id)
	bef, _ := mod.RowsAffected()
	if bef == 1 {
		return true
	} else {
		return false
	}
}

type Singleimgmsg struct {
	ID    int     `json:"id" form:"id"`
	Hrefs string  `json:"hrefs" form:"hrefs"`
	Mess  string  `json:"mess" form:"mess"`
	Price float64 `json:"price" form:"price"`
}

func SingleLittleShop() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	var k = rand.Intn(8)
	err := dB.Select(&mod, `select * from yunshop1 limit ?,4`, k)
	return mod, err
}

func SingleBigShop(id int) ([]Singleimgmsg, error) {
	mod := make([]Singleimgmsg, 0)
	var k = (id - 1) * 5
	err := dB.Select(&mod, `select * from singleimgmsg limit ?,5`, k)
	// fmt.Println(err)
	return mod, err
}

//single表的定义
type Singledetial struct {
	ID    int     `json:"id" form:"id"`
	Hrefs string  `json:"hrefs" form:"hrefs"`
	Mess  string  `json:"mess" form:"mess"`
	Price float64 `json:"price" form:"price"`
	Ids   string  `json:"ids" form:"ids"`
	Idjs  string  `json:"idjs" form:"idjs"`
}

func LunSingleImg(id int) ([]Singledetial, error) {
	mod := make([]Singledetial, 0)
	var k = (id - 1) * 5
	err := dB.Select(&mod, `select * from singledetial limit ?,5`, k)
	// fmt.Println(err)
	return mod, err
}

//产品详情信息
type Goodsmsg struct {
	ID           int     `json:"id" form:"id"`
	Imgs         string  `json:"imgs" form:"imgs"`
	Goodsname    string  `json:"goodsname" form:"goodsname"`
	GoodsMsg     string  `json:"goodsmsg"  form:"goodsmsg"`
	Oldprice     float64 `json:"oldprice" form:"oldprice"`
	Newprice     float64 `json:"newprice" form:"newprice"`
	Successgoods string  `json:"successgoods" form:"successgoods"`
}

func GoodsMSG(id int) (Goodsmsg, error) {
	mod := Goodsmsg{}
	err := dB.Get(&mod, `select * from gooodsmsg where id=?`, id)
	return mod, err
}

type InserAdd struct {
	ID      int    `json:"id" form:"id"`
	Userid  int    `json:"userid" form:"userid"`
	Goodsid int    `json:"goodsid" form:"goodsid"`
	Addr    string `json:"addr" form:"addr"`
	Num     int    `json:"num" form:"num"`
}

//发送用户的收获地址信息
func Inseradd(userid, goodsid int, address string, num int) bool {
	mod, _ := dB.Exec(`insert into address(userid,goodsid,addr,num) values(?,?,?,?)`, userid, goodsid, address, num)
	ok, _ := mod.LastInsertId()
	if ok == 1 {
		return true
	} else {
		return false
	}
}

//关于我们页面的相关数据库操作
type Personinfor struct {
	ID  int    `json:"id,omitempty" form:"id"`
	Msg string `json:"msg,omitempty" form:"msg"`
}

//查询平台特色的相关信息
func AllInfor() ([]Personinfor, error) {
	mod := make([]Personinfor, 0)
	err := dB.Select(&mod, `select * from perinfor`)
	return mod, err
}

func Movemsg() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	err := dB.Select(&mod, `select imgs,mess from yunshop1 where id<13`)
	return mod, err
}

func Movenine() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	err := dB.Select(&mod, `select imgs,mess from yunshop1 where id<3`)
	return mod, err
}

func Moveeight() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	err := dB.Select(&mod, `select imgs,mess from yunshop1 where id>=3 and id<6`)
	return mod, err
}
func Moveseven() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	err := dB.Select(&mod, `select imgs,mess from yunshop1 where id>=6 and id<10`)
	return mod, err
}
func Movesix() ([]YSJSpordu, error) {
	mod := make([]YSJSpordu, 0)
	err := dB.Select(&mod, `select imgs,mess from yunshop1 where id>=10 and id<13`)
	return mod, err
}

//用户反馈
type FeedBack struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Tel   string `json:"tel,omitempty"`
	Msg   string `json:"msg,omitempty"`
}

func FeedMsgs(name, email, tel, msg string) bool {
	mod, _ := dB.Exec(`insert into feedback(name, email, tel, msg) values(?,?,?,?)`, name, email, tel, msg)
	bef, _ := mod.LastInsertId()
	if bef == 1 {
		return true
	} else {
		return false
	}
}
