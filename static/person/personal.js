window.onload = function() {

    var nm = new Vue({
        el: '#subject',
        data: {
            msgs: [], //平台特色
            moves: [], //公司动态信息
            movenine: [], //2019年信息
            moveeight: [], //2018年信息
            moveseven: [], //2017年信息
            movesix: [], //2016年信息
        },
        methods: {
            Msgs: function() {
                axios.get('/api/per/msg').then(function(resp) {
                    // console.log("asdfsf")
                    // console.log(resp)
                    nm.msgs = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },

            moveMsgs: function() {
                axios.get('/api/per/movemsg').then(function(resp) {
                    // console.log("asdfsf")
                    // console.log(resp)
                    nm.moves = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            moveNine: function() {
                axios.get('/api/per/movemsgnine').then(function(resp) {
                    // console.log("asdfsf")
                    // console.log(resp)
                    nm.movenine = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            moveEight: function() {
                axios.get('/api/per/movemsgeight').then(function(resp) {
                    // console.log("asdfsf")
                    // console.log(resp)
                    nm.moveeight = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            moveSeven: function() {
                axios.get('/api/per/movemsgseven').then(function(resp) {
                    // console.log("asdfsf")
                    // console.log(resp)
                    nm.moveseven = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            moveSix: function() {
                axios.get('/api/per/movemsgsix').then(function(resp) {
                    // console.log("asdfsf")
                    // console.log(resp)
                    nm.movesix = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            btns: function() {
                var name = $("#inp1").val()
                var emails = $("#inp2").val()
                var tel = $("#inp3").val()
                var msgs = $("#inp4").val()

                $.post('/api/per/sends', {
                    name: name,
                    emails: emails,
                    tel: tel,
                    msgs: msgs
                }, function(resp) {
                    alert(resp)
                        // console.log(resp)
                        // console.log("asds")
                })
            },
        },
        mounted: function() {
            this.Msgs();
            this.moveMsgs();
            this.moveNine();
            this.moveEight();
            this.moveSeven();
            this.moveSix();
        }
    })

    var down = document.getElementsByClassName('down')[0];
    var bp = document.getElementsByClassName("word");
    var sp = document.getElementsByClassName("sp");
    //var bg=document.getElementsByClassName("bg");
    var input = document.getElementsByTagName("input");
    var turnimg = document.getElementsByClassName("turnimg");
    var imag = document.getElementsByClassName("imag1");
    //公司动态开始
    for (var i = 0; i < turnimg.length; i++) {
        (function(n) {
            turnimg[n].onmousedown = function() {
                for (var j = 0; j < imag.length; j++) {
                    imag[j].style.display = "none";
                }
                imag[n].style.display = "block";
            }
        }(i))
    }

    down.onclick = function() {
        window.scrollTo(0, 982);
    }

    for (var i = 0; i < sp.length; i++) {
        (function(n) {
            sp[n].onmouseenter = function() {
                for (var j = 0; j < bp.length; j++) {
                    bp[j].style.display = "none";
                    sp[j].style.border = "none"
                }
                this.style.border = "2px solid red";
                bp[n].style.display = "block";
            }
            sp[n].onmouseleave = function() {
                this.style.border = "none";
            }
        }(i))
    }
    for (var i = 0; i < input.length - 1; i++) {
        input[i].style.color = "#999"
        switch (i) {
            case 0:
                (function(n) {
                    input[n].onfocus = function() {
                        if (input[n].value == "姓名") {
                            input[n].value = "";
                            input[n].style.color = "black";
                        }
                    }
                    input[n].onblur = function() {
                        if (input[n].value == "") {
                            input[n].value = "姓名";
                            input[n].style.color = "#999";
                        }
                    }
                }(i))
                break;
            case 1:
                (function(n) {
                    input[n].onfocus = function() {
                        if (input[n].value == "邮箱地址") {
                            input[n].value = "";
                            input[n].style.color = "black";
                        }
                    }
                    input[n].onblur = function() {
                        if (input[n].value == "") {
                            input[n].value = "邮箱地址";
                            input[n].style.color = "#999";
                        }
                    }
                }(i))
                break;
            case 2:
                (function(n) {
                    input[n].onfocus = function() {
                        if (input[n].value == "联系方式") {
                            input[n].value = "";
                            input[n].style.color = "black";
                        }
                    }
                    input[n].onblur = function() {
                        if (input[n].value == "") {
                            input[n].value = "联系方式";
                            input[n].style.color = "#999";
                        }
                    }
                }(i))
                break;
            case 3:
                (function(n) {
                    input[n].onfocus = function() {
                        if (input[n].value == "留言信息填写...") {
                            input[n].value = "";
                            input[n].style.color = "black";
                        }
                    }
                    input[n].onblur = function() {
                        if (input[n].value == "") {
                            input[n].value = "留言信息填写...";
                            input[n].style.color = "#999";
                        }
                    }
                }(i))
                break;
        }
    }
}