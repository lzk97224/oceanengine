package union

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/union"
)

// FlowPackageReport 查看rit数据
// 查看穿山甲广告位数据，包括付费率、消耗等数据，默认查询历史30天数据。
func FlowPackageReport(clt *core.SDKClient, accessToken string, req *union.FlowPackageReportRequest) (*union.FlowPackageReportData, error) {
	var resp union.FlowPackageReportResponse
	if err := clt.Get("2/tools/union/flow_package/report/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
