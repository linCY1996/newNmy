window.onload = function() {
    new Vue({
        el: '#logonp',
        data: {

        },
        methods: {
            nec: function() {
                axios.get('/car.html', {
                    params: {
                        token: localStorage.getItem("token")
                    }
                }).then(function(resp) {
                    console.log(resp.data.code)
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
        }
    })
}