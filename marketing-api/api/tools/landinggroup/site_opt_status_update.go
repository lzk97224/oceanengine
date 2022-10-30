package landinggroup

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/landinggroup"
)

// SiteOptStatusUpdate 更新落地页组站点状态
// 通过此接口，用户可以修改落地页组站点的启用状态。
func SiteOptStatusUpdate(clt *core.SDKClient, accessToken string, req *landinggroup.SiteOptStatusUpdateRequest) error {
	return clt.Post("2/tools/landing_group/site_opt_status/update/", req, nil, accessToken)
}
