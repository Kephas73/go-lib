package util

import "time"

const (
    TimestampFormat = `2006-01-02T15:04:05.999Z07:00`
)

func GetTimestampData() string {
    t := time.Now()
    location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
    if err != nil {
        return t.Format(TimestampFormat)
    }
    return t.In(location).Format(TimestampFormat)
}
