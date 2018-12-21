window.onload = function() {

    var nm = new Vue({
        el: '#tall',
        data: {
            vides: [],
            navcs: [],
            messg: [],
            titl: [],
            item: [],
            item1: [],
            item2: [],
            friendlin: [],
            href: "car.html?token=" + localStorage.getItem("token")
                // goodid: nm.shoplist1.id
        },
        methods: {
            singles: function() {
                axios.get('/car.html', {
                    params: {
                        token: localStorage.getItem("token")
                    }
                }).then(function(resp) {
                    if (resp.data.code == 301 || resp.data.code == 302 || resp.data.code == 303) {
                        window.location.href = ("/check")
                    } else if (resp.data.code == 302) {
                        window.location.href = ("/check")
                            // window.location.href = "/single?id=" + id
                            // return   

                    } else {
                        window.location.href = "simble?token=" + localStorage.getItem("token")
                    }
                })
            },
            nec: function() {
                axios.get('/car.html', {
                    params: {
                        token: localStorage.getItem("token")
                    }
                }).then(function(resp) {
                    if (resp.data.code == 301 || resp.data.code == 302 || resp.data.code == 303) {
                        window.location.href = ("/check")
                    } else if (resp.data.code == 302) {
                        window.location.href = ("/check")
                            // window.location.href = "/single?id=" + id
                            // return   

                    } else {
                        window.location.href = "car.html?token=" + localStorage.getItem("token")
                    }
                })
            },
            tui: function() {
                localStorage.removeItem("token")
                window.location.href = ("/check")
                console.log("sadasd")
            },
            frien: function() {
                axios({
                    method: 'get',
                    url: '/api/user/frined'
                }).then(function(res) {
                    // console.log(res)
                    nm.friendlin = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },


            vide: function() {
                axios({
                    method: 'get',
                    url: '/api/user/vide'
                }).then(function(res) {
                    // console.log(res)
                    nm.vides = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            navs: function() {
                axios({
                    method: 'get',
                    url: '/api/user/navs'
                }).then(function(res) {
                    // console.log(res)
                    nm.navcs = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            mess: function() {
                axios({
                    method: 'get',
                    url: '/api/user/mess'
                }).then(function(res) {
                    // console.log(res)
                    nm.messg = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            titlc: function() {
                axios({
                    method: 'get',
                    url: '/api/user/produ'
                }).then(function(res) {
                    console.log(res)
                    nm.titl = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            send: function() {
                axios({
                    method: 'get',
                    url: '/api/user/list1'
                }).then(function(res) {
                    // console.log(res)
                    nm.item = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            send1: function() {
                axios({
                    method: 'get',
                    url: '/api/user/list2'
                }).then(function(res) {
                    // console.log(res)
                    nm.item1 = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            send2: function() {
                axios({
                    method: 'get',
                    url: '/api/user/list3'
                }).then(function(res) {
                    // console.log(res)
                    nm.item2 = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            //点击加入购物车
            // nowsend: function(e) {
            //     console.log(e)
            //     axios({
            //         method: 'post',
            //         url: '/api/user/gwc',
            //         data: {
            //             goodsid: e
            //         }
            //     }).then(function(resp) {
            //         console.log(resp)
            //     }).catch(function(error) {
            //         console.log(error)
            //     })
            // },


        },
        mounted: function() {
            this.vide(); //生命周期
            this.navs();
            this.mess();
            this.titlc();
            this.send();
            this.send1();
            this.send2();
            this.frien(); //友情链接

        },
        components: {
            'my-message': { //组件
                template: '<b>优选项目推荐<span>RECOMMENDED</span></b>'
            },
            'my-mess': {
                template: '<p>产品过关/质量保证/优质精选</p>'
            }
        }
    });
    $(".nav-body-text1-ch1 li").mouseover(function() {
        //通过 .index()方法获取元素下标，从0开始，赋值给某个变量
        var index = $(this).index();
        //让内容框的第 _index 个显示出来，其他的被隐藏
        $(".tab-box>div").eq(index).show().siblings().hide();
        //改变选中时候的选项框的样式，移除其他几个选项的样式
        $(this).addClass("change").siblings().removeClass("change");
        // $(this).addClass("change").show();
    });
    $(".nav-body-text1-ch2 li").mouseover(function() {
        //通过 .index()方法获取元素下标，从0开始，赋值给某个变量
        var index = $(this).index();
        //让内容框的第 _index 个显示出来，其他的被隐藏
        $(".tab-box2>div").eq(index).show().siblings().hide();
        //改变选中时候的选项框的样式，移除其他几个选项的样式
        $(this).addClass("change").siblings().removeClass("change");
        // $(this).addClass("change").show();
    });

    // //轮播图部分
    // var index = 0;
    // $(".btn").click(function() {
    //     index = $(this).index() //获取点击元素的下标值
    //     $(this).addClass("active").siblings().removeClass("active");
    //     $(".img-item").fadeOut();
    //     $(".img-item").eq(index).fadeIn();
    // });

    // $(".right").click(function() {
    //     index++; //一次加一
    //     if (index > 3) {
    //         index = 0;
    //     }
    //     $(".img-wrap .img-item").fadeOut();
    //     $(".img-wrap .img-item").eq(index).fadeIn();
    //     // alert(index);
    //     $(".btn").eq(index).addClass("active").siblings().removeClass("active");
    // });
    // $(".left").click(function() {
    //     index--; //一次减一
    //     if (index < 0) {
    //         index = 3;
    //     }
    //     $(".img-wrap .img-item").fadeOut();
    //     $(".img-wrap .img-item").eq(index).fadeIn();
    //     $(".btn").eq(index).addClass("active").siblings().removeClass("active");
    // });

}