window.onload = function() {
    // import qs from 'qs'
    // var id = this.$route.params.id
    var id = location.search.replace('?id=', "")


    var nsp = new Vue({
        el: '#tall',
        data: {
            littleshop: [], //侧栏小信息
            navcs: [],
            bigshop: [], //大图详情
            friendlin: [],
            sMsgimg: [], //头部展示轮播信息
            goodsmsgs: [], //产品价格详情
            checkadd: '', //收获地址
            checkadd1: '',
            checkadd2: '',
        },
        methods: {
            goodsMsg: function() {
                axios.get('/api/user/goodsmsg', {
                    params: { kid: id }
                }).then(function(resp) {
                    // console.log(id)
                    // console.log(resp)
                    nsp.goodsmsgs = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            nowbuy: function() {
                //添加用户地址信息
                // var qs = require('qs')
                axios.get('/api/single/address', {
                    params: {
                        address: nsp.checkadd + '.' + nsp.checkadd1 + '.' + nsp.checkadd2,
                        goodsid: id,
                    }
                }).then(function(resp) {
                    // console.log(resp)
                }).catch(function(error) {
                    console.log(error)
                })
            },
            sMsgimgs: function() {
                axios({
                    method: 'get',
                    url: '/api/user/sMsgimg',
                    params: {
                        kid: id
                    }
                }).then(function(res) {
                    // console.log(res)
                    nsp.sMsgimg = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            frien: function() {
                axios({
                    method: 'get',
                    url: '/api/user/frined'
                }).then(function(res) {
                    // console.log(res)
                    nsp.friendlin = res.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            bigsh: function() {
                axios.get('/api/single/bigshop', {
                    params: { kids: id }
                }).then(function(resp) {
                    // console.log(id)
                    // console.log(resp)
                    nsp.bigshop = resp.data
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
                    nsp.navcs = res.data

                }).catch(function(error) {
                    console.log(error)
                })
            },
            send1: function() {
                axios.get('/api/single/litshop', {
                    params: {
                        kid: id
                    }
                }).then(function(resp) {
                    // console.log(id)
                    // console.log(resp)
                    nsp.littleshop = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },

        },
        mounted: function() {
            //生命周期函数
            this.send1();
            this.navs();
            this.bigsh();
            this.frien();
            this.sMsgimgs();
            this.goodsMsg();
            this.nowbuy();
            // console.log(this.$route.params.id)
        },

    })
}