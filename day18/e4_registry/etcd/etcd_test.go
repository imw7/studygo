package etcd

import (
	"context"
	"registry"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	registryInst, err := registry.InitRegistry(context.TODO(), "etcd",
		registry.WithAddrs([]string{"127.0.0.1:2379"}),
		registry.WithTimeout(time.Second),
		registry.WithRegistryPath("/my/etcd/"),
		registry.WithHeartbeat(5),
	)
	if err != nil {
		t.Errorf("init registry failed, err:%v\n", err)
		return
	}
	service := &registry.Service{
		Name: "comment_service",
	}
	service.Nodes = append(service.Nodes,
		&registry.Node{
			IP:   "127.0.0.1",
			Port: 8801,
		},
		&registry.Node{
			IP:   "127.0.0.2",
			Port: 8801,
		},
	)
	if err = registryInst.Register(context.TODO(), service); err != nil {
		return
	}
	go func() {
		time.Sleep(time.Second * 10)
		service.Nodes = append(service.Nodes,
			&registry.Node{
				IP:   "127.0.0.3",
				Port: 8801,
			},
			&registry.Node{
				IP:   "127.0.0.4",
				Port: 8801,
			},
		)
		if err = registryInst.Register(context.TODO(), service); err != nil {
			return
		}
	}()
	select {}
}
