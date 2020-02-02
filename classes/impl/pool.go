package impl

import (
	"github.com/fogetu/miner_service_impl/classes/models/db/pool"
	"github.com/fogetu/miner_service_intf/mine_intf"
	"golang.org/x/net/context"
)

type Context = context.Context

func init() {

}

type PoolImpl struct {
}

func (base PoolImpl) GetList(ctx context.Context, request *mine_intf.PoolListRequest) (*mine_intf.PoolListResponse, error) {
	poolItem := pool.FindPool()
	var poolList []*mine_intf.PoolItem
	poolList = append(poolList, &poolItem)
	var response mine_intf.PoolListResponse
	response.Data = poolList
	response.Code = 201
	return &response, nil
}
