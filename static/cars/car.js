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
                axios({
                    method: 'POST',
                    url: '/api/car/remove',
                    headers: { 'content-type': 'application/x-www-form-urlencoded' },
                    params: {
                        kid: id
                    }
                }).then(function(resp) {
                    // console.log(resp)
                    console.log(id)
                    if (resp == "删除成功") {
                        alert("删除成功")
                    } else {
                        location.reload();
                    }
                }).catch(function(error) {
                    console.log(error)
                })
            },
            totalPrice: function() {
                var totalP = 0;
                for (var i = 0, len = nm.gods.length; i < len; i++) {
                    totalP += nm.gods[i].price * nm.gods[i].num;
                }
                return totalP;
            },
            Users: function() {
                axios.get('/api/gwc/user', {
                    params: { token: localStorage.getItem("token") } //用户的id
                }).then(function(resp) {
                    // console.log(resp)
                    nm.users = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },

        },
        mounted: function() {
            this.gds();
            this.navs();
            this.Users(); //用户的信息展示
            this.herfs();

        },
        // watch: {
        //     totalP: function() {
        //         this.totalPrice()
        //     }
        // }

    })
}