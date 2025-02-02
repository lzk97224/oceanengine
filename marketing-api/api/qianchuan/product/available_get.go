package product

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/product"
)

// AvailableGet 获取可投商品列表
// 获取千川账户下可投商品列表接口
func AvailableGet(clt *core.SDKClient, accessToken string, req *product.AvailableGetRequest) (*product.AvailableGetResponseData, error) {
	var resp product.AvailableGetResponse
	err := clt.Get("v1.0/qianchuan/product/available/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
