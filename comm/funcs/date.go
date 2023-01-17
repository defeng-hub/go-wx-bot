package funcs

import (
	"fmt"
	"time"

	"github.com/Lofanmi/chinese-calendar-golang/calendar"
)

const DefaultDateFormat = "2006-01-02"

func getYearDay(year int, date string) string {
	return fmt.Sprintf(`%d-%s`, year, date)
}

// 获取两个时间相差的天数
func GetDiffDaysSolar(curDate, futureDate string) (dd int) {
	curD, _ := time.ParseInLocation(DefaultDateFormat, getYearDay(time.Now().Year(), curDate), time.Local)
	furD, _ := time.ParseInLocation(DefaultDateFormat, getYearDay(time.Now().Year(), futureDate), time.Local)

	// 如果是负数说明已经过去了，加一年再计算
	dd = int(furD.Sub(curD).Hours() / 24)
	if dd > 0 {
		return dd
	}

	furD, _ = time.ParseInLocation(DefaultDateFormat,
		getYearDay(time.Now().AddDate(1, 0, 0).Year(), futureDate), time.Local)

	return int(furD.Sub(curD).Hours() / 24)
}

func GetDiffDaysLunar(curDate, futureDate string, lunarMonth, lunarDay int64) (dd int) {
	curD, _ := time.ParseInLocation(DefaultDateFormat, getYearDay(time.Now().Year(), curDate), time.Local)
	furD, _ := time.ParseInLocation(DefaultDateFormat, getYearDay(time.Now().Year(), futureDate), time.Local)

	// 如果是负数说明已经过去了，加一年再计算
	//fmt.Println(curD.String(), furD.String(), int(furD.Sub(curD).Hours()/24))
	dd = int(furD.Sub(curD).Hours() / 24)
	if dd > 0 {
		return dd
	}

	futureDate = getLunar2SolarDate(int64(time.Now().AddDate(1, 0, 0).Year()),
		lunarMonth, lunarDay)
	furD, _ = time.ParseInLocation(DefaultDateFormat,
		getYearDay(time.Now().AddDate(1, 0, 0).Year(), futureDate), time.Local)

	return int(furD.Sub(curD).Hours() / 24)
}

// TODO 这里可能不准,需要计算闰月
func getLunar2SolarDate(year, month, day int64) string {
	c := calendar.ByLunar(year, month, day, 0, 0, 0, false)
	return fmt.Sprintf("%02d-%02d", c.Solar.GetMonth(), c.Solar.GetDay())
}

func RemainingDays() int {
	curD, _ := time.ParseInLocation(DefaultDateFormat, time.Now().Format("2006-01-02"), time.Local)
	furD, _ := time.ParseInLocation(DefaultDateFormat, fmt.Sprintf("%d-12-31", time.Now().Year()), time.Local)

	return int(furD.Sub(curD).Hours() / 24)
}
