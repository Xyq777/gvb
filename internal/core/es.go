package core

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gvb/internal/global"
)

func InitES() *elasticsearch.TypedClient {
	cfg := elasticsearch.Config{
		Addresses: []string{
			global.Config.System.ES.Addr(),
		},
		Username: global.Config.System.ES.Username,
		Password: global.Config.System.ES.Password,
	}
	esClient, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		global.Log.Fatalf("es连接失败 %s", err)
		return nil
	}
	return esClient
}
