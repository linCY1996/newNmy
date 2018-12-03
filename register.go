package main

import (
	"NMY/check"
	"NMY/first"
	"NMY/gyfpmsg"
	"NMY/util"
	"encoding/json"
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
	var uName = r.FormValue(`name`)
	// var uPass = r.FormValue(`mima`)
	pass := r.Form.Get("mima")
	mods, err := check.CheckUser(uName)
	if err != nil {
		w.Write([]byte(`信息错误`))
		return
	}
	if uName != mods.Name {
		w.Write([]byte(`信息错误`))
		return
	}

	if mods.Pass != pass {
		w.Write([]byte(`信息错误`))
		return
	}
	data := check.Jwt{
		Id:   mods.ID,
		Name: mods.Name,
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

//将用户的注册信息提交到数据库
func RegUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name = r.FormValue(`name`)
	var idcard = r.FormValue(`idcard`)
	var email = r.FormValue(`email`)
	var address = r.FormValue(`address`)
	var tel = r.FormValue(`tel`)
	var newname = r.FormValue(`new_name`)
	var pass = r.FormValue(`pass`)
	var againpass = r.FormValue(`again_pass`)

	mod, err := check.Useradd(name, idcard, email, address, tel, newname, pass, againpass)
	if err != nil {
		w.Write([]byte(`注册失败`))
		return
	}
	if mod.Email == email {
		w.Write([]byte(`该邮箱已经被注册过，请重新注册新的邮箱`))
		return
	}
	// dd, _ := json.Marshal(mod)
	// w.Header().Set(`Content-Type`, `application/json`)
	w.Write([]byte(`注册成功`))
}

func CheckEmail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var email = r.FormValue(`email`)
	mods, err := check.CheckEmail(email)
	if err != nil {
		w.Write([]byte(`查询信息错误`))
		return
	}
	dd, _ := json.Marshal(mods)
	w.Header().Set(`Content-Ttpe`, `application/json`)
	w.Write(dd)
}

func Lostpa(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var mail = r.FormValue(`email`)
	var pass = r.FormValue(`pass1`)
	var againpass = r.FormValue(`againpass1`)

	lop, _ := check.Lostpass(mail)
	if lop.Email != mail {
		w.Write([]byte(`此邮箱未被注册`))
	} else {
		_, err := check.UpDataPass(pass, againpass, mail)
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
	r.ParseForm()
	var id = r.Form.Get(`kid`)
	kid, err := strconv.Atoi(id)
	if err != nil {
		w.Write([]byte(`类型转换错误`))
		return
	}
	// fmt.Println(id)
	mod, err := first.SingleLittleShop(kid)
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

//关于我们页面信息
// func ViewAbout(w http.ResponseWriter, r *http.Request) {
// 	buf, _ := ioutil.ReadFile(`view/nmyabout.html`)
// 	w.Write(buf)
// }

// func AboutMymsg(w http.ResponseWriter, r *http.Request) {
// 	mod, err := abouts.Aboutmymsg()
// 	if err != nil {
// 		w.Write([]byte(`查询信息错误`))
// 		return
// 	}
// 	md, _ := json.Marshal(mod)
// 	w.Header().Set(`Content-Type`, `application/json`)
// 	w.Write(md)
// }

// func AboutMy(w http.ResponseWriter, r *http.Request) {
// 	mod, err := abouts.Aboutmy()
// 	if err != nil {
// 		w.Write([]byte(`查询信息错误`))
// 		return
// 	}
// 	md, _ := json.Marshal(mod)
// 	w.Header().Set(`Content-Type`, `application/json`)
// 	w.Write(md)
// }

// func Aboutfriend(w http.ResponseWriter, r *http.Request) {
// 	mod, err := abouts.Aboutfr()
// 	if err != nil {
// 		w.Write([]byte(`查询信息错误`))
// 		return
// 	}
// 	md, _ := json.Marshal(mod)
// 	w.Header().Set(`Content-Type`, `application/json`)
// 	w.Write(md)
// }

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
			buf, _ := json.Marshal(id)
			// str := string(buf)
			// // fmt.Println(str)
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
	http.HandleFunc(`/register`, ViewLogin)
	http.HandleFunc(`/check`, Viewcheck)
	http.HandleFunc(`/`, ViewMain)
	// http.HandleFunc(`/viewmain`, ViewMain)
	http.HandleFunc(`/yushangjs`, Viewysjs)
	http.HandleFunc(`/regis`, RegUser)
	http.HandleFunc(`/regis/email`, CheckEmail)
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

	// http.HandleFunc(`/viewabout`, ViewAbout) //关于我们
	// http.HandleFunc(`/api/about/aboutmsgleft`, AboutMymsg)
	// http.HandleFunc(`/api/about/aboutmsg1`, AboutMy)
	// http.HandleFunc(`/api/about/aboutfri`, Aboutfriend)

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

	// http.ListenAndServe(":8080", nil)
	// fmt.Println("run on 8080")
	http.ListenAndServeTLS(":4320", "cert-1542427206238_www.linchongyang.cn.crt", "cert-1542427206238_www.linchongyang.cn.key", nil)
	// fmt.Println(err)
}
