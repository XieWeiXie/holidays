package main

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/holidays/holiday"
)

func main() {

	h := holidays.FetchByYear(2019)
	fmt.Println(h)

	h2 := holidays.FetchMonthHolidayCount(2018, 10)
	fmt.Println(h2)

	h3 := holidays.IsHoliday("2018/10/01")
	fmt.Println(h3)
}
