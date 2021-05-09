package conf

// LogTransferConf 全局配置
type LogTransferConf struct {
	KafkaConf `ini:"kafka"`
	ESConf    `ini:"es"`
}

// KafkaConf ...
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESConf ...
type ESConf struct {
	Address  string `ini:"address"`
	ChanSize int    `ini:"chan_size"`
	Nums     int    `ini:"nums"`
}
