package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type GinHandler struct {
	Key      func(*gin.Context) (string, error)
	Do       func(*gin.Context) (interface{}, error) // key, data, error
	Duration time.Duration

	AlwaysCache bool // avoid cache penetration
}

func DefaultRedisClient() (client *redis.Client, err error) {
	client = redis.NewClient(
		&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0},
	)

	statusCmd := client.Ping()
	if err := statusCmd.Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func (hdl *GinHandler) WithRedis(client *redis.Client) func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			key  string
			err  error
			bts  []byte
			cmd  *redis.StringCmd
			data interface{}
		)

		if key, err = hdl.Key(c); err != nil {
			c.JSON(http.StatusOK, err)
			return
		}

		respBytes := func(bts []byte) {
			c.Header("StatusCode", strconv.Itoa(http.StatusOK))
			c.Header("Status", http.StatusText(http.StatusOK))
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.Writer.Write(bts)
			return
		}

		cmd = client.Get(key)
		if err = cmd.Err(); err == nil { // get result from redis cache
			bts, _ = cmd.Bytes()
			respBytes(bts)
			return
		}

		if data, err = hdl.Do(c); err != nil { // process failed
			bts, _ = json.Marshal(err)
			if hdl.AlwaysCache && hdl.Duration > 0 {
				client.Set(key, bts, hdl.Duration) // avoid cache penetration
			}
			respBytes(bts)
			return
		}

		bts, _ = json.Marshal(data)
		if hdl.Duration > 0 {
			client.Set(key, bts, hdl.Duration)
		}

		respBytes(bts)
		return
	}
}

func GinWithRedis(
	do func(*gin.Context, bool) (string, interface{}, error),
	client *redis.Client, duration time.Duration,
	alwaysCache ...bool) func(*gin.Context) {

	if duration <= 0 { // don't allow no cache or never expire
		return nil
	}

	return func(c *gin.Context) {
		var (
			key  string
			err  error
			bts  []byte
			cmd  *redis.StringCmd
			data interface{}
		)

		if key, _, err = do(c, false); err != nil {
			c.JSON(http.StatusOK, err)
			return
		}

		defer func() {
			c.Header("StatusCode", strconv.Itoa(http.StatusOK))
			c.Header("Status", http.StatusText(http.StatusOK))
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.Writer.Write(bts)
			return
		}()

		cmd = client.Get(key)
		if err = cmd.Err(); err == nil { // get result from redis cache
			bts, _ = cmd.Bytes()
			return
		}

		if _, data, err = do(c, true); err != nil { // proccess failed
			bts, _ = json.Marshal(err)
			if len(alwaysCache) > 0 && alwaysCache[0] {
				client.Set(key, bts, duration) // avoid cache penetration
			}
			return
		}

		bts, _ = json.Marshal(data)
		client.Set(key, bts, duration)
		return
	}
}
