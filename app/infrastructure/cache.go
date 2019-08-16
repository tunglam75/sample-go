package infrastructure

import (
	"os"

	"github.com/garyburd/redigo/redis"
)

// Cache struct.
type Cache struct {
	// Cache is global redis connect
	Conn *redis.Conn
}

// NewCache returns new cacheHandler.
// repository: https://github.com/garyburd/redigo/redis
func NewCache() *Cache {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASS")

	var err error
	options := redis.DialPassword(pass)
	c, err := redis.Dial("tcp", host+":"+port, options)

	if err != nil {
		panic(err)
	}
	return &Cache{Conn: &c}
}

// CloseRedis close redis connection
func CloseRedis(conn *redis.Conn) {
	// close redis connection.
	if conn != nil {
		err := (*conn).Close()
		if err != nil {
			panic(err)
		}
	}
}
