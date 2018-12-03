window.onload = function() {

    var nm = new IDBCursorWithValue({
        el: '#tall',
        data: {
            navcs: [],
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
        },
        mounted: function() {
            this.navs();
        }
    })
    $(function() {
        $(".gulou-cooperate0").mouseover(function() {
            $("#gulou-cooperate0-name").css("display", "block");
            $("#gulou-cooperate0-img").css({
                "transition": "display linear 1s",
                "display": "block",
            });
        }).mouseleave(function() {
            $("#gulou-cooperate0-name").css("display", "none")
            $("#gulou-cooperate0-img").css({
                "display": "none",
            });
        });
        $(".gulou-cooperate0").click(function() {
            $(".travel-list").css("display", "block")
        })
    });
    $(window).bind("scroll", function() {
        var top = $(this).scrollTop(); // 当前窗口的滚动距离
        // console.log($(this))
        if (top > 1) {
            $(".nav").css({
                "background": "white",
                "opacity": "1",
                "transition": "background-color 1s"
            });
            $(".nav-table a").css("color", "black");
            $(".company a").css("color", "black");
            $(".logo span").css("color", "black");
            $(".am-input-group").css("opacity", "1")
            $(".title").attr("src", "img/标题-黑.png")
        }
        if (top == 0) {
            $(".nav").css({
                "background": "none",
                "opacity": "1",
            });
            $(".nav-table a").css("color", "white");
            $(".company a").css("color", "white");
            $(".logo span").css("color", "white");
            $(".am-input-group").css("opacity", "0.5")
            $(".title").attr("src", "img/标题.png")
        }
    });
    // <!--地图 -->

    require.config({
        paths: {
            echarts: 'assets/echarts-2.2.7/build/dist'
        }
    });
    require(
        [
            'echarts',
            'echarts/chart/map' // 使用柱状图就加载bar模块，按需加载
        ],
        function(ec) {

            // --- 地图 ---
            var myChart = ec.init(document.getElementById('main'));
            myChart.setOption({
                tooltip: {
                    trigger: 'item',
                    formatter: '{b}'
                },
                series: [{
                    name: '中国',
                    type: 'map',
                    mapType: 'china',
                    // tooltip: '1123',
                    selectedMode: 'single',
                    itemStyle: {
                        normal: {
                            areaStyle: {
                                color: 'rgba(255,255,255,0.8)',
                            },
                            borderColor: 'grey',
                            label: {
                                show: true
                            }
                        },
                        emphasis: {
                            areaStyle: {
                                color: 'rgba(255,130,71)',

                            },
                            label: {
                                show: true
                            }
                        }
                    },
                    data: [{
                        name: '四川',
                        selected: true
                    }]
                }]
            });
            var ecConfig = require('echarts/config');
            myChart.on(ecConfig.EVENT.MAP_SELECTED, function(param) {
                var selected = param.selected;
                var provence = param.target; //获取所选定的省份
                console.log(param.target)
                document.getElementById("provence").style.display = "block";
                mapDisplay(provence);
                chinaDisapper();
            });
        }
    );

    function chinaDisapper() {
        var t = document.getElementById("main");
        t.style.display = "none"
    };

    function mapDisplay(provence) {
        var myChart = echarts.init(document.getElementById('provence'));
        var option = {
            tooltip: {
                trigger: 'item',
                formatter: function(a) { //鼠标移到某个州市上弹出的提示内容。包括显示样式可以自定义，利用return返回样式即可。
                    return a[1] + ":" + a[2]; //a[1]:州市名称，a[2]:data中的valuez值。
                }
            },
            // dataRange: {
            // 	min: 0,
            // 	max: 10,
            // 	color: ['red', 'rgb(220,220,220)'],
            // 	//color: ['orange', 'blue'],
            // 	boder: 3,
            // 	text: ['10', '0'], // 文本，默认为数值文本
            // 	calculable: true
            // },
            series: [{
                //name: '数据名称',
                type: 'map',
                mapType: provence, //如果是其他省份，也可以改变，例如：上海，北京，天津等地。
                selectedMode: 'single',
                itemStyle: {
                    normal: {
                        areaStyle: {
                            color: 'rgba(255,255,255,0.8)',
                        },
                        label: {
                            show: true
                        },
                    },
                    emphasis: {
                        areaStyle: {
                            color: 'rgba(255,130,71)',

                        },
                        label: {
                            show: true
                        }
                    }
                },
                //data一定要有，不然没有图
                data: [{
                    name: '阿坝藏族羌族自治州',
                    value: 0
                }, {
                    name: '巴中市',
                    value: 0
                }, {
                    name: '成都市',
                    value: 0
                }, {
                    name: '达州市',
                    value: 0
                }, {
                    name: '德阳市',
                    value: 0
                }, {
                    name: '甘孜藏族自治州',
                    value: 0
                }, {
                    name: '广安市',
                    value: 0
                }, {
                    name: '广元市',
                    value: 0
                }, {
                    name: '乐山市',
                    value: 0
                }, {
                    name: '凉山彝族自治州',
                    value: 0
                }, {
                    name: '泸州市',
                    value: 0
                }, {
                    name: '眉山市',
                    value: 0
                }, {
                    name: '绵阳市',
                    value: 0
                }, {
                    name: '内江市',
                    value: 0
                }, {
                    name: '南充市',
                    value: 3,
                    selected: true
                }, {
                    name: '攀枝花市',
                    value: 0
                }, {
                    name: '遂宁市',
                    value: 0
                }, {
                    name: '雅安市',
                    value: 0
                }, {
                    name: '宜宾市',
                    value: 0
                }, {
                    name: '资阳市',
                    value: 0
                }, {
                    name: '自贡市',
                    value: 0
                }]
            }]
        };
        myChart.setOption(option);
        var ecConfig = require('echarts/config');
        myChart.on(ecConfig.EVENT.MAP_SELECTED, function(param) {
            var selected = param.selected;
            var provence = param.target;
            console.log(param.target)
            document.getElementById("provence").style.display = "none";
            document.getElementById("city").style.display = "block";
        });

    }
}