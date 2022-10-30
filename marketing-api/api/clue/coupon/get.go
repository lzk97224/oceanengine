package coupon

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/clue/coupon"
)

// Get 获取卡券列表
func Get(clt *core.SDKClient, accessToken string, req *coupon.GetRequest) (*coupon.GetResponseData, error) {
	var resp coupon.GetResponse
	if err := clt.Get("2/clue/coupon/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
