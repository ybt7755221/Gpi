package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	. "gpi/entities"
)
var pageSize = 20

func getParams(c *gin.Context, fields []string) map[string]string {
	condition := make(map[string]string)
	for _, val := range fields {
		condition[val] = c.Query(val)
	}
	return condition
}

func getCommonParams(c *gin.Context) gin.H{
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit == 0 {
		limit = pageSize
	}
	sortStr := c.Query("sort")
	sort := make(map[string]string, 0)
	if len(sortStr) > 0 {
		sort = getSort(sortStr)
	}
	commonParams := gin.H{
		"offset"  : offset,
		"limit"   : limit,
		"sort"    : sort,
	}
	return commonParams
}

func getSort(sortStr string) (sort map[string]string) {
	if err := json.Unmarshal([]byte(sortStr), &sort); err != nil {
		return nil
	}
	return
}

func resJson(c *gin.Context, httpCode int, data ApiResonse) {
	c.JSON(httpCode, data)
}

func resResult(c *gin.Context, code int, msg string, data interface{}) {
	dataStruct := ApiResonse{code, msg, data}
	resJson(c, http.StatusOK, dataStruct)
}

func resSuccess(c *gin.Context, data interface{}) {
	dateStruct := ApiResonse{EntityIsOk, GetStatusMsg(EntityIsOk), data}
	resJson(c, http.StatusOK, dateStruct)
}

func resError(c *gin.Context, code int, msg string) {
	dateStruct := ApiResonse{code, msg, gin.H{}}
	resJson(c, http.StatusOK, dateStruct)
}
/**
 * 字符格式转时间
 * "2006-01-02 15:04:05"
 */
func string2Time(tSter string) (theTime time.Time){
	loc, _ := time.LoadLocation("Local")
	theTime, _ = time.ParseInLocation("2006-01-02", tSter, loc)
	return
}