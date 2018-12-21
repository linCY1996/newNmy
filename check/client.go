package check

//用户收货地址信息
type UserAddre struct {
	ID          int    `json:"id,omitempty" xml:"id,omitempty"`
	Userid      int    `json:"userid,omitempty" xml:"userid,omitempty"`
	Address     string `json:"address,omitempty" xml:"address,omitempty"`
	Detaaddress string `json:"detaaddress,omitempty" xml:"detaaddress,omitempty"`
	Youbian     int    `json:"youbian,omitempty" xml:"youbian,omitempty"`
	Names       string `json:"names,omitempty" xml:"names,omitempty"`
	Tel         string `json:"tel,omitempty" xml:"tel,omitempty"`
}

//用户添加收货地址到数据库
func InserAddr(userid int, address, detaaddress string, youbian int, name, tel string) error {
	_, err := dB.Exec(`insert into dizhi(userid, address, detaaddress, youbian, names, tel) values(?,?,?,?,?,?)`, userid, address, detaaddress, youbian, name, tel)
	return err
}

//显示用户添加的地址信息
func ShowAddre(userid int) ([]UserAddre, error) {
	mod := make([]UserAddre, 0)
	err := dB.Select(&mod, `select * from dizhi where userid=?`, userid)
	return mod, err
}

//用户移除相应收货地址
func Removeaddr(id int) error {
	_, err := dB.Exec(`delete from dizhi where id=?`, id)
	return err
}
