# 节假日库

> 项目中需要用到节假日查询接口：1. 网络上免费的不稳定 2. 稳定的需要收费，而且带很多不需要的信息

> 基于此，开发一款极度精简的节假日库


## API List

###  FetchAll
> 获取当前公布的从 2010 ~ 至今 的官方节假日安排

### FetchByYear(year int)
> 获取当前公布的节假日安排，按年份查询

### FetchByMonth(year int, month int)
> 获取当前公布的节假日安排，按年份和月份查询

### FetchByChName(year int, name string)
> 获取当前公布的节假日安排，按年份和中文名称查询

### FetchByEnName(year int, name string)
> 获取当前节假日安排，按年份和英文名称查询

### FetchYearHolidayCount(year int)
> 获取某年存在多少天的假期

### IsHoliday(date string)
> 判断某天是否是节假日，某天的格式是："2006/01/02"

### IsWorkDay(date string)
> 判断某天是否是工作日，某天的格式是： "2006/01/02"

### IsWeekDay(date string)
> 判断某天是否是周末，某天的格式是： "2006/01/02"

### FetchLunarDay()
> 所有农历节日

### GuessLunarByMonth(month int)
> 通过月份查看农历节日

---
<使用愉快...>

## 英文文档

[英文文档](en_README.md)

## TODO

- 追踪具体信息来源：自动监控、自动抓取、自动解析
