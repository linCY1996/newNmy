/*
Navicat MariaDB Data Transfer

Source Server         : new
Source Server Version : 100311
Source Host           : localhost:3306
Source Database       : nmy

Target Server Type    : MariaDB
Target Server Version : 100311
File Encoding         : 65001

Date: 2018-12-15 20:35:12
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for aboutfr
-- ----------------------------
DROP TABLE IF EXISTS `aboutfr`;
CREATE TABLE `aboutfr` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `mess` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of aboutfr
-- ----------------------------
INSERT INTO `aboutfr` VALUES ('1', '/static/img/about/z.jpg', '曾梓涵', '三季度放牛班覅而那我也恢复IE未婚夫IE热闹可烦你');
INSERT INTO `aboutfr` VALUES ('2', '/static/img/about/z.jpg', '曾梓涵', '三季度放牛班覅而那我也恢复IE未婚夫IE热闹可烦你');
INSERT INTO `aboutfr` VALUES ('3', '/static/img/about/z.jpg', '曾梓涵', '三季度放牛班覅而那我也恢复IE未婚夫IE热闹可烦你');
INSERT INTO `aboutfr` VALUES ('4', '/static/img/about/z.jpg', '曾梓涵', '三季度放牛班覅而那我也恢复IE未婚夫IE热闹可烦你');

-- ----------------------------
-- Table structure for aboutmsg
-- ----------------------------
DROP TABLE IF EXISTS `aboutmsg`;
CREATE TABLE `aboutmsg` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `msg1` varchar(255) DEFAULT NULL,
  `msg2` varchar(255) DEFAULT NULL,
  `litemsg` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of aboutmsg
-- ----------------------------
INSERT INTO `aboutmsg` VALUES ('1', '基于农产品平台研发，以扶贫为初心，致力于产品上行，从生产源头把握农产品品控，保证农产品绿色生态。农牧云与具有一定养殖规模的贫困村的合作社、产业大户实现战略合作，消费者通过认购投资与农户共享收益、共担风险。', '通过宣传推广农特产品，开设贫困销售专栏，为村子全方位分析、打造，对其农产品以及环境资源进行最大化利用，打造“一村一品”，实现为农户增收脱贫。', '【创新模式】“互联网”+农村合作社+消费者');
INSERT INTO `aboutmsg` VALUES ('2', 'h', 'dfg', '【政府监督】政府监督、食品安全有保障');
INSERT INTO `aboutmsg` VALUES ('3', 'g', 'dfg', '【生态养殖】发展绿色农业、生产绿色作物');
INSERT INTO `aboutmsg` VALUES ('4', 'dfg', 'gdf', '【高校帮扶】高校根据具体问题进行对口帮扶');
INSERT INTO `aboutmsg` VALUES ('5', 'dfg', 'dfg', '【四个字词】你来想，你想不到了，我再想');
INSERT INTO `aboutmsg` VALUES ('6', 'dfg', 'dfg', '【可爱的我】帅气的你，苦逼的程序员的生活');

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `userid` int(128) DEFAULT NULL,
  `goodsid` int(128) DEFAULT NULL,
  `addr` varchar(255) DEFAULT '',
  `num` int(120) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=75 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES ('61', '3', '1', '四川省..', null);
INSERT INTO `address` VALUES ('62', '3', '1', '四川省.南充市.顺庆区', null);
INSERT INTO `address` VALUES ('63', '4', '1', '四川省.南充市.顺庆区', null);
INSERT INTO `address` VALUES ('64', '5', '2', '北京.北京.东城区', '0');
INSERT INTO `address` VALUES ('65', '2', '2', '北京.北京.东城区', '1');
INSERT INTO `address` VALUES ('66', '2', '2', '北京.北京.东城区', '1');
INSERT INTO `address` VALUES ('67', '2', '2', '北京.北京.东城区', '1');
INSERT INTO `address` VALUES ('68', null, '2', '上海.上海.静安区', '2');
INSERT INTO `address` VALUES ('69', null, '2', '四川.成都.青白江区', '4');
INSERT INTO `address` VALUES ('70', null, '2', '上海.上海.浦东新区', '4');
INSERT INTO `address` VALUES ('71', '28', '14', '北京.北京.东城区', '1');
INSERT INTO `address` VALUES ('72', '28', '14', '北京.北京.东城区', '1');
INSERT INTO `address` VALUES ('73', '28', '13', '北京.北京.东城区', '1');
INSERT INTO `address` VALUES ('74', '28', '13', '北京.北京.东城区', '1');

-- ----------------------------
-- Table structure for dizhi
-- ----------------------------
DROP TABLE IF EXISTS `dizhi`;
CREATE TABLE `dizhi` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `userid` int(128) DEFAULT NULL,
  `address` varchar(50) DEFAULT NULL,
  `detaaddress` varchar(50) DEFAULT NULL,
  `youbian` int(30) DEFAULT NULL,
  `names` varchar(30) DEFAULT NULL,
  `tel` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of dizhi
-- ----------------------------
INSERT INTO `dizhi` VALUES ('5', '2', '广东.广州.越秀区', '南红石', '611630', '张三', '13350058238');

-- ----------------------------
-- Table structure for feedback
-- ----------------------------
DROP TABLE IF EXISTS `feedback`;
CREATE TABLE `feedback` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL,
  `tel` varchar(30) DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of feedback
-- ----------------------------
INSERT INTO `feedback` VALUES ('1', '姓名dgdsfg', '邮箱地址dfg', '联系方式dfg', '留言信息填写.fgd..');
INSERT INTO `feedback` VALUES ('2', '姓名', '邮箱地址', '联系方式', '留言信息填写...');
INSERT INTO `feedback` VALUES ('3', '姓名sd', '邮箱地址sdf', '联系方式sdf', '留言信息填写...sdf');
INSERT INTO `feedback` VALUES ('4', '姓名sddsf', '邮箱地址sdf', '联系方式sdf', '留言信息填写...sdf');
INSERT INTO `feedback` VALUES ('5', '姓名sdfd', '邮箱地sdf址', '联系方式sdf', '留言信息sdf填写...');
INSERT INTO `feedback` VALUES ('6', 'hzangsan ', '465', '45645', '123161');
INSERT INTO `feedback` VALUES ('7', 'asd', '465', '4654', '1313');
INSERT INTO `feedback` VALUES ('8', 'sd', '1465', '132', '13');
INSERT INTO `feedback` VALUES ('9', '46', '6', '23', '23232');
INSERT INTO `feedback` VALUES ('10', 'dfdf', 'sdfsdf', 'fsdfds', 'fsdfsdfads');
INSERT INTO `feedback` VALUES ('11', 'sadfsdf', 'sdfasdf', 'sdfasdf', 'dfsdfs');
INSERT INTO `feedback` VALUES ('12', 'dsfgfd', '23', '23', '23');
INSERT INTO `feedback` VALUES ('13', 'sdsd', '151', '123132', '1321231');
INSERT INTO `feedback` VALUES ('14', 'zhangan ', '456', '1263', '4156');
INSERT INTO `feedback` VALUES ('15', '伤心城市', '叙', '想参照v', '淡出说的');
INSERT INTO `feedback` VALUES ('16', '林崇阳', '1653709489@qq.com', '13350558238', '大加好');

-- ----------------------------
-- Table structure for friendlink
-- ----------------------------
DROP TABLE IF EXISTS `friendlink`;
CREATE TABLE `friendlink` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `names` varchar(30) DEFAULT NULL,
  `hrefs` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of friendlink
-- ----------------------------
INSERT INTO `friendlink` VALUES ('1', '新农网', 'http://www.baidu.com');
INSERT INTO `friendlink` VALUES ('2', '光农网', 'http://www.baidu.com');
INSERT INTO `friendlink` VALUES ('3', '乐农之家', 'http://www.baidu.com');
INSERT INTO `friendlink` VALUES ('4', '农业大创', 'http://www.baidu.com');
INSERT INTO `friendlink` VALUES ('5', '新农业', 'http://www.baidu.com');

-- ----------------------------
-- Table structure for gooodsmsg
-- ----------------------------
DROP TABLE IF EXISTS `gooodsmsg`;
CREATE TABLE `gooodsmsg` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `goodsname` varchar(255) DEFAULT NULL,
  `goodsmsg` varchar(255) DEFAULT NULL,
  `oldprice` decimal(10,2) DEFAULT NULL,
  `newprice` varchar(25) DEFAULT '',
  `successgoods` int(128) DEFAULT NULL,
  `imgs` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gooodsmsg
-- ----------------------------
INSERT INTO `gooodsmsg` VALUES ('1', '两年正宗散养土鸡老母鸡农家公鸡乌鸡草鸡笨鸡走地鸡', '公鸡冠大而红，性烈好斗，母鸡鸡冠极小。土鸡也叫草鸡、笨鸡，是指从古代家养驯化而成，从未经过任何杂交和优化配种，长期以自然觅食或结合粗饲喂养而成，具有较强的野外觅食和生存能力。具有耐粗饲、就巢性强和抗病力强等特性，肉质鲜美', '56.90', '49.9', '450', '/static/img/1/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('2', '藏香猪', '藏香猪的品质上有“六个最”，即： 肉品中氨基酸含量最高，微量元素最高，脂肪含量最低，猪肠最长，猪皮最薄，鬃毛最长。是藏民的传统民族美食。尤其是猪皮，口味Q弹爽嫩，远异于一般生猪，藏香猪还是我国唯一的放牧型猪种生长在海拔3000——4000米的高原地带，以天然野生可食性植物及果实为主食，成年猪平均体重不足50公斤。', '12.90', '10.5', '123', '/static/img/2/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('3', '农家散养绿壳乌鸡蛋绿壳土鸡蛋30枚包邮', '绿壳蛋鸡是我国江苏泰州(姜堰河横区)、浙江宁波、湖北林区、湖南山区、山东山区、江西林区、河南卢氏山区、云贵高原山区的特产，仅在纯自然放养条件下生长，以虫类和草药为食以保证其产蛋量。相传乾隆年间，村民将黑羽乌骨鸡进贡给乾隆皇帝，乾隆皇帝见黑羽乌骨鸡不仅外貌超群，而且肉质细嫩，奇香味美，视为奇品异珍，列为“贡鸡”，其所产之蛋列为“贡蛋”。', '11.20', '10.8', '21546', '/static/img/3/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('4', '脆霞桃', '云南大理新鲜甜脆冬桃5斤装贡品18个中大果桃蟠桃油桃毛桃包邮\r\n桃子果肉中含有丰富的蛋白质、脂肪、糖、钙、磷、铁和维生素B、维生素C、胡萝卜素、有机酸(主要是苹果酸和柠檬酸)及大量的水分，这些成分能补充人体的营养所需。此外，桃子中含有较多的钾，而钠含量较少，因此适合水肿病人和缺铁性贫血患者食用', '12.50', '10.5', '23', '/static/img/4/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('5', ' 特级鸡棕鸡纵云南野生蘑菇包邮', '现货新鲜野生鸡枞菌500g 特级鸡棕鸡纵云南野生蘑菇包邮 鲜荔枝菌\r\n鸡㙡菌肉厚肥硕，质细丝白，味道鲜甜香脆。含人体所必须的氨基酸、蛋白质、脂肪，还含有各种维生素和钙、磷、核黄酸等物质。鸡枞的吃法很多，可以单料为菜，还能与蔬菜、鱼肉及各种山珍海味搭配，可无论炒、炸、腌、煎、拌、烩、烤、焖，清蒸或做汤，其滋味都很鲜，为菌中之冠', '99.80', '78.5', '45', '/static/img/5/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('6', '陕西眉县 徐香绿心猕猴桃 22个板盒装', '适时地浇灌，保证水分满溢，酸甜一人；太阳直晒，均衡的营养早就营养较丰富；\r\n充足阳光照射，展现魅力姿态，果农悉心呵护每一个果实，不催熟，不打蜡（亲，只要把奇异果和苹果或者是香蕉放置在塑料袋内，几天就可以吃到口感香甜的奇异果啦）', '51.50', '45.8', '1321', '/static/img/6/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('7', '蒲江农家红心柚子当季现摘现发蜜柚新鲜水果不到10斤整箱批发包邮', '美味高山种植，远离城市的喧嚣，一样的红心柚，职位带给您不一样的美味，自然种植的果子虽然表皮不好看，但是慢慢的心意确实不可辜负的。悉心种植，只为让您吃到甜蜜美味的果实。每一里都是颗粒饱满，果肉细嫩无渣，鲜嫩多汁，橙色的果皮光滑西能，晶莹剔透的果肉饱满多汁。', '56.90', '54.1', '456', '/static/img/7/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('8', '现摘现发会理突尼斯软籽石榴当季新鲜水果四川大凉山超甜石榴5斤', '石榴表面的黑点是糖斑，不是坏了哦，石榴没有入过冷库，采摘后自然存储，采摘后自然储存，果果糖分又很足，在快递运输过程中果子会反糖，果果都是自然生长，不打药、不催熟，因为阳光照射的不同，果皮的颜色会有差异，但是不影响口感哦。', '12.10', '10.8', '6', '/static/img/8/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('9', '四川安岳黄柠檬新鲜当季水果批发包邮发6斤 现摘一级果皮薄柠檬', '安岳柠檬主栽品种尤力克，经过80多年的栽培驯化，选优提纯，选育出了高产、优质、抗逆性强的株系。经中国柑桔研究所检测，并与美国加州、意大利西西里柠檬比较，安岳柠檬含柠檬油7.4‰，可溶性固形物9.5%，柠檬酸6.7%，Vc58/100ml，出汁率38%，Vp2.5%，果胶3%，同时富含肌醇、柠檬烯等多种维生素和微量元素，柠檬油含柠檬醛4%，柠檬烯95%，各项指标均超过美国、意大利柠檬，是理想的既能鲜食又能加工的优良品种。\r\n', '11.80', '10.8', '132', '/static/img/9/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('10', '四川攀枝花凯特大芒果应季新鲜水果包邮整箱10斤当季青芒', '攀枝花芒果具有纤维少、味甜芳香、质地腻滑、营养丰富的独特品质。攀枝花芒果在9—10月份成熟，比国内其他芒果产区晚熟1—2个月。攀枝花芒果（凯特）果大，卵圆形，果鼻明显，果顶圆，果肩较平，果蒂直，果皮薄光滑，果粉厚，果点密、小、白色。青果黄绿色，向阳面有红晕，熟果底色黄绿色、盖色紫红色，套袋果实熟果绿黄色。果肉橙黄色、纤维极少，汁多、味甜，质地腻滑，香气浓，几乎无松香味。', '45.70', '43.8', '23', '/static/img/10/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('11', '仟佰果 云南石林人参果5斤包邮人生果长果当季新鲜水果特产孕妇', '通常人们所说的人参果是一种产于我国甘肃武威民勤、酒泉玉门等地区日光温室里面的一种高营养水果。甘肃省武威市民勤县也被称为中国人参果之乡。其果肉清香多汁，腹内无核，风味独特，具有高蛋白、低脂肪、低糖等特点，同时富含蛋白质、氨基酸以及微量元素、维生素与矿物元素，具有保健功效。', '54.90', '49.9', '132', '/static/img/11/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('12', '红心芭乐番石榴新鲜水果特产5斤包邮当季巴乐潘石榴 芭乐', '番石榴味道甘甜多汁，果肉柔滑，果心较少无籽，常吃可以补充人体所缺乏的营养成分，可以强身健体提高身体素质。比起苹果，番石榴所含有的脂肪少38%，卡路里少42%。番石榴还广泛应用于食品加工业，主要目的就是为了增加食品维生素C的含量，使食品的营养得以强化和提高。', '39.90', '35.9', '3', '/static/img/12/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('13', '海南青金桔新鲜小青桔小柠檬水果橘子青桔子青桔果5斤包邮5斤装', '理气，解郁，化痰，止渴，消食，醒酒。金桔能增强机体抗寒能力，可以防治感冒、降低血脂。胸闷郁结，不思饮食，或伤食饱满，醉酒口渴之人食用；适宜急慢性气管炎，肝炎，胆囊炎，高血压，直管硬化者食用。脾弱气虚之人不宜多食，糖尿病人忌食。凡口舌碎痛，齿龈肿痛者忌食。', '32.90', '29.9', '4548', '/static/img/13/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('14', '土豆马铃薯洋芋新鲜农家自种蔬菜农产品陕北特产5斤装包邮', '“红皮土豆”经检测：花青素含量为每百克3.15毫克，淀粉含量16%，每千克“红皮土豆”中含粗淀粉100.7克、还原糖1.23克，每百克“红玉土豆”中含蛋白质2.11克，脂肪0.1克，每百克“红皮土豆”中微量元素及维生素含量为：碳水化合物15.5克，维生素C18.7毫克。红心土豆还含有丰富的维生素B1、B2、B6和泛酸等B群维生素及大量的优质纤维素及碘、硒、锰、钾、锌等微量元素，氨基酸、这些成分在人的肌体抗老防病过程中有着重要的作用', '56.80', '52.9', '4548', '/static/img/14/1.jpg');
INSERT INTO `gooodsmsg` VALUES ('15', '云南大理新鲜紫皮独头蒜农家产品红皮独蒜5斤低价包邮干蒜', '因当地独特的土壤、水质及气温、地势而形成，蒜皮呈紫色，薄而富有弹性，果实脆而白，半透明，辣味正，口感好，富含多种营养成份和微量元素，具有美容、健胃、杀菌功能，加工利用价值大，开发前景广阔', '23.60', '19.8', '23', '/static/img/15/1.jpg');

-- ----------------------------
-- Table structure for gwc
-- ----------------------------
DROP TABLE IF EXISTS `gwc`;
CREATE TABLE `gwc` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `goodsid` int(128) DEFAULT NULL,
  `userid` int(128) DEFAULT NULL,
  `num` int(20) DEFAULT NULL,
  `addr` varchar(255) DEFAULT NULL,
  `imgs` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT '',
  `msgs` varchar(255) DEFAULT '',
  `price` varchar(25) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gwc
-- ----------------------------
INSERT INTO `gwc` VALUES ('16', '2', '3', '3', '广东.广州.荔湾区', '/static/img/images/meat2.jpg', '海鲜拿到手就服你第三方的风格', '所发生的广东人哥哥多个人退热贴热特瑞ter热特瑞托付给的第三方个地方官第三个电饭锅电饭锅第三个放大豆腐干豆腐干发鬼地方', '10');
INSERT INTO `gwc` VALUES ('18', '2', '3', '4', '四川.阿坝藏族羌族自治州.马尔康县', '/static/img/images/meat03.jpg', '海鲜拿到手就服你第三方的风格', '所发生的广东人哥哥多个人退热贴热特瑞ter热特瑞托付给的第三方个地方官第三个电饭锅电饭锅第三个放大豆腐干豆腐干发鬼地方', '10');
INSERT INTO `gwc` VALUES ('20', '8', '56', '6', '上海.上海.浦东新区', '/static/img/images/meat4.jpg', '海鲜拿到手就服你第三方的风格', '所发生的广东人哥哥多个人退热贴热特瑞ter热特瑞托付给的第三方个地方官第三个电饭锅电饭锅第三个放大豆腐干豆腐干发鬼地方', '10');
INSERT INTO `gwc` VALUES ('38', '15', '28', '1', '北京.北京.东城区', '/static/img/15/1.jpg', '云南大理新鲜紫皮独头蒜农家产品红皮独蒜5斤低价包邮干蒜', '因当地独特的土壤、水质及气温、地势而形成，蒜皮呈紫色，薄而富有弹性，果实脆而白，半透明，辣味正，口感好，富含多种营养成份和微量元素，具有美容、健胃、杀菌功能，加工利用价值大，开发前景广阔', '19.8');
INSERT INTO `gwc` VALUES ('39', '15', '28', '1', '北京.北京.东城区', '/static/img/15/1.jpg', '云南大理新鲜紫皮独头蒜农家产品红皮独蒜5斤低价包邮干蒜', '因当地独特的土壤、水质及气温、地势而形成，蒜皮呈紫色，薄而富有弹性，果实脆而白，半透明，辣味正，口感好，富含多种营养成份和微量元素，具有美容、健胃、杀菌功能，加工利用价值大，开发前景广阔', '19.8');
INSERT INTO `gwc` VALUES ('48', '1', '86', '4', '黑龙江.大兴安岭地区.呼玛县', '/static/img/1/1.jpg', '两年正宗散养土鸡老母鸡农家公鸡乌鸡草鸡笨鸡走地鸡', '公鸡冠大而红，性烈好斗，母鸡鸡冠极小。土鸡也叫草鸡、笨鸡，是指从古代家养驯化而成，从未经过任何杂交和优化配种，长期以自然觅食或结合粗饲喂养而成，具有较强的野外觅食和生存能力。具有耐粗饲、就巢性强和抗病力强等特性，肉质鲜美', '49.9');
INSERT INTO `gwc` VALUES ('49', '3', '86', '4', '内蒙古.乌兰察布.四子王旗', '/static/img/3/1.jpg', '农家散养绿壳乌鸡蛋绿壳土鸡蛋30枚包邮', '绿壳蛋鸡是我国江苏泰州(姜堰河横区)、浙江宁波、湖北林区、湖南山区、山东山区、江西林区、河南卢氏山区、云贵高原山区的特产，仅在纯自然放养条件下生长，以虫类和草药为食以保证其产蛋量。相传乾隆年间，村民将黑羽乌骨鸡进贡给乾隆皇帝，乾隆皇帝见黑羽乌骨鸡不仅外貌超群，而且肉质细嫩，奇香味美，视为奇品异珍，列为“贡鸡”，其所产之蛋列为“贡蛋”。', '10.8');
INSERT INTO `gwc` VALUES ('51', '14', '99', '3', '四川.成都.成华区', '/static/img/14/1.jpg', '土豆马铃薯洋芋新鲜农家自种蔬菜农产品陕北特产5斤装包邮', '“红皮土豆”经检测：花青素含量为每百克3.15毫克，淀粉含量16%，每千克“红皮土豆”中含粗淀粉100.7克、还原糖1.23克，每百克“红玉土豆”中含蛋白质2.11克，脂肪0.1克，每百克“红皮土豆”中微量元素及维生素含量为：碳水化合物15.5克，维生素C18.7毫克。红心土豆还含有丰富的维生素B1、B2、B6和泛酸等B群维生素及大量的优质纤维素及碘、硒、锰、钾、锌等微量元素，氨基酸、这些成分在人的肌体抗老防病过程中有着重要的作用', '52.9');
INSERT INTO `gwc` VALUES ('53', '3', '100', '3', '四川.成都.新都区', '/static/img/3/1.jpg', '农家散养绿壳乌鸡蛋绿壳土鸡蛋30枚包邮', '绿壳蛋鸡是我国江苏泰州(姜堰河横区)、浙江宁波、湖北林区、湖南山区、山东山区、江西林区、河南卢氏山区、云贵高原山区的特产，仅在纯自然放养条件下生长，以虫类和草药为食以保证其产蛋量。相传乾隆年间，村民将黑羽乌骨鸡进贡给乾隆皇帝，乾隆皇帝见黑羽乌骨鸡不仅外貌超群，而且肉质细嫩，奇香味美，视为奇品异珍，列为“贡鸡”，其所产之蛋列为“贡蛋”。', '10.8');

-- ----------------------------
-- Table structure for gyfpcountry
-- ----------------------------
DROP TABLE IF EXISTS `gyfpcountry`;
CREATE TABLE `gyfpcountry` (
  `id` int(50) NOT NULL AUTO_INCREMENT,
  `countrysidename` varchar(255) DEFAULT NULL,
  `links` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gyfpcountry
-- ----------------------------
INSERT INTO `gyfpcountry` VALUES ('1', '险岩村', 'http://url.cn/5eFfzKI');
INSERT INTO `gyfpcountry` VALUES ('2', '瓦子坪村', 'http://url.cn/5V0bfAO');
INSERT INTO `gyfpcountry` VALUES ('3', '玉兰村', 'http://url.cn/5DXep2e');
INSERT INTO `gyfpcountry` VALUES ('4', '塘湾村', 'http://url.cn/508CmS2');
INSERT INTO `gyfpcountry` VALUES ('5', '桂花桥村', 'http://www.jcshys.com/shsj/shsjhd/18929.html');
INSERT INTO `gyfpcountry` VALUES ('6', '保平桥村', 'http://www.chouduol.com/html/2018/qingchunlizhi_0718/12286.html?from=groupmessage ');
INSERT INTO `gyfpcountry` VALUES ('7', '琳琅村', 'http://www.scportray.com/jiaoyuyingxiang/20180718/369.html?from=groupmessage');
INSERT INTO `gyfpcountry` VALUES ('8', '周家坝村', 'http://www.guojishibao.com/index.php?m=wap&siteid=1&a=show&catid=11&typeid=6&id=16087&from=groupmessage');
INSERT INTO `gyfpcountry` VALUES ('9', '方坝村', 'http://www.bjradio.org.cn/xinwen/guonei/2018/0718/11558.html?from=groupmessage');
INSERT INTO `gyfpcountry` VALUES ('10', '文昌村', 'http://appweb.scpublic.cn/news/wx/detail?newsid=109050&time=1531964873650&from=groupmessage');
INSERT INTO `gyfpcountry` VALUES ('11', '神山村', 'http://www.chouduol.com/html/2018/qingchunlizhi_0718/12273.html?from=groupmessage');
INSERT INTO `gyfpcountry` VALUES ('12', '飞来石村', 'http://www.chinaol.org/jiaokewenwei/20180718/3668.html?from=groupmessage');

-- ----------------------------
-- Table structure for gyfpimg
-- ----------------------------
DROP TABLE IF EXISTS `gyfpimg`;
CREATE TABLE `gyfpimg` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gyfpimg
-- ----------------------------
INSERT INTO `gyfpimg` VALUES ('1', '/static/img/gyfp/1.jpg');
INSERT INTO `gyfpimg` VALUES ('2', '/static/img/gyfp/2.jpg');
INSERT INTO `gyfpimg` VALUES ('3', '/static/img/gyfp/3.jpg');
INSERT INTO `gyfpimg` VALUES ('4', '/static/img/gyfp/4.jpg');
INSERT INTO `gyfpimg` VALUES ('5', '/static/img/gyfp/5.jpg');

-- ----------------------------
-- Table structure for gyfpmsg
-- ----------------------------
DROP TABLE IF EXISTS `gyfpmsg`;
CREATE TABLE `gyfpmsg` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `messtitle` varchar(255) DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gyfpmsg
-- ----------------------------
INSERT INTO `gyfpmsg` VALUES ('1', '/static/img/gyfp/1.jpg', '农牧云三下乡', '按时发违法避而不热不那我背负被违法与给贝尔国内微积分似懂非懂过分的话电饭锅电饭锅按时发违法避而不热不那我背负被违法与给贝尔国内微积分似懂非懂过分的话电饭锅电饭锅按时发违法避而不热不那我背负被违法与给贝尔国内微积分似懂非懂过分的话电饭锅电饭锅按时发违法避而不热不那我背负被违法与给贝尔国内微积分似懂非懂过分的话电饭锅电饭锅按时发违法避而不热不那我背负被违法与给贝尔国内微积分似懂非懂过分的话电饭锅电饭锅按时发违法避而不热不那我背负被违法与给贝尔国内微积分似懂非懂过分的话电饭锅电饭锅');
INSERT INTO `gyfpmsg` VALUES ('2', '/static/img/gyfp/1.jpg', '农牧云三下乡', '按时发违法避而不热不那我背负被违法与给贝尔国内微积分似懂非懂过分的话电饭锅电饭锅');

-- ----------------------------
-- Table structure for gyfpmsgs
-- ----------------------------
DROP TABLE IF EXISTS `gyfpmsgs`;
CREATE TABLE `gyfpmsgs` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `msgs` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gyfpmsgs
-- ----------------------------
INSERT INTO `gyfpmsgs` VALUES ('1', '西华师大学子走进瓦子坪村，实地调研脱贫攻坚工作', '7月18日上午西华师范大学“走进仪陇红色老区，助力乡村电商扶贫”暑假实践队一行来到了四川省南充市仪陇县马鞍镇瓦子坪村进行了实地的调研，学习该村的脱贫经验。村支部书记代旭初热情地接待了实践队员，并向他们讲解村中现在的脱贫的基本情况，并带领团队走访村中农户，实地考察脱贫后的现状，讲解之后村中脱贫攻坚工作的开展。脱贫攻坚的主战场正式干部作风的大考场，聚焦真扶贫，扶真贫，真脱贫，各地尽锐出战，砥砺干部作风，狠抓扶贫实效，确保精准脱贫，如期实践，瓦子坪村中各个干部真正做好脱贫工作，干实事。带领村民脱贫增收。');

-- ----------------------------
-- Table structure for gyfpstory
-- ----------------------------
DROP TABLE IF EXISTS `gyfpstory`;
CREATE TABLE `gyfpstory` (
  `id` int(15) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT '',
  `gytitle` varchar(255) DEFAULT '',
  `gymsg` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gyfpstory
-- ----------------------------
INSERT INTO `gyfpstory` VALUES ('1', '/static/img/gyfp/2.jpg', '农牧云的故事', '经历了昨天的炎热，团队顺利到达了四川省南充市仪陇县马鞍镇。成员们安顿好住宿后，立刻前往了我们此次目的地之一的险岩村。与该村敬振华书记亲切交谈和学习了村里的“扶贫经”，从昨天开始，我们暑期社会实践小分队就开始了为期七天以“走进仪陇红色老区，助力乡村电商扶贫”为主题的赴仪陇社会实践活动。');
INSERT INTO `gyfpstory` VALUES ('2', '/static/img/gyfp/3.jpg', '走访瓦子贫村', '骄阳似火的七月，随着时光的流逝，短暂的7天旅程，我们三下乡活动在大家共同努力下圆满结束，大家带着一份感动与思考起程，想想我们分别的那一幕，想想我们与增光学生共同度过的那点点滴滴，想想我们实践队队员们留下的足迹，想想这些天来我们所做的每一件事——真的为我的人生增添了不少的经验和收获，它也是这片天空下不灭的回忆，这将成为我人生的一笔财富，值得我永远珍藏。');
INSERT INTO `gyfpstory` VALUES ('3', '/static/img/gyfp/4.jpg', '走访调查西充古楼镇', '虽然不是盛产期，但还是看到许多菌桶上冒出了一个个白色的“小伞”，队员们纷纷与之合影，书记向我们介绍：“食用菌产业是我们村脱贫攻坚最重要的产业之一，而这里食用菌质量很高，销路更是不愁。”说到这里，书记脸上更是露出了欣慰的笑容。队员们也没有忘记正事，迅速的采集好视频、图片资料后，大家一起留下了珍贵的合影。\r\n');

-- ----------------------------
-- Table structure for indexmessage
-- ----------------------------
DROP TABLE IF EXISTS `indexmessage`;
CREATE TABLE `indexmessage` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `mess` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of indexmessage
-- ----------------------------
INSERT INTO `indexmessage` VALUES ('1', '/static/img/kapian/bg1.jpg', '政府支持', '西华师范大学最美校园歌声唱出我的自我心飞扬');
INSERT INTO `indexmessage` VALUES ('2', '/static/img/kapian/bg2.jpg', '政府支持', '西华师范大学最美校园歌声唱出我的自我心飞扬');
INSERT INTO `indexmessage` VALUES ('3', '/static/img/kapian/bg3.jpg', '政府支持', '西华师范大学最美校园歌声唱出我的自我心飞扬');

-- ----------------------------
-- Table structure for information2
-- ----------------------------
DROP TABLE IF EXISTS `information2`;
CREATE TABLE `information2` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `links` varchar(255) DEFAULT NULL,
  `times` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of information2
-- ----------------------------
INSERT INTO `information2` VALUES ('1', '四川大学生走进朱德故里 重温红色赤子心', 'http://url.cn/5eFfzKI', '8-15');
INSERT INTO `information2` VALUES ('2', '西华师大学子扬青春风帆 ：助力脱贫攻坚', 'http://url.cn/5V0bfAO', '8-16');
INSERT INTO `information2` VALUES ('3', '西华师范大学学子走进瓦子坪村', 'http://url.cn/5DXep2e', '9-12');
INSERT INTO `information2` VALUES ('4', '大学子生调研险岩村香菇种植基地', 'http://url.cn/508CmS2', '8-16');
INSERT INTO `information2` VALUES ('5', '耀眼红色老区，助力乡村电商', 'http://www.jcshys.com/shsj/shsjhd/18929.html', '11-14');
INSERT INTO `information2` VALUES ('6', '西华师大学子走进瓦子坪村', 'http://www.chouduol.com/html/2018/qingchunlizhi_0718/12286.html?from=groupmessage ', '11-14');
INSERT INTO `information2` VALUES ('7', '实地调研脱贫攻坚工作', 'http://www.scportray.com/jiaoyuyingxiang/20180718/369.html?from=groupmessage', '11-14');
INSERT INTO `information2` VALUES ('8', '西华师大学子深入食用菌基地调研乡村脱贫攻坚项目', 'https://url.cn/5cfQQ3k', '11-14');
INSERT INTO `information2` VALUES ('9', '大学子生调研险岩村香菇种植基地', 'http://url.cn/508CmS2', '11-15');
INSERT INTO `information2` VALUES ('10', '台立法部门：自造军舰步子太大 岛内造船能力跟不上', 'https://mil.news.sina.com.cn/china/2018-11-14/doc-ihmutuec0091971.shtml', '11-13');

-- ----------------------------
-- Table structure for information3
-- ----------------------------
DROP TABLE IF EXISTS `information3`;
CREATE TABLE `information3` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `links` varchar(255) DEFAULT NULL,
  `times` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of information3
-- ----------------------------
INSERT INTO `information3` VALUES ('1', '四川大学生走进朱德故里 重温红色赤子心', 'http://url.cn/5eFfzKI', '8-12');
INSERT INTO `information3` VALUES ('2', '西华师大学子扬青春风帆 ：助力脱贫攻坚', 'http://url.cn/5V0bfAO', '5-12');
INSERT INTO `information3` VALUES ('3', '西华师范大学学子走进瓦子坪村', 'http://url.cn/5DXep2e', '9-12');
INSERT INTO `information3` VALUES ('4', '大学子生调研险岩村香菇种植基地', 'http://url.cn/508CmS2', '5-16');
INSERT INTO `information3` VALUES ('5', '神吐槽：做发型+购新房!杜兰特书豪相约洛杉矶', 'http://sports.sina.com.cn/basketball/nba/2018-11-14/doc-ihnvukfe8465967.shtml', '11-13');
INSERT INTO `information3` VALUES ('6', '末节8分2助书豪爆勇士 教练不瞎就该知道用谁	', 'http://sports.sina.com.cn/basketball/nba/2018-11-14/doc-ihnvukfe8356675.shtml', '11-15');
INSERT INTO `information3` VALUES ('7', '泪奔！火箭防守复苏！只皆因那个男人去了低位	', 'http://sports.sina.com.cn/basketball/nba/2018-11-14/doc-ihmutuec0066833.shtml', '11-15');
INSERT INTO `information3` VALUES ('8', '勇火大战就要来了！ 最强西决却变菜鸡互啄？\r\n	', 'http://sports.sina.com.cn/basketball/nba/2018-11-14/doc-ihmutuec0100577.shtml', '11-12');
INSERT INTO `information3` VALUES ('9', '阿杜对喷格林时竟说要离开！ 赌气还是真心话', 'http://sports.sina.com.cn/basketball/nba/2018-11-14/doc-ihmutuec0088064.shtml', '11-14');
INSERT INTO `information3` VALUES ('10', '御用记者:湖人有超级得分手!但他得离开詹姆斯', 'http://sports.sina.com.cn/basketball/nba/2018-11-14/doc-ihnvukfe6696993.shtml', '11-14');

-- ----------------------------
-- Table structure for informations
-- ----------------------------
DROP TABLE IF EXISTS `informations`;
CREATE TABLE `informations` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `links` varchar(255) DEFAULT NULL,
  `times` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of informations
-- ----------------------------
INSERT INTO `informations` VALUES ('1', '西华师大学子走进瓦子坪村', 'http://www.chouduol.com/html/2018/qingchunlizhi_0718/12286.html?from=groupmessage ', '9-23');
INSERT INTO `informations` VALUES ('2', '实地调研脱贫攻坚工作', 'http://www.scportray.com/jiaoyuyingxiang/20180718/369.html?from=groupmessage', '11-14');
INSERT INTO `informations` VALUES ('3', '西华师大学子深入食用菌基地调研乡村脱贫攻坚项目', 'https://url.cn/5cfQQ3k', '11-01');
INSERT INTO `informations` VALUES ('4', '大学子生调研险岩村香菇种植基地', 'http://url.cn/508CmS2', '11-06');
INSERT INTO `informations` VALUES ('5', '西华师大学子走进瓦子坪村', 'http://www.chouduol.com/html/2018/qingchunlizhi_0718/12286.html?from=groupmessage ', '11-13');
INSERT INTO `informations` VALUES ('6', '实地调研脱贫攻坚工作', 'http://www.scportray.com/jiaoyuyingxiang/20180718/369.html?from=groupmessage', '11-13');
INSERT INTO `informations` VALUES ('7', '西华师大学子深入食用菌基地调研乡村脱贫攻坚项目', 'https://url.cn/5cfQQ3k', '11-13');
INSERT INTO `informations` VALUES ('8', '大学子生调研险岩村香菇种植基地', 'http://url.cn/508CmS2', '11-13');
INSERT INTO `informations` VALUES ('9', '西华师大学子深入食用菌基地调研乡村脱贫攻坚项目', 'http://sports.sina.com.cn/g/pl/2018-11-14/doc-ihnvukfe6726164.shtml', '11-13');
INSERT INTO `informations` VALUES ('10', '西华师大学子走进瓦子坪村', 'http://sports.sina.com.cn/china/national/2018-11-14/doc-ihmutuea9965782.shtml', '11-13');

-- ----------------------------
-- Table structure for infortitle
-- ----------------------------
DROP TABLE IF EXISTS `infortitle`;
CREATE TABLE `infortitle` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `links` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of infortitle
-- ----------------------------
INSERT INTO `infortitle` VALUES ('1', '师大要闻', 'fsdf');
INSERT INTO `infortitle` VALUES ('2', '新闻快讯', 'dfg');
INSERT INTO `infortitle` VALUES ('3', '视频新闻', 'fg');
INSERT INTO `infortitle` VALUES ('4', null, null);

-- ----------------------------
-- Table structure for lunboimg
-- ----------------------------
DROP TABLE IF EXISTS `lunboimg`;
CREATE TABLE `lunboimg` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `hre` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of lunboimg
-- ----------------------------
INSERT INTO `lunboimg` VALUES ('1', '/static/img/images/banner.jpg', 'http://www.baidu.com');
INSERT INTO `lunboimg` VALUES ('2', '/static/img/images/banne1.jpg', 'http://www.baidu.com');
INSERT INTO `lunboimg` VALUES ('3', '/static/img/images/banne2.jpg', 'http://www.baidu.com');
INSERT INTO `lunboimg` VALUES ('4', '/static/img/images/banne3.jpg', 'http://www.baidu.com');

-- ----------------------------
-- Table structure for navs
-- ----------------------------
DROP TABLE IF EXISTS `navs`;
CREATE TABLE `navs` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `hrefs` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of navs
-- ----------------------------
INSERT INTO `navs` VALUES ('1', '首页', 'https://www.linchongyang.cn:4320/');
INSERT INTO `navs` VALUES ('2', '云上集市', 'https://www.linchongyang.cn:4320/yushangjs');
INSERT INTO `navs` VALUES ('3', '云游天下', 'https://www.linchongyang.cn:4320/yytx.html');
INSERT INTO `navs` VALUES ('4', '云公益', 'https://www.linchongyang.cn:4320/viewgyfp');
INSERT INTO `navs` VALUES ('6', '关于我们', 'https://www.linchongyang.cn:4320/About.html');

-- ----------------------------
-- Table structure for perinfor
-- ----------------------------
DROP TABLE IF EXISTS `perinfor`;
CREATE TABLE `perinfor` (
  `id` int(25) NOT NULL AUTO_INCREMENT,
  `msg` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of perinfor
-- ----------------------------
INSERT INTO `perinfor` VALUES ('1', '【创新模式】“互联网”+农村合作社+消费者');
INSERT INTO `perinfor` VALUES ('2', '【政府监督】政府监督、食品安全有保障');
INSERT INTO `perinfor` VALUES ('3', '【生态养殖】发展绿色农业、生产绿色作物');
INSERT INTO `perinfor` VALUES ('4', '【高校帮扶】高校根据具体问题进行对口帮扶');
INSERT INTO `perinfor` VALUES ('5', '【四个字词】你来想，你想不到了，我再想');
INSERT INTO `perinfor` VALUES ('6', '【可爱的我】帅气的你，苦逼的程序员');

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
  `id` int(50) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `message` varchar(255) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `hrefs` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of products
-- ----------------------------
INSERT INTO `products` VALUES ('1', '/static/img/13/1.jpg', '海南青金桔', '29.90', 'https://www.linchongyang.cn:4320/single?id=13');
INSERT INTO `products` VALUES ('2', '/static/img/14/1.jpg', '新鲜农家自种蔬菜', '52.90', 'https://www.linchongyang.cn:4320/single?id=14');
INSERT INTO `products` VALUES ('3', '/static/img/15/1.jpg', '农家产品红皮独蒜', '19.80', 'https://www.linchongyang.cn:4320/single?id=15');

-- ----------------------------
-- Table structure for regis
-- ----------------------------
DROP TABLE IF EXISTS `regis`;
CREATE TABLE `regis` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `tel` varchar(30) DEFAULT '',
  `pass` varchar(30) DEFAULT '',
  `tagname` varchar(30) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `sex` varchar(15) DEFAULT NULL,
  `idcard` varchar(40) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `tel` (`tel`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of regis
-- ----------------------------

-- ----------------------------
-- Table structure for register
-- ----------------------------
DROP TABLE IF EXISTS `register`;
CREATE TABLE `register` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL,
  `idcard` varchar(30) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `tel` varchar(30) DEFAULT NULL,
  `newname` varchar(30) DEFAULT '',
  `pass` varchar(30) DEFAULT NULL,
  `againpass` varchar(30) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ema` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=106 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of register
-- ----------------------------
INSERT INTO `register` VALUES ('98', '刘悦', '510131199608125436', '1653709@qq.com', '南充市', '13350058238', '南岸', '123456789', '123456789');
INSERT INTO `register` VALUES ('99', '林重阳', '510131199608125436', '1653709489@qq.com', '南充市', '13350058238', '张三', '77777777', '77777777');
INSERT INTO `register` VALUES ('100', '张安', '510131199608125436', '16537@qq.com', '南充市', '13350058238', '在干撒', '123456789', '123456789');
INSERT INTO `register` VALUES ('103', '林重阳', '510131199608125436', 'asd@qq.com', 'mam', '13350058238', '收到', '123456789', '123456789');

-- ----------------------------
-- Table structure for singledetial
-- ----------------------------
DROP TABLE IF EXISTS `singledetial`;
CREATE TABLE `singledetial` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `hrefs` varchar(255) DEFAULT NULL,
  `mess` varchar(255) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `ids` varchar(30) DEFAULT NULL,
  `idjs` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of singledetial
-- ----------------------------
INSERT INTO `singledetial` VALUES ('1', '/static/img/1/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('2', '/static/img/1/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('3', '/static/img/1/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('4', '/static/img/1/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('5', '/static/img/1/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('6', '/static/img/2/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('7', '/static/img/2/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('8', '/static/img/2/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('9', '/static/img/2/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('10', '/static/img/2/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('11', '/static/img/3/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('12', '/static/img/3/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('13', '/static/img/3/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('14', '/static/img/3/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('15', '/static/img/3/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('16', '/static/img/4/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('17', '/static/img/4/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('18', '/static/img/4/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('19', '/static/img/4/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('20', '/static/img/4/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('21', '/static/img/5/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('22', '/static/img/5/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('23', '/static/img/5/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('24', '/static/img/5/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('25', '/static/img/5/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('26', '/static/img/6/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('27', '/static/img/6/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('28', '/static/img/6/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('29', '/static/img/6/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('30', '/static/img/6/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('31', '/static/img/7/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('32', '/static/img/7/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('33', '/static/img/7/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('34', '/static/img/7/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('35', '/static/img/7/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('36', '/static/img/8/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('37', '/static/img/8/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('38', '/static/img/8/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('39', '/static/img/8/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('40', '/static/img/8/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('41', '/static/img/9/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('42', '/static/img/9/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('43', '/static/img/9/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('44', '/static/img/9/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('45', '/static/img/9/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('46', '/static/img/10/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('47', '/static/img/10/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('48', '/static/img/10/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('49', '/static/img/10/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('50', '/static/img/10/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('51', '/static/img/11/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('52', '/static/img/11/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('53', '/static/img/11/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('54', '/static/img/11/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('55', '/static/img/11/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('56', '/static/img/12/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('57', '/static/img/12/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('58', '/static/img/12/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('59', '/static/img/12/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('60', '/static/img/12/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('61', '/static/img/13/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('62', '/static/img/13/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('63', '/static/img/13/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('64', '/static/img/13/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('65', '/static/img/13/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('66', '/static/img/14/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('67', '/static/img/14/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('68', '/static/img/14/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('69', '/static/img/14/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('70', '/static/img/14/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('71', '/static/img/15/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('72', '/static/img/15/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('73', '/static/img/15/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('74', '/static/img/15/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('75', '/static/img/15/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('76', '/static/img/16/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('77', '/static/img/16/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('78', '/static/img/16/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('79', '/static/img/16/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('80', '/static/img/16/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('81', '/static/img/17/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('82', '/static/img/17/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('83', '/static/img/17/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('84', '/static/img/17/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('85', '/static/img/17/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('86', '/static/img/18/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('87', '/static/img/18/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('88', '/static/img/18/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('89', '/static/img/18/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('90', '/static/img/18/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('91', '/static/img/19/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('92', '/static/img/19/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('93', '/static/img/19/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('94', '/static/img/19/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('95', '/static/img/19/5.jpg', 'img', '232.00', '5', '#5');
INSERT INTO `singledetial` VALUES ('96', '/static/img/20/1.jpg', 'img', '45.20', '1', '#1');
INSERT INTO `singledetial` VALUES ('97', '/static/img/20/2.jpg', 'img', '54.00', '2', '#2');
INSERT INTO `singledetial` VALUES ('98', '/static/img/20/3.jpg', 'img', '3.00', '3', '#3');
INSERT INTO `singledetial` VALUES ('99', '/static/img/20/4.jpg', 'img', '56.00', '4', '#4');
INSERT INTO `singledetial` VALUES ('100', '/static/img/20/5.jpg', 'img', '232.00', '5', '#5');

-- ----------------------------
-- Table structure for singleimg
-- ----------------------------
DROP TABLE IF EXISTS `singleimg`;
CREATE TABLE `singleimg` (
  `id` int(15) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `imgs` varchar(50) DEFAULT NULL,
  `ids` varchar(30) DEFAULT NULL,
  `idjs` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of singleimg
-- ----------------------------
INSERT INTO `singleimg` VALUES ('1', '海鲜', '/static/img/kapian/bg1.jpg', 'bg1', '#bg1');
INSERT INTO `singleimg` VALUES ('2', '海鲜', '/static/img/kapian/bg2.jpg', 'bg2', '#bg2');
INSERT INTO `singleimg` VALUES ('3', '海鲜', '/static/img/kapian/bg3.jpg', 'bg3', '#bg3');
INSERT INTO `singleimg` VALUES ('4', '海鲜', '/static/img/kapian/bg4.jpg', 'bg4', '#bg4');
INSERT INTO `singleimg` VALUES ('5', '海鲜', '/static/img/kapian/bg5.jpg', 'bg5', '#bg5');

-- ----------------------------
-- Table structure for singleimgmsg
-- ----------------------------
DROP TABLE IF EXISTS `singleimgmsg`;
CREATE TABLE `singleimgmsg` (
  `id` int(128) NOT NULL AUTO_INCREMENT,
  `hrefs` varchar(255) DEFAULT NULL,
  `mess` varchar(255) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of singleimgmsg
-- ----------------------------
INSERT INTO `singleimgmsg` VALUES ('1', '/static/img/1/a.jpg', 'img', '45.20');
INSERT INTO `singleimgmsg` VALUES ('2', '/static/img/1/b.jpg', 'img', '54.00');
INSERT INTO `singleimgmsg` VALUES ('3', '/static/img/1/c.jpg', 'img', '3.00');
INSERT INTO `singleimgmsg` VALUES ('4', '/static/img/1/d.jpg', 'img', '56.00');
INSERT INTO `singleimgmsg` VALUES ('5', '/static/img/1/e.jpg', 'img', '232.00');
INSERT INTO `singleimgmsg` VALUES ('6', '/static/img/2/a.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('7', '/static/img/2/b.jpg', 'img', '56.00');
INSERT INTO `singleimgmsg` VALUES ('8', '/static/img/2/c.jpg', 'img', '223.00');
INSERT INTO `singleimgmsg` VALUES ('9', '/static/img/2/d.jpg', 'img', '6.00');
INSERT INTO `singleimgmsg` VALUES ('10', '/static/img/2/e.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('11', '/static/img/3/a.jpg', 'img', '62.00');
INSERT INTO `singleimgmsg` VALUES ('12', '/static/img/3/b.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('13', '/static/img/3/c.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('14', '/static/img/3/d.jpg', 'img', '32.00');
INSERT INTO `singleimgmsg` VALUES ('15', '/static/img/3/e.jpg', 'img', '23.20');
INSERT INTO `singleimgmsg` VALUES ('16', '/static/img/4/a.jpg', 'img', '264.00');
INSERT INTO `singleimgmsg` VALUES ('17', '/static/img/4/b.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('18', '/static/img/4/c.jpg', 'img', '1264.00');
INSERT INTO `singleimgmsg` VALUES ('19', '/static/img/4/f.jpg', 'img', '123.00');
INSERT INTO `singleimgmsg` VALUES ('20', '/static/img/4/e.jpg', 'img', '3.00');
INSERT INTO `singleimgmsg` VALUES ('21', '/static/img/5/a.jpg', 'img', '13.00');
INSERT INTO `singleimgmsg` VALUES ('22', '/static/img/5/b.jpg', 'img', '13.00');
INSERT INTO `singleimgmsg` VALUES ('23', '/static/img/5/c.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('24', '/static/img/5/d.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('25', '/static/img/5/e.jpg', 'img', '23.00');
INSERT INTO `singleimgmsg` VALUES ('26', '/static/img/6/a.jpg', 'img', '15.00');
INSERT INTO `singleimgmsg` VALUES ('27', '/static/img/6/b.jpg', 'img', '12.00');
INSERT INTO `singleimgmsg` VALUES ('28', '/static/img/6/c.jpg', 'img', '12.00');
INSERT INTO `singleimgmsg` VALUES ('29', '/static/img/6/d.jpg', 'img', '12.00');
INSERT INTO `singleimgmsg` VALUES ('30', '/static/img/6/e.jpg', 'img', '1263.00');
INSERT INTO `singleimgmsg` VALUES ('31', '/static/img/7/a.jpg', 'img', '32.00');
INSERT INTO `singleimgmsg` VALUES ('32', '/static/img/7/b.jpg', 'img', '15.00');
INSERT INTO `singleimgmsg` VALUES ('33', '/static/img/7/c.jpg', 'img', '12.00');
INSERT INTO `singleimgmsg` VALUES ('34', '/static/img/7/d.jpg', 'img', '45.00');
INSERT INTO `singleimgmsg` VALUES ('35', '/static/img/7/e.jpg', 'img', '28.00');
INSERT INTO `singleimgmsg` VALUES ('36', '/static/img/8/a.jpg', 'img', '49.00');
INSERT INTO `singleimgmsg` VALUES ('37', '/static/img/8/b.jpg', 'img', '7.00');
INSERT INTO `singleimgmsg` VALUES ('38', '/static/img/8/c.jpg', 'img', '98.00');
INSERT INTO `singleimgmsg` VALUES ('39', '/static/img/8/d.jpg', 'img', '58.00');
INSERT INTO `singleimgmsg` VALUES ('40', '/static/img/8/e.jpg', 'img', '87.00');
INSERT INTO `singleimgmsg` VALUES ('41', '/static/img/9/a.jpg', 'img', '45.00');
INSERT INTO `singleimgmsg` VALUES ('42', '/static/img/9/b.jpg', 'img', '32.00');
INSERT INTO `singleimgmsg` VALUES ('43', '/static/img/9/c.jpg', 'img', '64.00');
INSERT INTO `singleimgmsg` VALUES ('44', '/static/img/9/d.jpg', 'img', '123.00');
INSERT INTO `singleimgmsg` VALUES ('45', '/static/img/9/e.jpg', 'img', '12.00');
INSERT INTO `singleimgmsg` VALUES ('46', '/static/img/10/a.jpg', 'img', '54.00');
INSERT INTO `singleimgmsg` VALUES ('47', '/static/img/10/b.jpg', 'img', '15.00');
INSERT INTO `singleimgmsg` VALUES ('48', '/static/img/10/c.jpg', 'img', '64.00');
INSERT INTO `singleimgmsg` VALUES ('49', '/static/img/10/d.jpg', 'img', '565.00');
INSERT INTO `singleimgmsg` VALUES ('50', '/static/img/10/e.jpg', 'img', '45.00');
INSERT INTO `singleimgmsg` VALUES ('51', '/static/img/11/a.jpg', 'img', '65.00');
INSERT INTO `singleimgmsg` VALUES ('52', '/static/img/11/b.jpg', 'img', '16.00');
INSERT INTO `singleimgmsg` VALUES ('53', '/static/img/11/c.jpg', '美好信息', '23.00');
INSERT INTO `singleimgmsg` VALUES ('54', '/static/img/11/d.jpg', '美好信息', '565.00');
INSERT INTO `singleimgmsg` VALUES ('55', '/static/img/11/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('56', '/static/img/12/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('57', '/static/img/12/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('58', '/static/img/12/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('59', '/static/img/12/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('60', '/static/img/12/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('61', '/static/img/13/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('62', '/static/img/13/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('63', '/static/img/13/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('64', '/static/img/13/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('65', '/static/img/13/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('66', '/static/img/14/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('67', '/static/img/14/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('68', '/static/img/14/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('69', '/static/img/14/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('70', '/static/img/14/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('71', '/static/img/15/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('72', '/static/img/15/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('73', '/static/img/15/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('74', '/static/img/15/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('75', '/static/img/15/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('76', '/static/img/16/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('77', '/static/img/16/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('78', '/static/img/16/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('79', '/static/img/16/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('80', '/static/img/16/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('81', '/static/img/17/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('82', '/static/img/17/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('83', '/static/img/17/f.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('84', '/static/img/17/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('85', '/static/img/17/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('86', '/static/img/18/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('87', '/static/img/18/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('88', '/static/img/18/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('89', '/static/img/18/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('90', '/static/img/18/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('91', '/static/img/19/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('92', '/static/img/19/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('93', '/static/img/19/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('94', '/static/img/19/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('95', '/static/img/19/e.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('96', '/static/img/20/a.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('97', '/static/img/20/b.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('98', '/static/img/20/c.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('99', '/static/img/20/d.jpg', '美好信息', '45.00');
INSERT INTO `singleimgmsg` VALUES ('100', '/static/img/20/e.jpg', '美好信息', '45.00');

-- ----------------------------
-- Table structure for usermsg
-- ----------------------------
DROP TABLE IF EXISTS `usermsg`;
CREATE TABLE `usermsg` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `tagname` varchar(30) DEFAULT NULL,
  `name` varchar(30) DEFAULT NULL,
  `sex` varchar(12) DEFAULT NULL,
  `idcard` varchar(50) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of usermsg
-- ----------------------------
INSERT INTO `usermsg` VALUES ('1', '发过火', '发', 'null', '4564564', '456465');
INSERT INTO `usermsg` VALUES ('2', '电饭锅', '电饭锅', 'null', '446', '654654');
INSERT INTO `usermsg` VALUES ('3', '地方', '电饭锅', 'null', '23434435', '345345');
INSERT INTO `usermsg` VALUES ('4', '电饭锅', '电饭锅', 'null', '34345', '45345');
INSERT INTO `usermsg` VALUES ('5', '玩儿', '玩儿', 'undefined', '23434', '234');
INSERT INTO `usermsg` VALUES ('6', 'axc', 'asd', 'undefined', '456456464', '4564');
INSERT INTO `usermsg` VALUES ('7', 'asd', 'asd', '女', '46545646545456456', '456');

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `video` varchar(255) DEFAULT NULL,
  `times` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES ('1', '/static/vedio/农牧云2.mp4', '2018-11-14 11:19:37');

-- ----------------------------
-- Table structure for ysjsgood
-- ----------------------------
DROP TABLE IF EXISTS `ysjsgood`;
CREATE TABLE `ysjsgood` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `mess` varchar(30) DEFAULT NULL,
  `momess` varchar(255) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `hrefs` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ysjsgood
-- ----------------------------
INSERT INTO `ysjsgood` VALUES ('1', '/static/img/13/1.jpg', '青金桔新鲜小青桔', '适宜急慢性气管炎，肝炎，胆囊炎，高血压，直管硬化者食用。脾弱气虚之人不宜多食，糖尿病人忌食。凡口舌碎痛，齿龈肿痛者忌食。', '35.80', 'https://www.linchongyang.cn:4320/single?id=13');
INSERT INTO `ysjsgood` VALUES ('2', '/static/img/14/1.jpg', '农家自种蔬菜农产品', '红心土豆还含有丰富的维生素B1、B2、B6和泛酸等B群维生素微量元素，氨基酸、这些成分在人的肌体抗老防病过程中有着重要的作用', '45.90', 'https://www.linchongyang.cn:4320/single?id=14');
INSERT INTO `ysjsgood` VALUES ('3', '/static/img/15/1.jpg', '新鲜紫皮独头蒜', '蒜皮呈紫色，薄而富有弹性，果实脆而白，半透明，辣味正，口感好，富含多种营养成份和微量元素，具有美容、健胃、杀菌功能', '36.90', 'https://www.linchongyang.cn:4320/single?id=15');

-- ----------------------------
-- Table structure for yunshop1
-- ----------------------------
DROP TABLE IF EXISTS `yunshop1`;
CREATE TABLE `yunshop1` (
  `id` int(30) NOT NULL AUTO_INCREMENT,
  `imgs` varchar(255) DEFAULT NULL,
  `mess` varchar(255) DEFAULT NULL,
  `hrefs` varchar(255) DEFAULT NULL,
  `price` decimal(10,2) DEFAULT NULL,
  `number` int(30) DEFAULT NULL,
  `moremess` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of yunshop1
-- ----------------------------
INSERT INTO `yunshop1` VALUES ('1', '/static/img/1/1.jpg', '十年老母鸡', 'https://www.linchongyang.cn:4320/single?id=1', '30.40', '502', '十多个电饭锅电饭锅电饭锅');
INSERT INTO `yunshop1` VALUES ('2', '/static/img/2/1.jpg', '似懂非懂', 'https://www.linchongyang.cn:4320/single?id=2', '30.50', '3', '法规的非官方个');
INSERT INTO `yunshop1` VALUES ('3', '/static/img/3/1.jpg', '十年老母鸡电饭锅', 'https://www.linchongyang.cn:4320/single?id=3', '12.15', '2', '个电饭锅电饭锅');
INSERT INTO `yunshop1` VALUES ('4', '/static/img/4/1.jpg', '给对方回复过后', 'https://www.linchongyang.cn:4320/single?id=4', '465.00', '2323', '是电饭锅电饭锅');
INSERT INTO `yunshop1` VALUES ('5', '/static/img/5/1.jpg', '肺根后方', 'https://www.linchongyang.cn:4320/single?id=5', '163.40', '3', '电饭锅电饭锅电饭锅');
INSERT INTO `yunshop1` VALUES ('6', '/static/img/6/1.jpg', '跗骨结核', 'https://www.linchongyang.cn:4320/single?id=6', '23.50', '23', '第三个大哥大法官好');
INSERT INTO `yunshop1` VALUES ('7', '/static/img/7/1.jpg', '房管局和规划局', 'https://www.linchongyang.cn:4320/single?id=7', '23.40', '23', '大范甘迪发货后');
INSERT INTO `yunshop1` VALUES ('8', '/static/img/8/1.jpg', '地方规划局规划局', 'https://www.linchongyang.cn:4320/single?id=8', '213.40', '23', '电饭锅电饭锅个');
INSERT INTO `yunshop1` VALUES ('9', '/static/img/9/1.jpg', '发的规划', 'https://www.linchongyang.cn:4320/single?id=9', '23.60', '356', '电饭锅电饭锅电饭锅');
INSERT INTO `yunshop1` VALUES ('10', '/static/img/10/1.jpg', '发是各回各家建行卡通', 'https://www.linchongyang.cn:4320/single?id=10', '236.40', '5656', '电饭锅电饭锅电饭锅发');
INSERT INTO `yunshop1` VALUES ('11', '/static/img/11/1.jpg', '惩戒光环', 'https://www.linchongyang.cn:4320/single?id=11', '56.53', '50', '规划局规划局规划局规划局');
INSERT INTO `yunshop1` VALUES ('12', '/static/img/12/1.jpg', '方土异同', 'https://www.linchongyang.cn:4320/single?id=12', '23.40', '56', '市房管局水电费她的头');
