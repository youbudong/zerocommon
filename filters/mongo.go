package filters

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FixRegexString(str string) string {
	//正则匹配出现的特殊字符串
	fbsArr := []string{"\\", "$", "(", ")", "*", "+", ".", "[", "]", "?", "^", "{", "}", "|"}
	for _, ch := range fbsArr {
		if strContainers := strings.Contains(str, ch); strContainers {
			str = strings.Replace(str, ch, "\\"+ch, -1)
		}
	}
	return str
}

// ParseQueryDate 解析时间
func ParseQueryDate(searchTime string) (startTime, endTime time.Time, err error) {
	var times []string
	if searchTime != "" {
		times = strings.Split(searchTime, ",")
	}
	if len(times) > 0 {
		startTime, err = time.Parse("2006-01-02", times[0])
		if err != nil {
			return
		}
	}
	if len(times) > 1 {
		endTime, err = time.Parse("2006-01-02", times[1])
		if err != nil {
			return
		}
	}
	return
}

// FilterDate 过滤时间
func FilterDateRange(searchTime string, key string, filter primitive.M) {
	start, end, _ := ParseQueryDate(searchTime)
	if !start.IsZero() && start == end {
		filter[key] = bson.M{"$gte": start, "$lt": end.AddDate(0, 0, 1)}
	}
	if end.IsZero() && !start.IsZero() {
		filter[key] = bson.M{"$gte": start}
	}
	if !end.IsZero() && start.IsZero() {
		filter[key] = bson.M{"$lt": end.AddDate(0, 0, 1)}
	}
	if !start.IsZero() && !end.IsZero() && start != end {
		filter[key] = bson.M{"$gte": start, "$lt": end.AddDate(0, 0, 1)}
	}
}

// FilterKey 过滤
func FilterValue(value, key string, filter primitive.M) {
	if value != "" {
		filter[key] = value
	}
}

// FilterKeyword 过滤关键字
func FilterKeyword(keyword string, keys []string, filter primitive.M) {
	if keyword != "" {
		filter["$or"] = bson.A{}
		for _, key := range keys {
			filter["$or"] = append(filter["$or"].(bson.A), primitive.M{key: primitive.Regex{Pattern: FixRegexString(keyword), Options: "i"}})
		}
	}
}

// FilterKeywordWithFilterA 过滤关键字使用filterA
func FilterKeywordWithFilterA(keyword string, keys []string, filterA primitive.A) {
	if keyword != "" {
		for _, key := range keys {
			filterA = append(filterA, primitive.M{key: primitive.Regex{Pattern: FixRegexString(keyword), Options: "i"}})
		}
	}
}
