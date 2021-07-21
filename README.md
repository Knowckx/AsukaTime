# AsukaTime

## English
### Time Utils:  
generate Time Range of day/week/month
> go get -u github.com/Knowckx/asukatime

- NewTime() return time.Time with string input
```
ti1, err := NewTime("2021-07-20") 
if err != nil {
    fmt.Println(err)
    return
}  
fmt.Println(ti1) // ti1 is time.Time 2021-07-20 00:00:00

ti2, err := NewTime("2021-07-20 11:11:11") 
if err != nil {
    fmt.Println(err)
    return
}  
fmt.Println(ti2) // ti2 is time.Time 2021-07-20 11:11:11
```

- Get time range of one day
```
input := time.Now()  // 2021-07-20 17:37:21
tr := asukatime.GetDayRange(input)
fmt.Println(tr.Start)  // 2021-07-20 00:00:00
fmt.Println(tr.End)  // 2021-07-20 23:59:59
```

- Get time range of week (monday is the first day of week)
```
input, err := NewTime("2021-07-20") // input is time.Time 2021-07-20 00:00:00
if err != nil {
    fmt.Println(err)
    return
}  
tr := asukatime.GetWeekRange(input)
fmt.Println(tr) // [2021-07-19 00:00:00, 2021-07-25 23:59:59.999]
```



- APIs

|  API   | time  |
|  ----  | ----  |
| GetDayRange  | day |
| GetWeekRange  | week |
| GetMonthRange  | month |
| GetYearRange  | Year |

---



## CN
### 获取一个时间点所在的时间范围
工作中经常遇到这样的需求，生成`当天`，`当周`，`当月`的时间范围，然后拿这个时间范围作为参数去查API，查数据库。

我把常见的时间范围实现了一下，欢迎大家直接调（懒惰是程序员进步的源泉233）

---

输入一个时间点，求该时间点所在的一天，一周，一月的起止时间范围，

> go get -u github.com/Knowckx/asukatime



- 获取某天的时间范围
```
input := time.Now()  // 2021-07-20 17:37:21
tr := asukatime.GetDayRange(input)
fmt.Println(tr.Start)  // 2021-07-20 00:00:00
fmt.Println(tr.End)  // 2021-07-20 23:59:59
```

- 获取某周的时间范围（默认周一是第一天）
```
input, err := NewTime("2021-07-20") // input is time.Time 2021-07-20 00:00:00
if err != nil {
    fmt.Println(err)
    return
}  
tr := asukatime.GetWeekRange(input)
fmt.Println(tr) // [2021-07-19 00:00:00, 2021-07-25 23:59:59.999]
```

- 获取某月的时间范围
```
input := time.Now()  // 2021-07-20 17:37:21
tr := asukatime.GetMonthRange(input)
fmt.Println(tr)  
// [2021-07-01 00:00:00, 2021-07-31 23:59:59.999]
```
