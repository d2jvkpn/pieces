package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
)

type Res struct {
	Key    string `json:"key"`    // redis cacke key
	Status int    `json:"status"` // http response StatueCode
	Err    error  `json:"-"`

	ResData
}

type ResData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewRes(code int, message string, errs ...error) (rd Res) {
	rd.Status = http.StatusOK

	if len(errs) > 0 {
		rd.Err = errs[0]
	}
	if message == "" && rd.Err != nil {
		message = rd.Err.Error()
	}

	rd.Code, rd.Message = code, message
	return rd
}

func (rd *Res) Error() string { // implememt error interface
	if rd.Err != nil {
		return rd.Err.Error()
	}
	return "<nil>"
}

func DefaultRedisClient() (client *redis.Client, err error) {
	client = redis.NewClient(
		&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0},
	)

	statusCmd := client.Ping(context.TODO())
	if err := statusCmd.Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func GinWithRedis(do func(*gin.Context, bool) Res,
	client *redis.Client, duration time.Duration,
	alwaysCache ...bool) func(*gin.Context) {

	if duration < 0 { // no cache never expire
		return nil
	}

	return func(c *gin.Context) {
		var (
			key string
			bts []byte
			err error
			mp  map[string]string
			cmd *redis.StringStringMapCmd
			rd  Res
		)

		if rd = do(c, false); rd.Err != nil {
			c.JSON(rd.Status, rd.ResData) // http.StatusBadRequest
			return
		}
		key = rd.Key //! important

		defer func() {
			c.Header("StatusCode", strconv.Itoa(rd.Status)) // strconv.Itoa(http.StatusOK)
			// c.Header("Status", http.StatusText(http.StatusOK))
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.Writer.Write(bts)
			return
		}()

		// *redis.StringStringMapCmd
		cmd = client.HGetAll(context.TODO(), rd.Key)
		if err = cmd.Err(); err == nil {
			if mp, err = cmd.Result(); len(mp) > 0 && err == nil {
				// println(">>> cache")
				rd.Status, _ = strconv.Atoi(mp["status"])
				bts = []byte(mp["resdata"])
				return
			}
		}

		hmset := func() {
			bts, _ = json.Marshal(rd.ResData)
			client.HMSet(context.TODO(), key, map[string]interface{}{
				"status":  rd.Status,
				"resdata": bts,
			})

			client.Expire(context.TODO(), key, duration)
		}

		// println(">>> process")
		if rd = do(c, true); rd.Err != nil { // proccess failed
			bts, _ = json.Marshal(rd.ResData)
			if len(alwaysCache) > 0 && alwaysCache[0] {
				hmset() // avoid cache penetration
			}
			return
		}

		hmset()
		return
	}
}
