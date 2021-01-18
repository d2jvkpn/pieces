package ginx

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Demo(addr string) (err error) {
	var client *redis.Client

	router := gin.Default()

	if client, err = DefaultRedisClient(); err != nil {
		return err
	}

	router.GET("/:name", GinWithRedis(demoDo, client, 10*time.Second))
	return router.Run(addr)
}

func demoDo(c *gin.Context, process bool) (res Res) {
	var (
		name string
		err  error
	)

	//// get key
	if !process {
		name = c.Param("name")
		res.Key = "Demo:" + name
		c.Set("name", name)
		return res
	}

	////
	var year int
	name = c.GetString("name")

	if year, err = strconv.Atoi(c.DefaultQuery("year", "")); err != nil {
		res = NewRes(-1, "year is invalid", err)
		res.Status = http.StatusBadRequest
		return res
	}

	age := time.Now().Year() - year
	if age < 0 {
		res = NewRes(-2, "imporper year", fmt.Errorf("invalid parameter year: %d", year))
		res.Status = http.StatusBadRequest
		return res
	}

	res = NewRes(0, "OK")
	res.Data = gin.H{"name": name, "age": age}
	return res
}
