package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"ginserver/pkg"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/marstr/guid"
)

type userRepository struct {
	pool *redis.Pool
}

func NewRedigoRepository(url string) pkg.UserRepository {
	fmt.Println("building client on " + url)
	return &userRepository{
		pool: newPool(url),
	}
}

const pattern = "u"

func (ur *userRepository) Create(ctx context.Context, in pkg.UserCreateInput) (pkg.User, error) {
	conn := ur.pool.Get()
	defer conn.Close()
	ret := pkg.User{ID: guid.NewGUID().String(), Name: in.Name}
	data, err := json.Marshal(ret)
	_, err = conn.Do("HSET", pattern, ret.ID, string(data))
	if err != nil {
		return ret, fmt.Errorf("error setting key %s to %s: %v", ret.ID, ret.Name, err)
	}
	return ret, nil
}

func (ur *userRepository) GetAll(context.Context) ([]pkg.User, error) {
	var userSlice []pkg.User
	conn := ur.pool.Get()

	result, err := redis.Values(conn.Do("HGETALL", pattern))
	if err != nil {
		fmt.Println("Error al ejecutar HGETALL:", err)
		return userSlice, err
	}

	// Imprimir los campos y valores del hash
	fmt.Println("Contenido del hash:")
	for i := 0; i < len(result); i += 2 {
		user := pkg.User{}
		value := result[i+1].([]byte)
		err := json.Unmarshal(value, &user)
		if err != nil {
			return nil, err
		}
		userSlice = append(userSlice, user)
	}

	return userSlice, nil
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (ur *userRepository) get(key string) ([]byte, error) {

	conn := ur.pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	return data, err
}
