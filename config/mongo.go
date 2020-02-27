package config

const(
	SYS_LOG = "sys_log"
)

type MgoStruct struct {
	name 					string
	host 					string
	port 					string
	user 					string
	pass 					string
	PoolLimit 				string
	Timeout 				string
	Direct 					string
	ReplicaSetName 			string
	Group 					string
}
var GMConfig map[string]MgoStruct

func init() {
	GMConfig = make(map[string]MgoStruct,0)
	GMConfig[SYS_LOG] = MgoStruct{
		name:        GetApolloString("SYS_LOG_NAME", ""),
		host:        GetApolloString("SYS_LOG_HOST", ""),
		port:        GetApolloString("SYS_LOG_PORT", ""),
		user:        GetApolloString("SYS_LOG_USER", ""),
		pass:        GetApolloString("SYS_LOG_PASS", ""),
		PoolLimit :  GetApolloString("MONGO_POOL_LIMIT", ""),
		Timeout:     GetApolloString("MONGO_TIMEOUT", ""),
		Direct:      GetApolloString("MONGO_DIRECT", ""),
		ReplicaSetName: GetApolloString("SYS_LOG_REPLICASET", ""),
		Group         :  GetApolloString("SYS_LOG_HOST_PORT_GROUP", ""),
	}
}