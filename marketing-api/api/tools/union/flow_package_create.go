package union

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/union"
)

// FlowPackageCreate 创建穿山甲流量包
func FlowPackageCreate(clt *core.SDKClient, accessToken string, req *union.FlowPackageCreateRequest) (uint64, error) {
	var resp union.FlowPackageCreateResponse
	if err := clt.Post("2/tools/union/flow_package/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.FlowPackageID, nil
}
