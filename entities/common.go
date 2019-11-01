package entities

import "sync"

type ApiResonse struct {
	Code 	int			`json:"code"`
	Msg		string		`json:"msg"`
	Data 	interface{}	`json:"data"`
}

const EntityIsOk = 1000
const EntityParametersMissing = 1001
const EntityTokenMissing = 1002
const EntitySystemError = 1003
const EntityUnauthorized = 1401
const EntityForbidden = 1403
const EntityTimeout = 1504
const EntityFailure = 1100

var lang string
var once sync.Once

func init() {
	once.Do(func() {
		lang = "cn"
	})
}

func GetStatusMsg(code int) string {
	return statusMsg[lang][code]
}

var statusMsg = map[string]map[int]string{
	"cn" : {
		EntityIsOk: "请求成功",
		EntityParametersMissing: "缺少请求参数",
		EntityUnauthorized: "签名验证失败",
		EntityForbidden: "请求被禁止",
		EntityTimeout: "请求超时",
		EntityTokenMissing: "缺少token值",
		EntityFailure: "请求失败",
		EntitySystemError: "系统错误",
	},
	"en" : {
		EntityIsOk: "Request Success",
		EntityParametersMissing: "The Some Parameters is Missing",
		EntityUnauthorized: "Request Unauthorized",
		EntityForbidden: "Request Forbidden",
		EntityTimeout: "Request Timeout",
		EntityTokenMissing: "The token is missing",
		EntityFailure: "Request Failure",
		EntitySystemError: "System Error",
	},
}
