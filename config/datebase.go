package config

type MysqlConf struct {
	Host string
	Port string
	Name string
	User string
	Passwd string
	Charset string
	OpenMax int
	IdleMax int
}

const(
	Gin = "gin"
)

var MysqlConfMap map[string]MysqlConf

func init() {
	//读库操作
	msqConfMap := map[string]MysqlConf{
		Gin : {
			Host : GetApolloString("GIN_HOST", "127.0.0.1"),
			Port : GetApolloString("GIN_PORT", "3306"),
			Name : GetApolloString("GIN_NAME", "gin"),
			User : GetApolloString("GIN_USER", "GinUser"),
			Passwd : GetApolloString("GIN_PASSWD", "userGin"),
			Charset: "utf8",
			OpenMax: GetApolloInt("MYSQL_MAX_OPEN_CONN", 150),
			IdleMax: GetApolloInt("MYSQL_MAX_IDEL_CONN", 80),
		},
 	}
	MysqlConfMap = msqConfMap
}
