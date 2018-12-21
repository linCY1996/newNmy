package main

import (
	"NMY/check"
	"NMY/control"
	"NMY/first"
	"NMY/gyfpmsg"
	"NMY/util"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"

	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const tokenKey = "this is a bdd"

//核对用户登录时的用户名和密码是否一致
func Checklogo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var num = r.FormValue("num")
	// var uPass = r.FormValue(`mima`)
	var pass = r.Form.Get("mima")
	mods, err := check.CheckUser(num)
	if err != nil {
		w.Write([]byte(`信息错误1`))
		return
	}
	if num != mods.Num {
		w.Write([]byte(`信息错误2`))
		return
	}

	if mods.Pass != pass {
		w.Write([]byte(`信息错误3`))
		return
	}
	data := check.Jwt{
		Id:  mods.ID,
		Tel: mods.Tel,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}
	jwts := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err := jwts.SignedString([]byte(tokenKey))
	// fmt.Println(err)
	if err != nil {
		w.Write(util.NewResult(util.Fail, "token生成失败，请重试"))
		return
	}
	w.Write(util.NewResult(util.Success, `登录成功`, token))

}

func Lostpa(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var mail = r.FormValue(`tel`)
	var pass = r.FormValue(`pass1`)

	lop, _ := check.Lostpass(mail)
	if lop.Email != mail {
		w.Write([]byte(`此邮箱未被注册`))
	} else {
		_, err := check.UpDataPass(pass, mail)
		if err != nil {
			w.Write([]byte(`输入不合法`))
			return
		}
		// v, _ := json.Marshal(uDP)
		// w.Header().Set(`Content-Type`, `application/json`)

		w.Write([]byte(`修改密码成功`))
	}
}

//显示登录页面
func Viewcheck(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/check.html`)
	w.Write(buf)
}

//显示注册页面
func ViewLogin(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("view/register.html")
	w.Write(buf)
}

//云上集市页面
func Viewysjs(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("view/ysjs.html")
	w.Write(buf)
}

//忘记密码页面
func ViewLost(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("view/lost.html")
	w.Write(buf)
}

//首页
func ViewMain(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("view/viewmain.html")
	w.Write(buf)
}

func ViewSingle(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("view/single.html")
	w.Write(buf)
}

//信息栏目
func Infor1(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleInfor1()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

func Infor2(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleInfor2()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

func Infor3(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleInfor3()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//导航
func Titls(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleProdu()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//条目
func Indexmess(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleIndexMess()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//导航
func Navx(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleNavs()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//video数据
func Vide(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleVideo()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//云上集市中优质产品
func Goodmes(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleYSJS()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//轮播图
func YunsjsImg(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SleLunImg()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//云上集市产品第一栏
func Yunsproduct(w http.ResponseWriter, r *http.Request) {
	mod, err := first.Sleprodu()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//云上集市产品第二栏
func Yunsproduct2(w http.ResponseWriter, r *http.Request) {
	mod, err := first.Sleprodu1()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//友情链接
func Frined(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SelFriends()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	bn, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(bn)
}

//加入购物车的信息
func Gwcmess(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`)
	// fmt.Println(uid)
	userid, _ := strconv.Atoi(uid)
	var goodsid = r.FormValue(`gid`)
	goodis, _ := strconv.Atoi(goodsid)
	var num = r.FormValue(`count`)
	nums, _ := strconv.Atoi(num)
	var addre = r.FormValue(`address`)
	var imgs = r.FormValue(`imgs`)
	var gname = r.FormValue(`gname`)
	var gmsgs = r.FormValue(`gmsgs`)
	var gprice = r.FormValue(`gprice`)
	if nums == 0 {
		w.Write([]byte(`未选择数量`))
		return
	}
	// fmt.Println(3)
	ok := first.InsertGWC(goodis, userid, nums, addre, imgs, gname, gmsgs, gprice)
	if ok {
		w.Write([]byte(`插入成功`))
		return
	}
	w.Write([]byte(`插入失败`))
}

//singleLittleShop信息
func Singlelittleshop(w http.ResponseWriter, r *http.Request) {

	mod, err := first.SingleLittleShop()
	if err != nil {
		w.Write([]byte(`信息错误延时`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

//singleBigShop
func SingleBShop(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var id = r.FormValue(`kids`)
	// fmt.Println(id)
	kid, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte(`转换类型错误`))
		return
	}
	mod, err := first.SingleBigShop(kid)
	if err != nil {
		w.Write([]byte(`信息错误延时`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func Lunsingleimg(w http.ResponseWriter, r *http.Request) {
	var id = r.FormValue(`kid`)
	kids, _ := strconv.Atoi(id)

	mod, err := first.LunSingleImg(kids)
	if err != nil {
		w.Write([]byte(`信息错误延时`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func Goodsmsgs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	kids, err := strconv.Atoi(r.Form.Get(`kid`))
	// fmt.Println(id)
	if err != nil {
		w.Write([]byte(`转换类型错误`))
		return
	}
	mod, err := first.GoodsMSG(kids)
	if err != nil {
		w.Write([]byte(`信息错误延时`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}
func Inseraddress(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var addr = r.FormValue(`address`)
	if addr == ".." {
		w.Write([]byte(`地址不能为空`))
		return
	}
	var uid = r.FormValue(`id`)
	userid, _ := strconv.Atoi(uid)
	var godsid = r.Form.Get(`goodsid`)
	goodsid, _ := strconv.Atoi(godsid)
	var nums = r.Form.Get(`count`)
	num, _ := strconv.Atoi(nums)
	if num == 0 {
		w.Write([]byte(`未购买数量`))
		return
	}
	// fmt.Println(addr)
	ok := first.Inseradd(userid, goodsid, addr, num)
	w.Header().Set(`Content-Type`, `application/json`)
	if ok {
		w.Write([]byte(`地址错误`))
		return
	} else {
		w.Write([]byte(`地址正确`))
	}
}

//公益扶贫
func viewGYFP(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/gyfp.html`)
	w.Write([]byte(buf))
}

func GyImg(w http.ResponseWriter, r *http.Request) {
	mod, err := gyfpmsg.GyfpImg()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func GyMsg(w http.ResponseWriter, r *http.Request) {
	mod, err := gyfpmsg.GyMsgs()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func Gycountry(w http.ResponseWriter, r *http.Request) {
	mod, err := gyfpmsg.GyCountry()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func Gystory(w http.ResponseWriter, r *http.Request) {
	mod, err := gyfpmsg.GyTextmsg()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func viewCar(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/car.html`)
	w.Write(buf)
}

//将物品显示在购物车栏目中
func Lookgods(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`)
	// fmt.Println(uid)
	// fmt.Println("sdhbfu")
	userid, _ := strconv.Atoi(uid)
	mod, err := first.Showgwc(userid)
	// fmt.Println(456123)
	// fmt.Println(err)
	if err != nil {
		w.Write([]byte(`数据显示失败`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func usermsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`)
	userid, _ := strconv.Atoi(uid)
	mod, err := first.ShowUserMsg(userid)
	if err != nil {
		w.Write([]byte(`信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

//从购物车中移除相应商的逻辑操作
func Remove(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var id = r.Form.Get(`kid`)
	kid, err := strconv.Atoi(id)
	if err != nil {
		buf := util.NewResult(300, "删除的数据有错误，请重试")
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
		return
	}
	mod := first.RemoveGod(kid)
	if mod {
		buf := util.NewResult(200, "删除成功")
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	} else {
		buf := util.NewResult(300, "删除失败")
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
	}
}

//云游天下
func Viewyuntx(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/cloudTravel.html`)
	w.Write(buf)
}

//关于我们相关的逻辑函数处理
func ViewAboutUs(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/personalCenter.html`)
	w.Write(buf)
}

// 展示相应的新闻信息
func PerInfo(w http.ResponseWriter, r *http.Request) {
	mod, err := first.AllInfor()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

// 公司动态信息
func MoveMsgs(w http.ResponseWriter, r *http.Request) {
	mod, err := first.Movemsg()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

//2019年展示信息
func MoveNine(w http.ResponseWriter, r *http.Request) {
	mod, err := first.Movenine()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

//2018年展示信息
func MoveEight(w http.ResponseWriter, r *http.Request) {
	mod, err := first.Moveeight()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

//2017年展示信息
func MoveSeven(w http.ResponseWriter, r *http.Request) {
	mod, err := first.Moveseven()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

//2016年展示信息
func MoveSix(w http.ResponseWriter, r *http.Request) {
	mod, err := first.Movesix()
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)

}

//用户反馈的信息
func FedBackMsg(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	var name = r.FormValue(`name`)
	var email = r.FormValue(`emails`)
	var tel = r.FormValue(`tel`)
	var msgs = r.FormValue(`msgs`)
	ok := first.FeedMsgs(name, email, tel, msgs)
	if ok {
		w.Write([]byte(`提交成功`))
	} else {
		w.Write([]byte(`提交成功`))
	}
}

func ViewSimble(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/client.html`)
	w.Write(buf)
}

//用户注册的逻辑操作
func Uregister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var tel = r.FormValue(`tel`)
	// fmt.Println(tel)
	var pass = r.FormValue(`pass`)
	var againpass = r.FormValue(`againpass`)
	var tagname = r.FormValue(`tagname`)
	var name = r.FormValue(`name`)
	var sex = r.FormValue(`sex`)
	var idcard = r.FormValue(`idcard`)
	var email = r.FormValue(`email`)
	var imgs = r.FormValue(`imgs`)
	if tel == "" || len(tel) != 11 {
		w.Write([]byte(`电话号码有误`))
		return
	}
	if pass != againpass {
		w.Write([]byte(`密码不一致`))
		return
	}
	mods := check.UserRegis(tel, pass, tagname, name, sex, idcard, email, imgs)
	if mods != nil {
		w.Write([]byte(`添加失败`))
	} else {
		w.Write([]byte(`添加成功`))
	}
}

////个人中心用户添加收货地址到数据库
func InsertAddres(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`) // 用户ID
	userid, _ := strconv.Atoi(uid)
	// fmt.Println(userid)
	var address = r.FormValue(`addres`)
	// fmt.Println(address)
	var detaaddress = r.FormValue(`detaaddress`)
	var youbian = r.FormValue(`youbian`)
	youb, _ := strconv.Atoi(youbian)
	var name = r.FormValue(`names`)
	var tel = r.FormValue(`tel`)

	err := check.InserAddr(userid, address, detaaddress, youb, name, tel)
	if err != nil {
		w.Write([]byte(`添加失败`))
	} else {
		w.Write([]byte(`添加成功`))
	}
}

//显示用户添加得到收货地址
func ShowUAddre(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`)
	userid, _ := strconv.Atoi(uid)
	mod, err := check.ShowAddre(userid)
	if err != nil {
		w.Write([]byte(`查询失败`))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	w.Write(md)
}

//用户删除相应的地址信息
func RemoveAddre(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var id = r.FormValue(`kid`)
	kid, _ := strconv.Atoi(id)
	err := check.Removeaddr(kid)
	if err != nil {
		w.Write([]byte(`操作失败`))
		return
	}
	w.Write([]byte(`移除成功`))
}

//显示用户昵称
func ShowUserMsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`)
	userid, _ := strconv.Atoi(uid)
	mod, err := check.ShowUsermsg(userid)
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	w.Header().Set(`Content-Type`, `application/json`)
	md, _ := json.Marshal(mod)
	// fmt.Println(md)

	w.Write(md)
}

//显示用户所有信息
func ShowUserMsgAll(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`)
	userid, _ := strconv.Atoi(uid)
	mod, err := check.ShowUsermsgall(userid)
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	md, _ := json.Marshal(mod)
	// fmt.Println(md)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

//用户添加个人信息
func InsertUsermsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uid = r.FormValue(`id`)
	// fmt.Println(uid)
	id, _ := strconv.Atoi(uid)
	// fmt.Println(id)
	var tagname = r.FormValue(`tagname`)
	var name = r.FormValue(`name`)
	var sex = r.FormValue(`sex`)
	var idcard = r.FormValue(`idcard`)
	var email = r.FormValue(`email`)
	err := check.InsertMsg(id, tagname, name, sex, idcard, email)
	if err != nil {
		w.Write([]byte(`信息错误`))
		return
	}
	w.Write([]byte(`添加成功`))
}

//上传头像
var Userid int

func InsertUserImg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Access-Control-Allow-Origin`, `*`)
	r.ParseForm()
	var uid = r.Form.Get(`id`)
	Userid, _ = strconv.Atoi(uid)

}

//上传用户头像
func Upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println(Userid)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Header().Set(`Access-Control-Allow-Origin`, `*`)
	//	fmt.Println(r.Method)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	r.ParseMultipartForm(1 << 32)
	_, h, _ := r.FormFile("file")
	// fmt.Println(err)
	ext := path.Ext(h.Filename)
	dir := `static/img/singleuser` + time.Now().Format("2006-01-02") + "/" //yyyy-MM-rr
	os.MkdirAll(dir, 0666)
	f, _ := h.Open()
	name := util.RandStr() + ext
	var da = "/" + dir + name

	fmt.Println(da)
	err := check.UpUserimg(da, Userid)
	if err != nil {
		w.Write([]byte(`上传失败`))
		return
	}
	w.Write([]byte(`上传成功`))
	f1, _ := os.Create(dir + name)
	io.Copy(f1, f)
	f.Close()
	f1.Close()
	w.Write([]byte("/" + dir + name))
}

//查看生成账号页面
func ViewNumber(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`view/number.html`)
	w.Write(buf)
}

func RandNum(w http.ResponseWriter, r *http.Request) {

	nums := util.Randone() + util.Rand() + util.Randone() + util.Randone() + util.Randone()
	//	num, _ := strconv.Atoi(nums)
	r.ParseForm()
	var tel = r.FormValue(`tel`)
	//	fmt.Println("tel:" + tel)
	//	num, _ := strconv.Atoi(nums)
	//	fmt.Println(nums)
	err := check.InsertUserNum(nums, tel)
	if err != nil {
		w.Write([]byte(`失败`))
		return
	}
	w.Write([]byte(nums))
}

//中间件
func mid(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		token := r.Form.Get("token")
		if token == "" {
			w.Write(util.NewResult(301, "没有token,请重试"))
			return
		}
		d1 := &check.Jwt{}
		j, err := jwt.ParseWithClaims(token, d1, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenKey), nil
		})
		// j, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(tokenKey), nil
		// })

		if err != nil {
			w.Write(util.NewResult(302, "token 非法,请重试"))
			return
		}
		// 合法
		if j.Valid {
			// fmt.Println(d1)

			// var id = j.Claims.(jwt.MapClaims)["Id"]
			// buf, _ := json.Marshal(d1)
			var id = d1.Id
			// fmt.Println(id)
			buf, _ := json.Marshal(id)
			// str := string(buf)
			// fmt.Println(str)
			r.Form.Add("id", string(buf))
			next.ServeHTTP(w, r)
		} else { // 不合法
			w.Write(util.NewResult(303, "不合法,请重试"))
			return
		}
	})
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// http.Handle(`/check`, mid(http.HandlerFunc(Viewcheck)))
	http.HandleFunc(`/check`, Viewcheck)
	http.HandleFunc(`/register1`, control.VeiwRegister1)
	http.HandleFunc(`/api/user/regis`, Uregister) //用户注册表

	// http.HandleFunc(`/register`, ViewLogin)
	http.HandleFunc(`/`, ViewMain)
	// http.HandleFunc(`/viewmain`, ViewMain)
	http.HandleFunc(`/yushangjs`, Viewysjs)
	// http.HandleFunc(`/regis`, RegUser)
	// http.HandleFunc(`/regis/email`, CheckEmail)
	http.HandleFunc(`/checknp`, Checklogo)
	http.HandleFunc(`/lostpass`, Lostpa)
	http.HandleFunc(`/viewlost`, ViewLost)
	http.HandleFunc(`/api/user/list1`, Infor1)
	http.HandleFunc(`/api/user/list2`, Infor2)
	http.HandleFunc(`/api/user/list3`, Infor3)
	http.HandleFunc(`/api/user/produ`, Titls)
	http.HandleFunc(`/api/user/mess`, Indexmess)
	http.HandleFunc(`/api/user/navs`, Navx)
	http.HandleFunc(`/api/user/vide`, Vide)
	http.HandleFunc(`/api/user/lunims`, YunsjsImg)
	http.HandleFunc(`/api/user/gysjs`, Goodmes)
	http.HandleFunc(`/api/user/spyjsj`, Yunsproduct)
	http.HandleFunc(`/api/user/spyjsj2`, Yunsproduct2)

	http.HandleFunc(`/viewNumber`, ViewNumber) //生成账号页面
	//	http.Handle(`/viewNumber`, mid(http.HandlerFunc(ViewNumber)))
	http.HandleFunc(`/api/user/nums`, RandNum) //生成账号
	//	http.Handle(`/api/user/nums`, mid(http.HandlerFunc(RandNum)))

	http.HandleFunc(`/api/user/frined`, Frined)
	// http.HandleFunc(`/api/single/sgoods`, Gwcmess)
	http.Handle(`/api/single/sgoods`, mid(http.HandlerFunc(Gwcmess))) //加入购物车的信息
	http.HandleFunc(`/single`, ViewSingle)
	http.HandleFunc(`/api/single/litshop`, Singlelittleshop)
	http.HandleFunc(`/api/single/bigshop`, SingleBShop)
	http.HandleFunc(`/api/user/sMsgimg`, Lunsingleimg)
	http.HandleFunc(`/api/user/goodsmsg`, Goodsmsgs)
	http.Handle(`/api/single/address`, mid(http.HandlerFunc(Inseraddress)))
	// http.HandleFunc(`/api/single/address`, Inseraddress)
	http.HandleFunc(`/viewgyfp`, viewGYFP) //公益扶贫
	http.HandleFunc(`/api/gyfp/gyfpimg`, GyImg)
	http.HandleFunc(`/api/gyfp/gyfpmsg`, GyMsg)
	http.HandleFunc(`/api/gyfp/country`, Gycountry)
	http.HandleFunc(`/api/gyfp/textdetial`, Gystory)

	// http.HandleFunc(`/car.html`, viewCar)
	http.Handle(`/car.html`, mid(http.HandlerFunc(viewCar)))
	// http.HandleFunc(`/api/gwc/goods`, Lookgods)
	http.Handle(`/api/gwc/goods`, mid(http.HandlerFunc(Lookgods)))
	// http.HandleFunc(`/api/gwc/user`, usermsg) //购物车中显示用户信息
	http.Handle(`/api/gwc/user`, mid(http.HandlerFunc(usermsg)))
	http.HandleFunc(`/api/car/remove`, Remove)
	// http.Handle(`/api/car/remove`, mid(http.HandlerFunc(Remove)))

	http.HandleFunc(`/yytx.html`, Viewyuntx) //云游天下

	http.HandleFunc(`/About.html`, ViewAboutUs) //关于我们
	http.HandleFunc(`/api/per/msg`, PerInfo)
	http.HandleFunc(`/api/per/movemsg`, MoveMsgs) // 公司动态信息
	http.HandleFunc(`/api/per/movemsgnine`, MoveNine)
	http.HandleFunc(`/api/per/movemsgeight`, MoveEight)
	http.HandleFunc(`/api/per/movemsgseven`, MoveSeven)
	http.HandleFunc(`/api/per/movemsgsix`, MoveSix)
	http.HandleFunc(`/api/per/sends`, FedBackMsg) //用户反馈信息

	//	http.HandleFunc(`/simble`, ViewSimble) // 个人中心页面
	http.Handle(`/simble`, mid(http.HandlerFunc(ViewSimble)))
	//	http.HandleFunc(`/api/client/send`, InsertAddres) //////个人中心用户添加收货地址到数据库
	http.Handle(`/api/client/send`, mid(http.HandlerFunc(InsertAddres)))
	//	http.HandleFunc(`/api/client/showaddr`, ShowUAddre) // 个人中心显示用户添加地址信息
	http.Handle(`/api/client/showaddr`, mid(http.HandlerFunc(ShowUAddre)))
	http.HandleFunc(`/api/client/remove`, RemoveAddre) // 个人中心用户移除相应地址信息
	//	http.Handle(`/api/client/remove`, mid(http.HandlerFunc(RemoveAddre)))
	http.Handle(`/api/client/usermsg`, mid(http.HandlerFunc(ShowUserMsg)))        //显示用户的昵称
	http.Handle(`/api/client/umsgshowall`, mid(http.HandlerFunc(ShowUserMsgAll))) //显示用户所有个人信息
	//	http.HandleFunc(`/api/client/msg`, InsertUsermsg) //用户添加个人信息
	http.Handle(`/api/client/msg`, mid(http.HandlerFunc(InsertUsermsg)))
	http.HandleFunc(`/upload`, Upload) //上传头像
	//	http.Handle(`/upload`, mid(http.HandlerFunc(Upload)))
	http.Handle(`/api/uploaders`, mid(http.HandlerFunc(InsertUserImg)))

	// http.HandleFunc(`/api/name`, ShowUserName) //个人中心显示用户信息
	// http.Handle(`/api/name`, mid(http.HandlerFunc(ShowUserName)))

	//	http.ListenAndServe(":8080", nil)

	// fmt.Println("run on 8080")
	http.ListenAndServeTLS(":4320", "cert-1542427206238_www.linchongyang.cn.crt", "cert-1542427206238_www.linchongyang.cn.key", nil)
	// fmt.Println(err)
}
