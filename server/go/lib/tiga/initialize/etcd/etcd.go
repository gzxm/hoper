package etcd

import (
	"context"
	"github.com/actliboy/hoper/server/go/lib/tiga/initialize"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConfig clientv3.Config

func (conf *EtcdConfig) generate() *clientv3.Client {
	client, _ := clientv3.New((clientv3.Config)(*conf))
	resp, _ := client.Get(context.Background(), initialize.InitKey)
	initialize.InitConfig.UnmarshalAndSetV2(resp.Kvs[0].Value)
	return client
}

func (conf *EtcdConfig) Generate() interface{} {
	return conf.generate()
}

type Eecd struct {
	*clientv3.Client
	Conf EtcdConfig
}

func (e *Eecd) Config() interface{} {
	return &e.Conf
}

func (e *Eecd) SetEntity(entity interface{}) {
	if client, ok := entity.(*clientv3.Client); ok {
		e.Client = client
	}
}
