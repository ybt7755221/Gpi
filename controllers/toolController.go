package controllers

import (
	et "gpi/entities"
	"gpi/libraries/jwtGo"
	"gpi/libraries/kafka"
	"gpi/libraries/verify"
	"gpi/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type ToolController struct {
	serv *service.ToolService
}

func (t *ToolController) Token(ctx *gin.Context) {
	var ts string
	methodStr := ctx.Request.Method
	if methodStr == "GET" || methodStr == "DELETE"{
		ts = ctx.Query("ts")
	} else {
		ts = ctx.PostForm("ts")
	}
	if len(ts) == 0 {
		resError(ctx, et.EntityParametersMissing, et.GetStatusMsg(et.EntityParametersMissing))
		return
	}
	rawStr, tokenStr := verify.GenerateToken(ctx)
	fmt.Println(rawStr, tokenStr)
	resSuccess(ctx, gin.H{
		"raw" : rawStr,
		"token" : tokenStr,
	})
	return
}

// @Tags 工具库
// @Summary 【api】根据请求信息返回token
// @Description 根据请求信息返回token
// @Accept json
// @Produce json
// @Success 200 {object} SgrResp
// @Router /tool/jwt-token [post]
func (t *ToolController) GetJWTToken(ctx *gin.Context) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var in map[string]string
	json.Unmarshal(data, &in)
	resp, err := t.serv.GetJWTToken(in)
	if err != nil {
		resError(ctx, et.EntitySystemError, err.Error())
		return
	}
	resSuccess(ctx, resp)
	return
}

// @Tags 工具库
// @Summary 【api】解析jwt token
// @Description 解析jwt token
// @Accept json
// @Produce json
// @Success 200 {object} SgrResp
// @Router /tool/jwt-token [post]
func (t *ToolController) ParseJWTToken(ctx *gin.Context) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var in map[string]string
	json.Unmarshal(data, &in)
	resp, err := jwtGo.ParseToken(in["token"])
	if err != nil {
		resError(ctx, et.EntitySystemError, err.Error())
		return
	}
	resSuccess(ctx, resp)
	return
}
// @Tags 工具库
// @Summary 【api】推送kafka消息-生产者接口
// @Description 推送kafka消息
// @Accept json
// @Param   topic	path	string 	true	"kafka topic-消费topic"
// @Param   key		path	string 	true	"kafka key-消费用的key"
// @Param   value	path	string 	true	"kafka value-消费用value"
// @Produce json
// @Success 200 {object} SgrResp
// @Router /kafka/ [post]
func (t *ToolController) SendKafka(ctx *gin.Context) {
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	var req map[string]string
	if err := json.Unmarshal(data, &req); err != nil {
		resError(ctx, et.EntitySystemError, err.Error())
		return
	}
	if req["topic"] == "" {
		resError(ctx, et.EntityParametersMissing, "topic不能为空")
	}
	byteValue, err := json.Marshal(req["value"])
	if err != nil {
		resError(ctx, et.EntitySystemError, err.Error())
		return
	}
	producer := kafka.Kafka{
		req["topic"],
		req["key"],
		"",
		byteValue,
	}
	if err := kafka.Producer(producer); err != nil {
		resError(ctx, et.EntitySystemError, err.Error())
	}else{
		resSuccess(ctx,producer)
	}
}