package audience

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/report/audience"
)

// InterestActionList 行为兴趣数据
// 此接口用于获取：行为+兴趣数据，如果你在计划的受众定向中设置了行为/兴趣定向，那么在这能看到所选的行为/兴趣指标所对应的数据，包括花费、展示、点击等。
// 行为兴趣数据不支持获取当天的数据；
func InterestActionList(clt *core.SDKClient, accessToken string, req *audience.ListRequest) (*audience.ListResponseData, error) {
	var resp audience.ListResponse
	err := clt.Get("2/report/audience/interest_action/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
