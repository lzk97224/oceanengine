package agent

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/agent"
)

// AdvertiserUpdate 修改广告主
// 修改广告主信息，可更改内容包括账户名称、联系人、手机号码、固定电话，除此之外其他内容不允许修改。
func AdvertiserUpdate(clt *core.SDKClient, accessToken string, req *agent.AdvertiserUpdateRequest) (*agent.AdvertiserUpdateResponseData, error) {
	var resp agent.AdvertiserUpdateResponse
	err := clt.Post("2/agent/advertiser/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
