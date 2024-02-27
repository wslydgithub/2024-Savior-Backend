package signup

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// 连接redis
func Buildclient() *redis.Client {
	Client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return Client
}

func Savecode(client *redis.Client, Email string, Code string) {
	key := "user:" + Email // 设置一个唯一的键
	// 存入哈希表
	err := client.HSet(key, "email", Email).Err()
	if err != nil {
		fmt.Println("err(email save):", err)
		return
	}
	// 存入验证码字段
	err = client.HSet(key, "code", Code).Err()
	if err != nil {
		fmt.Println("err(code save):", err)
		return
	}
	// 设置过期时间
	err = client.Expire(key, 5*time.Minute).Err()
	if err != nil {
		fmt.Println("err(expire):", err)
		return
	}
	// 关联邮箱和验证码
	err = client.HSet("codes", Email, Code).Err()
	if err != nil {
		fmt.Println("err(link wrong):", err)
		return
	}
}

func Getcode(client *redis.Client, Email string) string {
	key := "user:" + Email
	// 从哈希表中获取数据
	email, err := client.HGet(key, "email").Result()
	if err != nil {
		fmt.Println("Get email wrong:", err)
		return ""
	}
	fmt.Println("Get email :", email)

	// 从哈希表中获取关联的验证码
	linkcode, err := client.HGet("codes", Email).Result()
	if err != nil {
		fmt.Println("Get linkcode wrong:", err)
		return ""
	}
	defer client.Close()
	return linkcode
}
