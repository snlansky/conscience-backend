package dao

import (
	"conscience-backend/config"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var DB *sqlx.DB
var Mongo *mongo.Client
var mongoDB string

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func Init(dao config.Dao) (err error) {
	//DB = initSqlxDB(dsn, 32, 2)
	Mongo, err = initMongo(dao.Mongo.Uri)
	mongoDB = dao.Mongo.DB
	return
}

func initSqlxDB(dbConfig string, maxOpen, maxIdle int) *sqlx.DB {
	db := sqlx.MustConnect("mysql", dbConfig)
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	// https://github.com/go-sql-driver/mysql/issues/446
	db.SetConnMaxLifetime(time.Second * 14400)
	return db
}

func initMongo(uri string) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

