package registry

// 抽象出一些结构体
// Node：单个节点的结构体，包含 id ip port weight(权重)
// Service：里面有服务名，还有节点列表，一个服务多台服务器支撑

// Service 抽象服务
type Service struct {
	// 服务名
	Name string `json:"name"`
	// 节点列表
	Nodes []*Node `json:"nodes"`
}

// Node 单个服务节点的抽象
type Node struct {
	Id     string `json:"id"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}
