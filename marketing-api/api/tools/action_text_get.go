package tools

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools"
)

// ActionTextGet 行动号召字段内容获取
// 获取行动号召字段内容，注意：结合附加创意类型以及广告主行业参数可以查询出更多细纬度的行动号召内容。
func ActionTextGet(clt *core.SDKClient, accessToken string, req *tools.ActionTextGetRequest) ([]string, error) {
	var resp tools.ActionTextGetResponse
	if err := clt.Get("2/tools/action_text/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
