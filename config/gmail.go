package config

type EmailConf struct {
	Host 	string
	Port 	int
	User 	string
	Passwd	string
	To		string
	ErrTopic string
}

var EmailConfStruct EmailConf

func init () {
	EmailConfStruct = EmailConf{
		Host 	: GetApolloString("MAIL_HOST","smtp.163.com"),
		Port 	: GetApolloInt("MAIL_PORT",25),
		User 	: GetApolloString("MAIL_USER","burt_yu@example.com"),
		Passwd	: GetApolloString("MAIL_PASS","xxxxxxxxxx"),
		To		: GetApolloString("MAIL_SEND",""),
		ErrTopic : "【Gpi系统】错误",
	}
}
