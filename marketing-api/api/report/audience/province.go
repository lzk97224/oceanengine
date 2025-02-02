package audience

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/report/audience"
)

// Province 省级数据
// 获取省级受众定向数据，包括你所投放省份的总花费、展示数、点击数等数据。
// 受众数据不支持获取当天的数据；
func Province(clt *core.SDKClient, accessToken string, req *audience.Request) ([]audience.ResponseData, error) {
	var resp audience.Response
	err := clt.Get("2/report/audience/province/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
