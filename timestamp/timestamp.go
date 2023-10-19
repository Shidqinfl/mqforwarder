package Timestamp

import (
	"fmt"
	"time"
)

//"2023-08-05T16:51:03+07:00"
//2023/09/08 17:17:38
func GetTime() (timestamp string) {
	// // time.Location = "Jakarta"
	tm := time.Now()
	timestamp = fmt.Sprint(timestamp,  tm.Year(), "-", uint16(tm.Month()),"-", tm.Day(),"T", tm.Hour(),":", tm.Minute(),":", tm.Second(),"+07:00")
	return timestamp
}
func Delay(t time.Duration) {
	time.Sleep(t * time.Millisecond)
}
