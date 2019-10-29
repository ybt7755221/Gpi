package entities

type ApiResonse struct {
	Code 	int			`json:"code"`
	Msg		string		`json:"msg"`
	Data 	interface{}	`json:"data"`
}

const ReqIsOk = 1000
const ReqParametersMissing = 1001
const ReqUnauthorized = 1401
const ReqForbidden = 1403
const ReqTimeout = 1408

var statusMsg = map[string]map[int]string{
	"cn" : {
		ReqIsOk: "请求成功",
		ReqParametersMissing: "缺少请求参数",
		ReqUnauthorized: "签名验证失败",
		ReqForbidden: "请求被禁止",
		ReqTimeout: "请求超时",
	},
	"en" : {
		ReqIsOk: "Request Success",
		ReqParametersMissing: "The Some Parameters is Missing",
		ReqUnauthorized: "Request Unauthorized",
		ReqForbidden: "Request Forbidden",
		ReqTimeout: "Request Timeout",
	},
}

func GetStatusMsg(code int, lang string) string {
	return statusMsg[lang][code]
}
