package impl

import (
	"github.com/pkg/errors"
	"gitlab.33.cn/site-service/miner-service-intf/mine_intf"
	"golang.org/x/net/context"
	"miner-service-impl/models"
)

type Context = context.Context

func init() {

}

type PoolImpl struct {
}

func (base PoolImpl) GetList(ctx context.Context, request *mine_intf.PoolListRequest) (*mine_intf.PoolListResponse, error) {
	var fields []string
	var sortby []string
	var order []string
	var err = errors.New("未知错误")
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	if request.PageSize > 0 {
		limit = int64(request.PageSize)
	}
	if request.Page > 0 {
		offset = (int64(request.Page) - 1) * limit;
	}
	l, err := models.GetAllPool(query, fields, sortby, order, offset, limit)
	if l != nil {

	}
	if err != nil {
		err = errors.New("MODEL查询失败")
		return nil, err
	}
	var poolList [] *mine_intf.PoolItem
	var response *mine_intf.PoolListResponse
	response.Code = 200
	response.Data = poolList
	return response, err
}
