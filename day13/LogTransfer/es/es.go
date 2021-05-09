package es

import (
	"LogTransfer/utils"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

// 初始化ES，准备接收kafka那边发来的数据

// LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

func Init(addr string, chanSize, nums int) (err error) {
	if strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL("http://" + addr))
	if err != nil {
		return
	}
	fmt.Println("connect to es succeed.")
	ch = make(chan *LogData, chanSize)
	for i := 0; i < nums; i++ {
		go sendToES()
	}
	return
}

// SendToESChan 发送数据到ES
func SendToESChan(msg *LogData) {
	ch <- msg
}

func sendToES() {
	for {
		select {
		case msg := <-ch:
			// fmt.Println(indexStr, data)
			ip, err := utils.GetOutboundIP()
			if err != nil {
				fmt.Println("get outbound ip failed, err:", err)
				return
			}
			put, err := client.Index().Index(msg.Topic).Type(ip).BodyJson(msg).Do(context.Background())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Indexed %s to index %s, type %s\n", put.Id, put.Index, put.Type)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
