package agent

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/agent"
)

// AdvertiserSelect 广告主列表
// 获取代理商下的广告主ID列表，如果存在二级代理，要获取二级代理商列表请使用agent/child_agent/select/接口。如果您需要获取二级代理商下的广告主，需要传递二级代理商id调用此接口即可获得。如果要查看广告主ID的详细信息请参考广告主信息接口。
func AdvertiserSelect(clt *core.SDKClient, accessToken string, req *agent.AdvertiserSelectRequest) (*agent.AdvertiserSelectResponseData, error) {
	var resp agent.AdvertiserSelectResponse
	err := clt.Get("2/agent/advertiser/select", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
