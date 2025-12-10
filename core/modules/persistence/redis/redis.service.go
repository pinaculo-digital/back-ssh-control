package redis_service

import (
	"fmt"
	text "go_service/core/util/debug"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const EMAIL_CODE_KEY = "email_code:"
const IP_ACCESS_KEY = "ip_try_access:"

type RedisService struct {
	client *redis.Client
}

func NewRedisService() (*RedisService, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	text.Println("Conectando ao redis em:", client.Options().Addr)

	_, err := client.Ping().Result()
	if err != nil {
		text.Errorln("Failed to connect to Redis: ", err)
	}
	return &RedisService{
		client: client,
	}, err
}

func (s *RedisService) SetEmailCode(email string, code string) {
	key := EMAIL_CODE_KEY + email
	s.client.Set(key, code, time.Hour)
}

func (s *RedisService) GetEmailCode(email string) (code string, err error) {
	key := EMAIL_CODE_KEY + email
	get := s.client.Get(key)
	return get.Val(), get.Err()
}

func (s *RedisService) DeleteEmailCode(email string) error {
	key := EMAIL_CODE_KEY + email
	return s.client.Del(key).Err()
}

func (s *RedisService) AddIpTryAccess(ip string) {

	pipe := s.client.TxPipeline()

	key := IP_ACCESS_KEY + ip

	pipe.Incr(key)
	pipe.Expire(key, time.Hour*24)

	pipe.Exec()

}

func (s *RedisService) GetIpTryAccess(ip string) (value int, err error) {
	key := IP_ACCESS_KEY + ip
	get := s.client.Get(key)

	value, err = strconv.Atoi(get.Val())
	if err != nil {
		return 0, fmt.Errorf("valor não encontrado")
	}
	return value, err
}
