package agent

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/agent"
)

// AdvertiserRecharge 代理商转账
// 通过此接口可给代理商下管理的广告主或二级代理商进行转账操作。
// 提醒：1月15日代理商平台和API都新增了转账类型字段，代理商如果使用了转账退款接口，请尽快调整上传最新的transfer_type字段，历史转账会按一定的规则和比例进行转账，新功能上线后需要您实际选择是希望转现金还是赠款，在1月31日之前如果不传transfer_type我们将按历史逻辑处理，在31日后如果不传transfer_type将默认优先转现金，请您尽快调整。
func AdvertiserRecharge(clt *core.SDKClient, accessToken string, req *agent.AdvertiserRechargeRequest) (string, error) {
	var resp agent.AdvertiserRechargeResponse
	err := clt.Post("2/agent/advertiser/recharge/", req, &resp, accessToken)
	if err != nil {
		return "", err
	}
	return resp.Data.TransactionSeq, nil
}
