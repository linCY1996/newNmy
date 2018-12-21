var can = document.getElementById("canvas");
var ctx = can.getContext("2d");
var w = can.width = window.innerWidth;
var h = can.height = window.innerHeight; 
var count = 30; 
var drops = []; 

window.onresize = function() {
	w = can.width = window.innerWidth;
	h = can.height = window.innerHeight;
}

function Drop() {} 

Drop.prototype = {
	init: function() {
		this.x = random(0, w);
		this.y = 0;
		this.r = 1;
		this.vy = random(4, 5); 
		this.vr = 1; 
		this.a = 1;
		this.va = 0.96; 
		this.l = random(h * 0.8, h * 0.9);
	},
	draw: function() { //绘制
		if(this.y > this.l) {
			ctx.beginPath(); 
			ctx.arc(this.x, this.y, this.r, 0, 2 * Math.PI, false); 
			ctx.strokeStyle = "rgba(230,230,230," + this.a + ")";
			ctx.stroke(); 
		} else {
			ctx.fillStyle = color(this.a); 
			ctx.fillRect(this.x, this.y, 2, 10, ); 
		}
		this.update();
	},
	update: function() {
		if(this.y < this.l) {
			this.y += this.vy;
		} else {
			if(this.a > 0.03) {
				this.r += this.vr;
				if(this.r > 50) {
					this.a *= this.va;
				}
			} else {
				this.init();
			}
		}
	}
}
function move() {
	ctx.fillStyle = "rgba(255,255,255,.1)";
	ctx.fillRect(0, 0, w, h);
	for(var i = 0; i < drops.length; i++) {
		drops[i].draw();
	}
	requestAnimationFrame(move);
}
function setup() {
	for(var i = 0; i < count; i++) {
		(function(j) {
			setTimeout(function() {
				var drop = new Drop();
				drop.init();
				drops.push(drop);
			}, j * 200);
		}(i))
	}
}
setup();
move();
function random(min, max) {
	return Math.random() * (max - min) + min;
}
function color(a) {
	var r = Math.floor(Math.random() * 255 + 150);
	return "rgba(" + r + "," + r + "," + r + "," + a + ")"
}