package utils

import (
	sf "github.com/yzw/snowflake"
	"math/rand"
	"strconv"
	"time"
)

var (
	snowflake *sf.Snowflake
)

// 判断当前是不是月初2天
func Early() bool {
	t := time.Now()
	return (t.Month() != t.Add(-24*time.Hour).Month() || t.Month() != t.Add(-48*time.Hour).Month())
}

// 判断当前是不是月末2天
func MonthEnd() bool {
	t := time.Now()
	return (t.Month() != t.Add(24*time.Hour).Month() || t.Month() != t.Add(48*time.Hour).Month())
}

// 获取随机字符串
func GetStr() (str string) {
	str = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
	return
}

// 初始化 snowflake
func GetSnowFlake(workerid, centerid int64) (err error) {
	snowflake, err = sf.NewSnowflake(workerid, centerid)
	if err != nil {
		return
	}
	_, err = snowflake.Id()
	return
}

// 获取唯一id
func GetTradeId() (str string, err error) {
	id, err := snowflake.Id()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(id, 10), nil
}
