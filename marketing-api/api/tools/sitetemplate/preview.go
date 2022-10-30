package sitetemplate

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/sitetemplate"
)

// Preview 获取模版预览链接
// 通过此接口，用户可以预览通过【基于站点创建模板】接口创建的落地页模板
// 落地页模板的预览链接有效时间为20分钟
func Preview(clt *core.SDKClient, accessToken string, req *sitetemplate.PreviewRequest) (string, error) {
	var resp sitetemplate.PreviewResponse
	if err := clt.Get("2/tools/site_template/preview/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.PreviewURL, nil
}
