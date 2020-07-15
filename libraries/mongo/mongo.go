package mongo

import (
	"context"
	"fmt"
	"gpi/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MgoClient *mongo.Client

func init() {
	connect()
}

func connect() {
	mgoConf := config.GMConfig["system_log"]
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", mgoConf.User, mgoConf.Pass, mgoConf.Host, mgoConf.Port)
	var err error
	MgoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("connnect mongo error :" + err.Error())
	}
}

func InsertOne(db string, table string, data bson.M) {
	MgoClient.Database(db).Collection(table).InsertOne(context.Background(), data)
}
