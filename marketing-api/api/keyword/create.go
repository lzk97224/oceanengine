package keyword

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/keyword"
)

// Create 创建关键词
// 在指定的ad_id下创建搜索词，可以在原有基础上进行新增。
func Create(clt *core.SDKClient, accessToken string, req *keyword.CreateRequest) (*keyword.ResponseData, error) {
	var resp keyword.Response
	err := clt.Post("2/keyword/create_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
