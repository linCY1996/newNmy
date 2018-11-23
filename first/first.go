package first

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
	ID      int     `json:"id" form:"id"`
	Imgs    string  `json:"imgs" form:"imgs"`
	Message string  `json:"message" form:"message"`
	Price   float64 `json:"price" form:"price"`
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
	ID     int     `json:"id" form:"id"`
	Imgs   string  `json:"imgs" form:"imgs"`
	Mess   string  `json:"mess" form:"mess"`
	Number string  `json:"number" form:"number"`
	Price  float64 `json:"price" form:"price"`
	Hrefs  string  `json:"hrefs" form:"hrefs"`
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
	fmt.Println(err)
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

//购物车信息
type GWC struct {
	Userid   int `json:"userid" form:"userid"`
	Goodsid  int `json:"goodsid" form:"goodsid"`
	Sellerid int `json:"sellerid" form:"sellerid"`
}

//购物车
func InsertGWC(userid, goodsid int) (GWC, error) {
	mod := GWC{}
	_, err := dB.Exec(`insert into gwc(userid, goodsid, sellerid) values(?,?,'张三')`, userid, goodsid)
	return mod, err
}

type Singleimgmsg struct {
	ID    int     `json:"id" form:"id"`
	Hrefs string  `json:"hrefs" form:"hrefs"`
	Mess  string  `json:"mess" form:"mess"`
	Price float64 `json:"price" form:"price"`
}

func SingleLittleShop(id int) ([]Singleimgmsg, error) {
	mod := make([]Singleimgmsg, 0)
	var k = (id - 1) * 2
	err := dB.Select(&mod, `select * from singleimgmsg limit ?,7`, k)
	return mod, err
}

func SingleBigShop(id int) ([]Singleimgmsg, error) {
	mod := make([]Singleimgmsg, 0)
	var k = (id - 1) * 6
	err := dB.Select(&mod, `select * from singleimgmsg limit ?,6`, k)
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
	Userid  string `json:"userid" form:"userid"`
	Goodsid string `json:"goodsid" form:"goodsid"`
	Addr    string `json:"addr" form:"addr"`
}

//发送用户的收获地址信息
func Inseradd(goodsid, address string) bool {
	mod, _ := dB.Exec(`insert into address(goodsid,addr) values(?,?)`, goodsid, address)
	ok, _ := mod.LastInsertId()
	if ok == 1 {
		return true
	} else {
		return false
	}
}
