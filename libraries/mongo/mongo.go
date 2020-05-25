package mongo

import (
	"fmt"
	conf "gpi/config"
	"gpi/libraries/elog"
	"log"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	once        sync.Once
	GlobalMongo struct {
		//m_lock  sync.RWMutex
		Session map[string]*mgo.Session
	}
)

// 实现 mongo.Logger 的接口
//type MongoLog struct {
//}

//func (MongoLog)Output(calldepth int, s string) error {
//	log.SetFlags(log.Lshortfile)
//	return log.Output(calldepth,s)
//}

func init() {
	GlobalMongo.Session = make(map[string]*mgo.Session)
}

func getApolloConfig(db string) map[string]string {
	dbConfig, ok := conf.GMConfig[db]
	if !ok {
		elog.New(fmt.Sprintf("mongo配置文件没有配置{%s}", db), elog.FileMsg{})
	}
	config := make(map[string]string)
	getType := reflect.TypeOf(dbConfig)
	getValue := reflect.ValueOf(dbConfig)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		config[field.Name] = value.String()
	}
	return config
}

func in_array(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
func GetMongoSession(db string) {
	_, ok := GlobalMongo.Session[db]
	if !ok {
		once.Do(func() {
			//GlobalMongo.m_lock.Lock()
			dbConfig := getApolloConfig(db)
			host_port := fmt.Sprintf("%s:%s", dbConfig["host"], dbConfig["port"])
			arr_addrs := []string{host_port}
			group := dbConfig["Group"]
			if len(group) > 0 {
				arr_group := strings.Split(group, ",")
				for _, v := range arr_group {
					arr_addrs = append(arr_addrs, v)
				}
			}
			timeout, _ := strconv.Atoi(dbConfig["Timeout"])
			poollimit, _ := strconv.Atoi(dbConfig["PoolLimit"])
			direct_type := false
			if dbConfig["Direct"] == "true" {
				direct_type = true
			}
			dialInfo := &mgo.DialInfo{
				Addrs:          arr_addrs,
				Timeout:        time.Duration(timeout) * time.Second,
				Database:       db,
				Direct:         direct_type,
				Username:       dbConfig["user"],
				Password:       dbConfig["pass"],
				PoolLimit:      poollimit,
				ReplicaSetName: dbConfig["ReplicaSetName"],
			}
			//fmt.Println(dialInfo)
			//mgo.SetDebug(true)
			//mgo.SetLogger(new(MongoLog)) // 设置日志.
			s, err := mgo.DialWithInfo(dialInfo)
			if err != nil {
				elog.New("Create Session: "+err.Error()+"\n", elog.FileMsg{})
			}
			GlobalMongo.Session[db] = s
		})
		//GlobalMongo.m_lock.Unlock()
	}

}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	GetMongoSession(db)
	ms := GlobalMongo.Session[db].Copy()
	c := ms.DB(db).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}

func getDb(db string) (*mgo.Session, *mgo.Database) {
	GetMongoSession(db)
	ms := GlobalMongo.Session[db].Copy()
	return ms, ms.DB(db)
}
func EnsureIndexKey(db, collection string, args ...string) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.EnsureIndexKey(args...)
}

func EnsureIndex(db, collection string, index mgo.Index) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.EnsureIndex(index)
}

func IsEmpty(db, collection string) bool {
	ms, c := connect(db, collection)
	defer ms.Close()
	count, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0
}

func Count(db, collection string, query interface{}) (int, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Count()
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func FindIter(db, collection string, query interface{}) *mgo.Iter {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Find(query).Iter()
}

func Update(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Update(selector, update)
}

func Upsert(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.Upsert(selector, update)
	return err
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.UpdateAll(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	return c.Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()

	_, err := c.RemoveAll(selector)
	return err
}

//insert one or multi documents
func BulkInsert(db, collection string, docs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Insert(docs...)
	return bulk.Run()
}

func BulkRemove(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()

	bulk := c.Bulk()
	bulk.Remove(selector...)
	return bulk.Run()
}

func BulkRemoveAll(db, collection string, selector ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.RemoveAll(selector...)
	return bulk.Run()
}

func BulkUpdate(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Update(pairs...)
	return bulk.Run()
}

func BulkUpdateAll(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.UpdateAll(pairs...)
	return bulk.Run()
}

func BulkUpsert(db, collection string, pairs ...interface{}) (*mgo.BulkResult, error) {
	ms, c := connect(db, collection)
	defer ms.Close()
	bulk := c.Bulk()
	bulk.Upsert(pairs...)
	return bulk.Run()
}

func PipeAll(db, collection string, pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.All(result)
}

func PipeOne(db, collection string, pipeline, result interface{}, allowDiskUse bool) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}
	return pipe.One(result)
}

func PipeIter(db, collection string, pipeline interface{}, allowDiskUse bool) *mgo.Iter {
	ms, c := connect(db, collection)
	defer ms.Close()
	var pipe *mgo.Pipe
	if allowDiskUse {
		pipe = c.Pipe(pipeline).AllowDiskUse()
	} else {
		pipe = c.Pipe(pipeline)
	}

	return pipe.Iter()

}

func Explain(db, collection string, pipeline, result interface{}) error {
	ms, c := connect(db, collection)
	defer ms.Close()
	pipe := c.Pipe(pipeline)
	return pipe.Explain(result)
}
func GridFSCreate(db, prefix, name string) (*mgo.GridFile, error) {
	ms, d := getDb(db)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Create(name)
}

func GridFSFindOne(db, prefix string, query, result interface{}) error {
	ms, d := getDb(db)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).One(result)
}

func GridFSFindAll(db, prefix string, query, result interface{}) error {
	ms, d := getDb(db)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Find(query).All(result)
}

func GridFSOpen(db, prefix, name string) (*mgo.GridFile, error) {
	ms, d := getDb(db)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Open(name)
}

func GridFSRemove(db, prefix, name string) error {
	ms, d := getDb(db)
	defer ms.Close()
	gridFs := d.GridFS(prefix)
	return gridFs.Remove(name)
}
