package dpa

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/dpa"
)

// TemplateGet 获取DPA模板
func TemplateGet(clt *core.SDKClient, accessToken string, req *dpa.TemplateGetRequest) (*dpa.TemplateGetResponseData, error) {
	var resp dpa.TemplateGetResponse
	if err := clt.Get("2/dpa/template/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
