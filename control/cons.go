package control

import (
	"io/ioutil"
	"net/http"
)

func VeiwRegister1(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("view/register1.html")
	w.Write(buf)
}
