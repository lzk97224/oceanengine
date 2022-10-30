package v3

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	v3 "github.com/lzk97224/oceanengine/marketing-api/model/report/v3"
)

// ProjectGet 项目数据报表
func ProjectGet(clt *core.SDKClient, accessToken string, req *v3.ProjectGetRequest) (*v3.ProjectGetResult, error) {
	var resp v3.ProjectGetResponse
	if err := clt.Get("v3.0/report/project/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
