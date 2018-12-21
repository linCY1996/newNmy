window.onload = function() {

    var gy = new Vue({
        el: '#tall',
        data: {
            num: 0,
            navcs: [],
            gyfpimg: [], //公益扶贫上部信息
            gyfpmsg: '',
            country: [], //乡村信息
            textdetial: [], //农牧云的故事
        },
        methods: {
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
            navs: function() {
                axios({
                    method: 'get',
                    url: '/api/user/navs'
                }).then(function(res) {
                    // console.log(res)
                    gy.navcs = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            Gyfpimg: function() {
                axios.get('/api/gyfp/gyfpimg', ).then(function(resp) {
                    // console.log(resp)
                    gy.gyfpimg = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            Gyfpmsg: function() {
                axios.get('/api/gyfp/gyfpmsg', ).then(function(resp) {
                    // console.log(resp)
                    gy.gyfpmsg = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            gycountry: function() {
                axios.get('/api/gyfp/country').then(function(resp) {
                    // console.log(resp)
                    gy.country = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            gytextdetial: function() {
                axios.get('/api/gyfp/textdetial').then(function(resp) {
                    // console.log(resp)
                    gy.textdetial = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            //轮播图实现部分
            autoplay: function() {
                var that = this
                that.num++;
                if (that.num == that.gyfpimg.length) {
                    that.num = 0
                }
            },
            play: function() {
                setInterval(this.autoplay, 2000)
            },
        },
        mounted: function() {
            this.play();
            this.navs();
            this.Gyfpimg();
            this.Gyfpmsg();
            this.gycountry();
            this.gytextdetial();
        }

    })
}