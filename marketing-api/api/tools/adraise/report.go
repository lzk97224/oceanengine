package adraise

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/adraise"
)

// Report “获取一键起量报告”用于获取开启过一键起量功能的计划，在一键起量功能生效期间产生的报告数据；接口支持获取总体报告和分时报告，需要根据计划的起量版本号及起止时间进行数据获取。
func Report(clt *core.SDKClient, accessToken string, req *adraise.ReportRequest) (*adraise.ReportResponseData, error) {
	var resp adraise.ReportResponse
	err := clt.Get("2/tools/ad_raise_result/get_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
