package mongostore

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const (
	_conn = "mongodb://localhost:27017"
	_dbName = "poc"
	_default = "test1"
)

type store struct {
	conn *mongo.Client
	database *mongo.Database
	*mongo.Collection
}

func NewStore() *store {
	connOption := options.Client().ApplyURI(_conn)
	duration := time.Second * 10
	connOption.SetConnectTimeout(duration)
	
	ctx,cancel := context.WithTimeout(context.Background(), time.Millisecond*2)
	defer cancel()

	client,err := mongo.Connect(ctx, connOption)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("connected to MongoDB !!")
	db := client.Database(_dbName)
	c := db.Collection(_default)
	o := store {
		conn: client,
		database: db,
		Collection: c,
	}
	return &o
}

func (s *store) Healthy() bool{
//	ctx,cancel := context.WithTimeout(context.Background(), time.Second*1)
	//defer cancel()
	err := s.conn.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return false 
	}
	return true
}

func (s *store) Close() {
	ctx,cancel := context.WithTimeout(context.Background(), time.Second * 1)
	defer cancel()
	err := s.conn.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connection closed")
}
