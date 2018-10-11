package impl

import (
	"gitlab.33.cn/site-service/miner-service-impl/classes/models/db/poolmodel"

	"github.com/pkg/errors"
	"gitlab.33.cn/site-service/miner-service-intf/mine_intf"
	"golang.org/x/net/context"
)

type Context = context.Context

func init() {

}

type PoolImpl struct {
}

func (base PoolImpl) GetList(ctx context.Context, request *mine_intf.PoolListRequest) (*mine_intf.PoolListResponse, error) {
	//var fields []string
	//var sortby []string
	//var order []string
	var err = errors.New("未知错误")
	//var query = make(map[string]string)
	//var limit int64 = 10
	//var offset int64
	if request.PageSize > 0 {
		//limit = int64(request.PageSize)
	}
	if request.Page > 0 {
		//offset = (int64(request.Page) - 1) * limit;
	}
	city := poolmodel.FindOne()
	// err = errors.New("MODEL查询失败")
	// return nil, err
	var poolList []*mine_intf.PoolItem
	poolList = append(poolList, city)
	var response *mine_intf.PoolListResponse
	response.Code = 200
	response.Data = poolList
	return response, err
}
