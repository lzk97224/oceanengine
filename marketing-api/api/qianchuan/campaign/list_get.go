package campaign

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/campaign"
)

// ListGet 获取广告组
// 本接口用于获取账户下的广告组列表
func ListGet(clt *core.SDKClient, accessToken string, req *campaign.ListGetRequest) (*campaign.ListGetResponseData, error) {
	var resp campaign.ListGetResponse
	err := clt.Get("v1.0/qianchuan/campaign_list/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
