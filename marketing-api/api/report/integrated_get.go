package report

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/report"
)

// IntegratedGet 多合一数据报表接口
// 数据报表使用方式：报表、监控通知、自动化调整
func IntegratedGet(clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponseData, error) {
	var resp report.IntegratedResponse
	err := clt.Get("2/report/integrated/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
