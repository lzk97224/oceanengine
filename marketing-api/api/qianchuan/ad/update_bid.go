package ad

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/ad"
)

// UpdateBid 更新计划出价
// 用于更新广告计划的出价 bid，一次可以处理10个计划
// 注意： 修改的出价不能大于当前预算
func UpdateBid(clt *core.SDKClient, accessToken string, req *ad.UpdateBidRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post("v1.0/qianchuan/ad/bid/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
