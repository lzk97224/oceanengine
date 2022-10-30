package ad

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/ad"
)

// Update 更新计划
// 计划包括了计划名称、投放范围、投放目标、用户定向、预算与出价，对于其中的概念解释可以参考：【广告计划】
// 未展示的字段，表示不可以进行修改（修改时注意字段的描述）
func Update(clt *core.SDKClient, accessToken string, req *ad.UpdateRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post("2/ad/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
