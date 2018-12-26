package util

var mp map[int]string

func init() {
	mp = make(map[int]string, 20)
}

//保存或者更新登录信息
func Set(id int, token string) {
	mp[id] = token
}

//获取登录信息
func Get(id int) (string, bool) {
	v, ok := mp[id]
	return v, ok
}
