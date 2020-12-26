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

	AlwayCache bool // avoid cache penetration
}

type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Err     error       `json:"err"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResData(code int, message string, err error) (rd *ResData) {
	if message == "" && err != nil {
		message == err.Err.Error()
	}
	return &ResData{Code: code, Message: message, Err: err}
}

func (rd *ResData) Error() string { // implememt error interface
	if rd.Err != nil {
		return rd.Err.Error()
	}
	return rd.Message // in general, message is ok when code == 0
}

func GinWithRedis(hdl *GinHandler, client *redis.Client) func(*gin.Context) {
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
			if hdl.AlwayCache && hdl.Duration > 0 {
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
