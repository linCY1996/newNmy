var down = document.getElementsByClassName('down')[0];
var bp = document.getElementsByClassName("word");
var sp = document.getElementsByClassName("sp");
//var bg=document.getElementsByClassName("bg");
var input = document.getElementsByTagName("input");


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