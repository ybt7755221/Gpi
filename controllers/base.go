package controllers

import (
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
	sortField := c.Query("sortField")
	sort, err := strconv.Atoi(c.Query("sort"))
	if err != nil || sort != 1 {
		sort = 2
	}
	commonParams := gin.H{
		"offset"  : offset,
		"limit"   : limit,
		"sortField" : sortField,
		"sort"    : sort,
	}
	return commonParams
}

func resJson(c *gin.Context, httpCode int, data ApiResonse) {
	c.JSON(httpCode, data)
}

func resResult(c *gin.Context, code int, msg string, data interface{}) {
	dataStruct := ApiResonse{code, msg, data}
	resJson(c, http.StatusOK, dataStruct)
}

func resSuccess(c *gin.Context, data interface{}) {
	dateStruct := ApiResonse{ReqIsOk, GetStatusMsg(ReqIsOk), data}
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