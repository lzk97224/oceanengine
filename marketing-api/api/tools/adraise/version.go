package adraise

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/adraise"
)

// Version “获取起量版本信息”用于获取计划在多次起量过程中产生的起量版本号及对应的起止时间。
func Version(clt *core.SDKClient, accessToken string, req *adraise.VersionRequest) (*adraise.VersionResponseData, error) {
	var resp adraise.VersionResponse
	err := clt.Get("2/tools/ad_raise_version/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
