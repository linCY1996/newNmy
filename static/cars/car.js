window.onload = function() {
    // $(".margA1").click(function(e) {
    //     console.log(e)
    // })

    var nm = new Vue({
        el: '#tall',
        data: {
            navcs: [],
            users: [], //用户的信息
            gods: [], //货物的信息
        },
        methods: {
            gds: function() {
                axios({
                    method: 'get',
                    url: '/api/gwc/goods',
                    params: {
                        token: localStorage.getItem("token")
                    }
                }).then(function(resp) {
                    // console.log(resp)
                    nm.gods = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            totalPrice: function() {
                var totalP = 0;
                for (var i = 0, len = this.gods.length; i < len; i++) {
                    totalP += nm.gods[i].price * nm.gods[i].num;
                }
                return totalP;
            },
            navs: function() {
                axios({
                    method: 'get',
                    url: '/api/user/navs',
                }).then(function(res) {
                    // console.log(res)
                    nm.navcs = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            remove: function(id) {
                var that = this
                var params = new URLSearchParams()
                params.append('kid', id)
                axios.post('/api/car/remove', params, {
                        headers: {
                            'content-type': 'application/x-www-form-urlencoded'
                        }
                    }).then(function(resp) {
                        console.log(id)
                        console.log(resp.data.msg)
                        if (resp.data.msg == "删除成功") {
                            // nm.gods = resp.data
                            that.gds()
                                // alert("删除成功")
                        }
                    })
                    // axios({
                    //     method: 'POST',
                    //     url: '',
                    //     headers: {  },
                    //     params: {
                    //         kid: id
                    //     }
                    // }).then(function(resp) {
                    //     // console.log(resp)

                // }).catch(function(error) {
                //     console.log(error)
                // })
            },
            Users: function() {
                axios.get('/api/gwc/user', {
                    params: { token: localStorage.getItem("token") } //用户的id
                }).then(function(resp) {
                    // console.log(resp)
                    nm.users = resp.data
                })
            },
            singles: function() {
                axios.get('/simble', {
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
        },
        mounted: function() {
            this.gds();
            this.navs();
            this.Users(); //用户的信息展示
            // this.herfs();

        },
        // watch: {
        //     totalP: function() {
        //         this.totalPrice()
        //     }
        // }

    })

    $("#btn").click(function() {
        console.log("123")
        layui.use('layer', function() {
            var layer = layui.layer;
            // layer.open({

            //     content: '本网站暂无实际商品，还无法立即购买'
            // });
            var text = '本网站暂无实际商品，还无法立即购买'
            layer.open({
                skin: 'btn',
                type: 1,
                offset: 'auto' //具体配置参考：http://www.layui.com/doc/modules/layer.html#offset
                    ,
                // id: 'layerDemo' + type //防止重复弹出
                // ,
                area: ['500px', '400px'],
                content: '<div style="width:500px;height:285px;background:black;color:white;">' + text + '</div>',
                btn: '关闭全部',
                btnAlign: 'c' //按钮居中
                    ,
                shade: 0 //不显示遮罩
                    ,
                time: 2000,
                yes: function() {
                    layer.closeAll();
                }
            });
        });
    })
}