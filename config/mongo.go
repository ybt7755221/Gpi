package config

const(
	Log = "sys_log"
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
	GMConfig[Log] = MgoStruct{
		name:        GetApolloString("MONGO_LOG_NAME", ""),
		host:        GetApolloString("MONGO_LOG_HOST", ""),
		port:        GetApolloString("MONGO_LOG_PORT", ""),
		user:        GetApolloString("MONGO_LOG_USER", ""),
		pass:        GetApolloString("MONGO_LOG_PASS", ""),
		PoolLimit :  GetApolloString("MONGO_POOL_LIMIT", ""),
		Timeout:     GetApolloString("MONGO_TIMEOUT", ""),
		Direct:      GetApolloString("MONGO_DIRECT", ""),
		ReplicaSetName: GetApolloString("MONGO_LOG_REPLICASET", ""),
		Group         :  GetApolloString("MONGO_LOG_HOST_PORT_GROUP", ""),
	}
}