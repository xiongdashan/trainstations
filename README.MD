全国火车站及所属城市名称、代码
=========================================

火车站所属地数据名称、 拼音、 所在城市、 所在城市代码... 结构如下：

```go
type TrainStation struct {
	PinYinShortName string  
	Name            string
	Code            string
	PinYinFull      string
	City            string
	CityCode        string
	Province        string
}

```

## 实现步骤

1. 从12306获取火车站信息的Js文件

2. 去高德注册个人账号，按火车站名查询所在地

3. 组合信息拼装Js文件

## 拿来主义

整理好的Js文件 [./data/train_stations.js](./data/train_stations.js)

格式如下：

```js
var station_names =[{"PinYinShortName":"bjb","Name":"北京北","Code":"VAP","PinYinFull":"beijingbei","City":"","CityCode":"","Province":""}];
```


## 自已运行

替换 ```main.go```中Key

```go
	
	const geoKey = "你的高德Key"


```