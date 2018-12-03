window.onload = function() {
    // new Vue({
    //     el: '#logonp',
    //     data: {
    //         hrr: "car.html?token=" + localStorage.getItem("token")
    //     }
    // })
    var ys = new Vue({
        el: '#tall',
        data: {
            num: 0,
            navcs: [],
            lunimg: [],
            goodrecom: [], //优质精选推荐
            shoplist1: [],
            shoplist2: [],
            hrr: "car.html?token=" + localStorage.getItem("token")
        },
        methods: {
            nec: function() {
                axios.get('/car.html', {
                    params: {
                        token: localStorage.getItem("token")
                    }
                }).then(function(resp) {
                    if (resp.data.code == 301) {
                        window.location.href = ("/check")
                    } else {
                        window.location.href = "/car.html?token=" + localStorage.getItem("token")
                    }
                })
            },
            tui: function() {
                localStorage.removeItem("token")
                window.location.href = ("/check")
                    // console.log("sadasd")
            },
            navs: function() {
                axios({
                    method: 'get',
                    url: '/api/user/navs'
                }).then(function(res) {
                    // console.log(res)
                    ys.navcs = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            goodre: function() {
                axios({
                    method: 'get',
                    url: '/api/user/gysjs',
                }).then(function(res) {
                    // console.log(res)
                    // console.log(id)
                    ys.goodrecom = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            shops: function() {
                axios({
                    method: 'get',
                    url: '/api/user/spyjsj'
                }).then(function(res) {
                    // console.log(res)
                    ys.shoplist1 = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            shops2: function() {
                axios({
                    method: 'get',
                    url: '/api/user/spyjsj2'
                }).then(function(res) {
                    console.log(res)
                    ys.shoplist2 = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            lunim: function() {
                axios({
                    method: 'get',
                    url: '/api/user/lunims'
                }).then(function(res) {
                    // console.log(res)
                    ys.lunimg = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            //轮播图实现部分
            autoplay: function() {
                var that = this
                that.num++;
                if (that.num == that.lunimg.length) {
                    that.num = 0
                }
            },
            play: function() {
                setInterval(this.autoplay, 2000)
            },
        },
        mounted: function() {
            this.navs();
            this.goodre(); //云上集市优质精选
            this.shops();
            this.shops2();
            this.play();
            this.lunim();
        }
    })
}