package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

func demoDo(c *gin.Context, process bool) (key string, data interface{}, err error) {
	if !process {
		key = c.Param("name")
		if key == "" {
			return "", nil, fmt.Errorf("no name provided")
		}
		key = "Demo2:" + key
		c.Set("key", key)
		return key, nil, nil
	}

	intf, _ := c.Get("key")
	key, _ = intf.(string)

	var year int
	yearStr, ok := c.GetQuery("year")
	if !ok {
		err = NewResData(-1, "year not provided", fmt.Errorf("parameter year not available"))
		return key, nil, err
	}

	if year, err = strconv.Atoi(yearStr); err != nil {
		return key, nil, NewResData(-2, "year is invalid", err)
	}

	age := time.Now().Year() - year
	if year <= 0 || age < 0 {
		err = NewResData(-3, "imporper year", fmt.Errorf("invalid parameter year: %d", year))
		return key, nil, err
	}

	resp := NewResData(0, "OK")
	resp.Data = gin.H{"name": key, "age": age}
	return key, resp, nil
}
