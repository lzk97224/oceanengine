package audience

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/report/audience"
)

// Gender 性别数据
// 获取性别受众本数据，可以查看受众中男性和女性分别的花费和展现、点击情况数据。
// 受众数据不支持获取当天的数据；
func Gender(clt *core.SDKClient, accessToken string, req *audience.Request) ([]audience.ResponseData, error) {
	var resp audience.Response
	err := clt.Get("2/report/audience/gender/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
