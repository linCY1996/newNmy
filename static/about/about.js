window.onload = function() {
    var ab = new Vue({
        el: '#tall',
        data: {
            aboutmsglef: '',
            aboutmsg1: [],
            aboutfri: [], //团队成员
        },
        methods: {
            msgleft: function() {
                axios.get('/api/about/aboutmsgleft').then(function(resp) {
                    // console.log(resp)
                    ab.aboutmsglef = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            msg: function() {
                axios.get('/api/about/aboutmsg1').then(function(resp) {
                    // console.log(resp)
                    ab.aboutmsg1 = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
            aboutfris: function() {
                axios.get('/api/about/aboutfri').then(function(resp) {
                    console.log(resp)
                    ab.aboutfri = resp.data
                }).catch(function(error) {
                    console.log(error)
                })
            },
        },
        mounted: function() {
            this.msg(); //关于我们
            this.msgleft(); //关于我们左边信息
            this.aboutfris(); //团队成员
        }
    })
}