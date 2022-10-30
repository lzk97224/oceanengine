package dpa

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/dpa"
)

// MetaGet 获取商品库元信息
func MetaGet(clt *core.SDKClient, accessToken string, req *dpa.MetaGetRequest) ([]dpa.Meta, error) {
	var resp dpa.MetaGetResponse
	err := clt.Get("2/dpa/meta/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
