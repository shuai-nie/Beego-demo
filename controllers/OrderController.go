package controllers

import (
	"Beego-demo/models"
	"Beego-demo/serices/mq"
	redisClient "Beego-demo/serices/redis"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"strconv"
	"time"
)

type OrderController struct {
	BaseController
}

func (o *OrderController) Seckill() {
	skuId, _ := o.GetInt("sku_id")
	address := o.GetString("address")

	if skuId == 0 {
		o.Data["json"] = &models.JsonResult{Code: 4001, Msg: ""}
		o.ServeJSONP()
		return
	}
	if address == "" {
		o.Data["json"] = &models.JsonResult{Code: 4002, Msg: ""}
		o.ServeJSONP()
		return
	}

	redisConn := redisClient.PoolConnect()
	defer redisConn.Close()

	// 1.在缓存中校验库存
	stockKey := "go_stock:" + strconv.Itoa(skuId)
	stock, err := redis.Int(redisConn.Do("get", stockKey))
	if err != nil {
		o.Data["json"] = &models.JsonResult{Code: 4003, Msg: ""}
		o.ServeJSONP()
		return
	}

	if stock < 1 {
		o.Data["json"] = &models.JsonResult{Code: 4005, Msg: "该商品库存不足"}
		o.ServeJSON()
		return
	}

	// 2:在缓存中检验秒杀是否开始
	xt := ExpireTime{}
	expire, err := redis.String(redisConn.Do("get", "go_expire:1"))
	if err != nil {
		o.Data["json"] = &models.JsonResult{Code: 4006, Msg: ""}
		o.ServeJSON()
		return
	}else{
		json.Unmarshal([]byte(expire), &xt)
		start := xt.Start
		end := xt.End

		local, _ := time.LoadLocation("Local")
		startTime, _ := time.ParseInLocation("", start, local)
		endTime, _ := time.ParseInLocation("", end, local)
		new := time.Now()

		if startTime.After(new) {
			o.Data["json"] = &models.JsonResult{Code: 4007, Msg: "秒杀还未开始"}
			o.ServeJSON()
			return
		}

		if endTime.Before(new) {
			o.Data["json"] = &models.JsonResult{Code: 4007, Msg: ""}
			o.ServeJSON()
			return
		}
	}

	// 3: 在缓存中校验该用户是否秒杀过
	userId := rand.Intn(99) +1
	userOrderKey := "go_user_order_" + strconv.Itoa(skuId) + ":" +strconv.Itoa(userId)
	order, err := redis.String(redisConn.Do("get", userOrderKey))
	if err != nil {
		o.Data["json"] = &models.JsonResult{Code: 4008, Msg: "该用户已经下过单"+order}
		o.ServeJSON()
		return
	}

	// 4：秒杀入队
	msg := &MsgData{}
	msg.TaskName = "seckill_order"
	msg.SkuId = userId
	msg.Address = address
	msg.Time = time.Now().Format("2006-01-02 15:04:05")
	msgStr, _ := json.Marshal(msg)
	mq.Publish("", "go_seckill", string(msgStr))

	o.Data["json"] = &models.JsonResult{Code: 2000, Msg: "秒杀中"}
	o.ServeJSON()
	return
	
}

type MsgData struct {
	TaskName string `json:"task_name"`
	UserId   int    `json:"user_id"`
	SkuId    int    `json:"sku_id"`
	Address  string `json:"address"`
	Time     string `json:"time"`
}

type ExpireTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func (o *OrderController) Set() {
	redisConn := redisClient.PoolConnect()
	defer redisConn.Close()
	redisConn.Do("set", "go_stock:1", 10)

	xt := &ExpireTime{}
	xt.Start = "2021-01-15 00:00:00"
	xt.End = "2021-02-15 00:00:00"
	data, _ := json.Marshal(xt)
	redisConn.Do("set", "go_expire:1", string(data))

	o.Ctx.WriteString("设置缓存成功")
}

func (o *OrderController) Get() {
	redisConn := redisClient.PoolConnect()
	defer redisConn.Close()

	xt := ExpireTime{}
	expire, err := redis.String(redisConn.Do("get", "go_expire:1"))
	if err != nil {
		o.Ctx.WriteString("商品秒杀时间缓存不存在")
	}else{
		err := json.Unmarshal([]byte(expire), &xt)
		fmt.Println(err)
		fmt.Println(xt.Start)
		fmt.Println(xt.End)
	}

	stock, err := redis.String(redisConn.Do("get", "go_stock:1"))
	if err != nil {
		o.Ctx.WriteString("商品sku 库存缓存key 不存在")
	} else {
		o.Ctx.WriteString(stock)
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	skuInfo, err := models.GetSkuInfo(1)
	if err == nil {
		fmt.Println(skuInfo.Stock)
	} else {
		fmt.Println("没有sku")
		fmt.Println(err)
	}

	//res, err := models.UpdateStock(1)
	//if err == nil {
	//	fmt.Println("mysql row affected nums: ", res)
	//
	//} else {
	//	fmt.Println("更新库存失败")
	//	fmt.Println(err)
	//}
	//
	//id,err := models.SaveOrder("golang3567891qwqw2",2, "广州")
	//if err == nil{
	//	err := models.SaveItem(id,1)
	//	if err != nil{
	//		fmt.Println(err)
	//	}
	//}else{
	//	fmt.Println(err)
	//}
}