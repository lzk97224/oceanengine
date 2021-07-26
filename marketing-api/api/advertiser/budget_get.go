package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// Budget 获取账户日预算
// 此接口可以获取广告主账号设置的预算类型与预算，可以一次查询100个广告主账号预算；
func Budget(clt *core.SDKClient, accessToken string, advertiserIDs []uint64) (*advertiser.BudgetGetResponse, error) {
	req := &advertiser.BudgetGetRequest{
		AdvertiserIDs: advertiserIDs,
	}
	var resp advertiser.BudgetGetResponse
	err := clt.Get("2/advertiser/budget/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
