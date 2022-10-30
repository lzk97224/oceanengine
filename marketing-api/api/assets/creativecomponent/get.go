package creativecomponent

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/assets/creativecomponent"
)

// Get 查询组件列表 https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1696710673645580
func Get(clt *core.SDKClient, accessToken string, req *creativecomponent.GetRequest) (*creativecomponent.GetResponseData, error) {
	var resp creativecomponent.GetResponse
	if err := clt.Get("2/assets/creative_component/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
