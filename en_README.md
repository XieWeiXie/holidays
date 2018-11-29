# Holidays API

> A simple package of chinese holidays


## Installation

```
go get 	"github.com/wuxiaoxiaoshen/holidays/holiday"
```

## Simple Example

```
package main
import (
	"fmt"
	"github.com/wuxiaoxiaoshen/holidays/holiday"
)

func main() {

	h := holidays.FetchByYear(2019)
	fmt.Println(h)
}
```
```
[{2018/12/30 2019/01/01 元旦 New Year\'s Day} {2019/02/04 2019/02/10 春节 Spring Festival} {2019/04/05 2019/04/07 清明节 Tomb-sweeping Day} {2019/04/29 2019/05/01 劳动节 Labour Day} {2019/06/07 2019/06/09 端午节 Dragon Boat Festival} {2019/09/13 2019/09/15 国庆节 National Day} {2019/10/01 2019/10/07 中秋节 Mid-autumn Festival}]
```

```go
package main
import (
	"fmt"
	"github.com/wuxiaoxiaoshen/holidays/holiday"
)
func main(){
	h2 := holidays.FetchMonthHolidayCount(2018, 10)
	fmt.Println(h2)
}
```
```
7
```

```
package main
import (
	"fmt"
	"github.com/wuxiaoxiaoshen/holidays/holiday"
)
func main(){
	h3 := holidays.IsHoliday("2018/10/01")
	fmt.Println(h3)
}
```

```
true
```
---

## FetchAll

> get all holidays in chinese (2010 ~ 2019)

## FetchByChName(year int, name string)

> get holidays in chinese search by year and chinese name

## FetchByEnName(year int, name string)

> get holidays in chinese search by year and english name

## FetchByMonth(year int, month int)

> get holidays in chinese search by year and month

## FetchByYear(year int)

> get holidays in chinese search by year

## FetchMonthHolidayCount(year int, month int)

> get count days in chinese holidays search by year and month

## FetchYearHolidayCount(year int)

> get count days in chinese holidays search by year

## IsHoliday

> is date  with format(2006/01/02) holiday?

## IsWeekDay

> is date with format(2006/01/02) weekday?

## IsWorkDay

> is date with format(2006/01/02) workday?

