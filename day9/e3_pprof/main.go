package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// runtime_pprof https://imw7.github.io/p/performance_optimisation.html

// 一段有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Println("recv from chan, value:", v)
		default:
			// 通过分析发现大部分CPU资源被18行占用，我们分析出select语句中的default没有内容会导致上面的
			// `case v := <-c:`一直执行。我们在default分支添加一行`time.time.Sleep(time.Second)`即可。
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var isCPUPprof bool // 是否开启CPU profile的标志位
	var isMemPprof bool // 是否开启内存 profile的标志位

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof") // 在当前路径下创建一个cpu.pprof文件
		if err != nil {
			fmt.Println("create cpu pprof failed, err:", err)
			return
		}
		_ = pprof.StartCPUProfile(file) // 往文件中记录CPU profile信息
		defer func() {
			pprof.StopCPUProfile()
			_ = file.Close()
		}()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Println("create mem pprof failed, err:", err)
			return
		}
		_ = pprof.WriteHeapProfile(file)
		_ = file.Close()
	}
}
