package tools

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools"
)

// CountryInfo 查询国家/区域信息
func CountryInfo(clt *core.SDKClient, accessToken string, req *tools.CountryInfoRequest) (*tools.CountryInfoResponseData, error) {
	var resp tools.CountryInfoResponse
	if err := clt.Get("2/tools/country/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
