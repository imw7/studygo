package main

import (
	"fmt"
	"time"
)

// 时间

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time.Unix()
	ret := time.Unix(1591526331, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())
	// 时间间隔
	fmt.Println(time.Second)
	// now + 24h
	fmt.Println(now.Add(24 * time.Hour))
	// Sub 两个时间相减
	nextYear, err := time.Parse("2006-01-02", "2021-01-01")
	if err != nil {
		fmt.Println("parse time failed, err: ", err)
		return
	}
	d := nextYear.Sub(now)
	fmt.Println(d)
	// 定时器
	// timer := time.Tick(time.Second)
	// for t:=range timer{
	//	fmt.Println(t) // 1秒钟执行一次
	// }

	// 格式化时间 把语言中时间对象，转换成字符串类型的时间
	// 2020-11-13
	fmt.Println(now.Format("2006-01-02"))
	// 2020/02/03 11:55:02
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 2020/02/03 11:55:02 AM
	fmt.Println(now.Format("2006/01/02 15:04:05 PM"))
	// 2020/02/03 11:55:02.342
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))
	// 按照对应的格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2000-11-13")
	if err != nil {
		fmt.Println("parse time failed, err:", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// Sleep
	n := 3 // int
	fmt.Println("开始sleep了")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("%d秒钟过去了\n", n)
	time.Sleep(5 * time.Second)
}

// 时区
func f2() {
	now := time.Now() // 本地的时间
	fmt.Println(now)
	// 明天的这个时间
	// 按照指定格式去解析一个字符串格式的时间
	_, _ = time.Parse("2006-01-02 15:04:05", "2050-01-01 10:00:30")
	// 按照东八区的时区和格式解析一个字符串格式的时间
	// 根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load loc failed, err:", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2050-01-01 10:00:30", loc)
	if err != nil {
		fmt.Print("parse time failed, err:", err)
		return
	}
	fmt.Println(timeObj)
	// 时间相减
	duration := timeObj.Sub(now)
	fmt.Println(duration)
}

func main() {
	f1()
	f2()
}
