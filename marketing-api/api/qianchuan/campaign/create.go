package campaign

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/campaign"
)

// Create 创建广告组
// 一个广告组最多容纳500条广告
// 创建广告计划引用一个满500条的广告组，系统会自动创建一个新的广告组
func Create(clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (uint64, error) {
	var resp campaign.CreateResponse
	err := clt.Post("v1.0/qianchuan/campaign/create/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
