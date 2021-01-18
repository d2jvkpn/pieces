package ginx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	RFC3339ms = "2006-01-02T15:04:05.000Z07:00"

	RESPONSE_User_Key = "user_id"
	RESPONSE_Data_Key = "response_data"
)

type RecordIntf interface {
	GetCode() int
	GetMessage() string
	GetError() error
}

type RecordData struct {
	Time   time.Time `json:"time"`
	Ip     string    `json:"ip"`
	UserId string    `json:"user_id"`

	Method    string `json:"method"`
	Path      string `json:"path"`
	Query     string `json:"query"`
	Referer   string `json:"referer"`
	UserAgent string `json:"user_agent"`

	Status    int           `json:"status"`
	BytesSent int           `json:"bytes_sent"`
	Latency   time.Duration `json:"latency"` // ms
	Level     string        `json:"level"`   // INFO, ERROR, PANIC
	Code      string        `json:"code"`    // NA, NaN, a number
	Message   string        `json:"message"`
	Errorx    error         `json:"errorx"`
}

func NewRespData() *RecordData {
	return &RecordData{
		Time: time.Now(),
		Code: "NA",
	}
}

// time, ip, user_id, method, uri, query, referer, user_agent, bytes_sent, status,
// level, latency, code, message, error
func NewRecord(achive func(*RecordData)) (hf gin.HandlerFunc) {

	hf = func(c *gin.Context) {
		rd, r := NewRespData(), c.Request

		rd.Ip, rd.Method = c.ClientIP(), r.Method
		rd.Path, rd.Query = r.URL.Path, r.URL.RawQuery
		rd.Referer, rd.UserAgent = c.GetHeader("Referer"), c.GetHeader("User-Agent")
		rd.UserId = c.GetString(RESPONSE_User_Key)

		//// achive and recover from panic
		defer func() { achive(rd) }()

		defer func() {
			if intf := recover(); intf == nil {
				return
			} else {
				bts, _ := json.Marshal(intf)
				rd.Errorx = NewErrorx("!!! panic", string(bts))
			}

			rd.Status = http.StatusInternalServerError
			rd.Latency = time.Since(time.Time(rd.Time)) / time.Duration(1_000_000)
			rd.Level, rd.Code = "PANIC", "NaN"
		}()

		//// handle *gin.Context
		c.Next()

		//// process
		var (
			ok  bool
			err error
			ri  RecordIntf
		)

		w := c.Writer
		if user_id := c.GetString(RESPONSE_User_Key); user_id != "" {
			rd.UserId = user_id
		}

		rd.Status, rd.BytesSent = w.Status(), w.Size()
		rd.Latency = time.Since(time.Time(rd.Time)) / time.Duration(1_000_000)

		if rr, _ := c.Get(RESPONSE_Data_Key); rr == nil {
			return
		} else if ri, ok = rr.(RecordIntf); !ok {
			return
		}

		rd.Code, rd.Message = strconv.Itoa(ri.GetCode()), ri.GetMessage()
		if err = ri.GetError(); err != nil {
			rd.Level, rd.Errorx = "ERROR", err
		} else {
			rd.Level = "INFO"
		}
	}

	return hf
}

// print RecordData to stdout in json format
func PrintRecordData(rd *RecordData, levels ...string) func(*RecordData) {
	mp := make(map[string]bool, len(levels))
	for i := range levels {
		mp[levels[i]] = true
	}

	return func(rd *RecordData) {
		if rd == nil || !mp[rd.Level] {
			return
		}
		bts, _ := json.Marshal(rd)
		fmt.Printf("%s\n", bts)
	}
}
