window.onload = function() {
    new Vue({
        el: "#logonp",
        data: {
            href: "car.html?token=" + localStorage.getItem("token")
        }
    })
}