window.onload = function() {
    var flag1 = 0,
        flag2 = 0;
    //获取验证码
    var varcode;
    var strcode;
    //选取验证码警告
    var codeWarn = document.getElementsByClassName("codeWarnHint")[0];
    //获取类名为code的元素
    var code = document.getElementsByClassName('varcode')[0];
    //选取验证码提示
    var codeHint = document.getElementsByClassName("codeHint")[0];
    //字符串
    var str = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
        //生成随机数并返回
    function getRandom() {
        return Math.round(Math.random() * 61);
    }
    //生成字符串并放入元素内
    function getCode() {
        strcode = "";
        var num;
        for (var i = 0; i < 4; i++) {
            num = getRandom();
            strcode += str[num];
        }
        code.innerHTML = strcode;
    }
    //getCode()函数执行
    getCode();
    //点击执行函数getCode()更换验证码
    code.onclick = function() {
        getCode();
    }




    var nm = new Vue({
        el: '#box',
        data: {
            tel: '',
            pass: '',
            againpass: '',
            tagname: '填写昵称',
            name: '填写姓名',
            sex: '男',
            idcard: '填写身份证',
            email: '填写邮箱',
            img: '/static/img/kapian/bg4.jpg',
            codes: '', //验证码
        },
        methods: {
            send: function() {
                // console.log(this.tel)
                // console.log(this.pass)
                var params = new URLSearchParams();
                params.append('tel', this.tel);
                params.append('pass', this.pass);
                params.append('againpass', this.againpass)
                params.append('tagname', this.tagname);
                params.append('name', this.name);
                params.append('sex', this.sex);
                params.append('idcard', this.idcard);
                params.append('email', this.email);
                params.append(`imgs`, this.img)
                    // params.append(`token`, localStorage.getItem("token"))
                axios.post('/api/user/regis', params).then(function(resp) {
                    console.log(nm.codes)
                    strcode = strcode.toUpperCase();

                    ns = nm.codes.toUpperCase();
                    console.log(ns)
                        // if (nm.codes != strcode) {
                        //     alert("验证码错误")
                        //     return
                        // }
                        // console.log(resp.data)
                    if (nm.tel == "" || nm.tel.length != 11) {
                        alert("电话号码有误")
                        return
                    }

                    if (nm.pass != nm.againpass) {
                        alert("密码不一致")
                    } else if (ns != strcode) {
                        alert("验证码错误")
                        return
                    } else if (resp.data == '添加成功') {
                        // alert("添加成功")
                        window.location.href = "/viewNumber?tel=" + nm.tel
                    } else {
                        alert("该手机号已被注册")
                    }
                })
            }
        },

    })

    //选取所有的input
    var input = document.getElementsByTagName("input");
    //选取所有的提示
    var hint = document.getElementsByClassName("hint");
    //选取所有的警告
    var warn = document.getElementsByClassName("warnHint");



    //密码长度验证

    input[1].onblur = function() {
        if (input[1].value.length < 6 || input[1].value.length > 15) {
            hint[1].style.display = "none";
            warn[1].style.display = "block";
            flag1 = 0;
        } else {
            hint[1].style.display = "block";
            warn[1].style.display = "none";
            flag1 = 1;
        }
    }

    input[2].onblur = function() {
        if (input[1].value != input[2].value) {
            hint[2].style.display = "none";
            warn[2].style.display = "block";
            flag2 = 0;
        } else {
            hint[2].style.display = "block";
            warn[2].style.display = "none";
            flag2 = 1;
        }
    }

    //对比验证码
    input[4].onclick = function() {
        strcode = strcode.toUpperCase();
        varcode = input[3].value;
        varcode = varcode.toUpperCase();
        if (flag1 == 0 || flag2 == 0) {
            alert("密码有误");
        }
    }

}