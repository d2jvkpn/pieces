package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Err     error       `json:"-"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResData(code int, message string, errs ...error) (rd *ResData) {
	var err error

	if len(errs) > 0 {
		err = errs[0]
	}

	if message == "" && err != nil {
		message = err.Error()
	}

	return &ResData{Code: code, Message: message, Err: err}
}

func (rd *ResData) Error() string { // implememt error interface
	if rd.Err != nil {
		return rd.Err.Error()
	}
	return "<nil>"
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

func GinWithRedis(
	do func(*gin.Context, bool) (string, interface{}, error),
	client *redis.Client, duration time.Duration,
	alwaysCache ...bool) func(*gin.Context) {

	if duration < 0 { // don't allow no cache or never expire
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
