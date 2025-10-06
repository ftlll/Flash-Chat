package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
	ctx   = context.Background()
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.Get("mysql.DB_USER"),
		viper.Get("mysql.DB_PASSWORD"),
		viper.Get("mysql.DB_HOST"),
		viper.Get("mysql.DB_PORT"),
		viper.Get("mysql.DB_NAME"))

	fmt.Println(dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	} else {
		fmt.Println("database connected")
	}
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conn"),
	})
	pong, err := Redis.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("failed to connect redis: %v", err))
	} else {
		fmt.Println("Redis connected", pong)
	}
}

const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	err = Redis.Publish(ctx, channel, msg).Err()
	fmt.Println("publish...", msg)
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Redis.Subscribe(ctx, channel)
	fmt.Println("subscribe...1", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("subscribe...", msg)
	return msg.Payload, err
}
