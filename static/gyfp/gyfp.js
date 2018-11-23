window.onload = function() {
    var gy = new Vue({
        el: '#tall',
        data: {
            navcs: [],
            gyfpmsg: [], //公益扶贫上部信息
            country: [], //乡村信息
            textdetial: [], //农牧云的故事
        },
        methods: {
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
            gyfpmsgs: function() {
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
        },
        mounted: function() {
            this.navs();
            this.gyfpmsgs();
            this.gycountry();
            this.gytextdetial();
        }

    })
}