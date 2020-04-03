package config

type MysqlConf struct {
	Host    string
	Port    string
	Name    string
	User    string
	Passwd  string
	Charset string
	OpenMax int
	IdleMax int
}

const (
	Gin = "gin"
)

var MysqlConfMap map[string]MysqlConf

func init() {
	//库操作
	msqConfMap := map[string]MysqlConf{
		Gin: {
			Host:    GetApolloString("DB_HOST", "127.0.0.1"),
			Port:    GetApolloString("DB_PORT", "3306"),
			Name:    GetApolloString("DB_NAME", "gin"),
			User:    GetApolloString("DB_USER", "GinUser"),
			Passwd:  GetApolloString("DB_PASS", "userGin"),
			Charset: "utf8",
			OpenMax: GetApolloInt("MYSQL_MAX_OPEN_CONN", 100),
			IdleMax: GetApolloInt("MYSQL_MAX_IDEL_CONN", 60),
		},
	}
	MysqlConfMap = msqConfMap
}
