package impl

import (
	"gitlab.33.cn/site-service/miner-service-impl/classes/models/db/poolmodel"

	"gitlab.33.cn/site-service/miner-service-intf/mine_intf"
	"golang.org/x/net/context"
)

type Context = context.Context

func init() {

}

type PoolImpl struct {
}

func (base PoolImpl) GetList(ctx context.Context, request *mine_intf.PoolListRequest) (*mine_intf.PoolListResponse, error) {
	poolItem := poolmodel.FindPool()
	var poolList []*mine_intf.PoolItem
	poolList = append(poolList, &poolItem)
	var response mine_intf.PoolListResponse
	response.Data = poolList
	response.Code = 201
	return &response, nil
}
