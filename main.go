package main

import (
	"fmt"
	"holidays/holiday"
)

func main() {
	fmt.Println(holidays.FetchByYear(2019))
}
