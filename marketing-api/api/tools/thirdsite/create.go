package thirdsite

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/thirdsite"
)

// Create 创建第三方落地页站点
// 通过此接口，用户可以创建第三方落地页站点，创建成功后接口会返回"code_0"。
func Create(clt *core.SDKClient, accessToken string, req *thirdsite.CreateRequest) (uint64, error) {
	var resp thirdsite.CreateResponse
	err := clt.Post("2/tools/third_site/create/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	if resp.Data == nil {
		return 0, nil
	}
	return resp.Data.SiteID, nil
}
