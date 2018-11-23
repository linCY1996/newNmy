package control

import (
	"NMY/check"
	"NMY/first"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//核对用户登录时的用户名和密码是否一致
func Checklogo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var uName = r.FormValue(`name`)
	var uPass = r.FormValue(`mima`)

	mods, err := check.CheckUser(uName, uPass)
	dd, _ := json.Marshal(mods)
	if err != nil {
		w.Write([]byte(`信息错误`))
		return
	}
	if mods.Name != uName {
		w.Write([]byte(dd))
		return
	}
	if mods.Pass != uPass {
		w.Write([]byte(dd))
		return
	}

	w.Write([]byte(dd))

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
	var userid = r.FormValue(`userid`)
	useris, _ := strconv.Atoi(userid)
	var goodsid = r.FormValue(`goodsid`)
	goodis, _ := strconv.Atoi(goodsid)

	_, err := first.InsertGWC(useris, goodis)
	if err != nil {
		w.Write([]byte(`没有响应的信息`))
		return
	}
	w.Write([]byte(`插入成功`))
}

//singleLittleShop信息
func Singlelittleshop(w http.ResponseWriter, r *http.Request) {
	mod, err := first.SingleLitteShop()
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
	mod, err := first.SingleBigShop()
	if err != nil {
		w.Write([]byte(`信息错误延时`))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(md)
}

func Lunsingleimg(w http.ResponseWriter, r *http.Request) {
	mod, err := first.LunSingleImg()
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
	var id = r.Form.Get(`kid`)
	kids, err := strconv.Atoi(id)
	fmt.Println(id)
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
	var addr = r.FormValue(`address`)
	ok := first.Inseradd(addr)
	if ok {
		w.Write([]byte(`地址错误`))
		return
	} else {
		w.Write([]byte(`地址正确`))
	}
}
