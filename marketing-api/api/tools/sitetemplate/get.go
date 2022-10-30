package sitetemplate

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/sitetemplate"
)

// Get 获取站点模版列表
// 通过此接口，用户可以获取通过【基于站点创建模板】接口创建的落地页模板。
func Get(clt *core.SDKClient, accessToken string, req *sitetemplate.GetRequest) (*sitetemplate.GetResponseData, error) {
	var resp sitetemplate.GetResponse
	if err := clt.Get("2/tools/site_template/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
