window.onload = function() {
    new Vue({
        el: '#logonp',
        data: {

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
                    // console.log("sadasd")
            },
        }
    })
}