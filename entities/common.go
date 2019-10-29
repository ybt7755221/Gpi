package entities

type ApiResonse struct {
	Code 	int			`json:"code"`
	Msg		string		`json:"msg"`
	Data 	interface{}	`json:"data"`
}

const ReqIsOk = 1000
const ReqParametersMissing = 1001
const TokenMissing = 1002
const ReqUnauthorized = 1401
const ReqForbidden = 1403
const ReqTimeout = 1408
const ReqFailure = 1100

var lang string

func init() {
	lang = "cn"
}

func GetStatusMsg(code int) string {
	return statusMsg[lang][code]
}

var statusMsg = map[string]map[int]string{
	"cn" : {
		ReqIsOk: "请求成功",
		ReqParametersMissing: "缺少请求参数",
		ReqUnauthorized: "签名验证失败",
		ReqForbidden: "请求被禁止",
		ReqTimeout: "请求超时",
		TokenMissing: "缺少token值",
		ReqFailure: "请求失败",
	},
	"en" : {
		ReqIsOk: "Request Success",
		ReqParametersMissing: "The Some Parameters is Missing",
		ReqUnauthorized: "Request Unauthorized",
		ReqForbidden: "Request Forbidden",
		ReqTimeout: "Request Timeout",
		TokenMissing: "The token is missing",
		ReqFailure: "Request Failure",
	},
}
