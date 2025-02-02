package creative

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/creative"
)

// Get 获取账户下创意列表
// 此接口用于获取生成后创意列表，包括：
// 自定义创意：计划创建后即可获取
// 程序化创意：创意过审后即可获取
func Get(clt *core.SDKClient, accessToken string, req *creative.GetRequest) (*creative.GetResponseData, error) {
	var resp creative.GetResponse
	err := clt.Get("v1.0/qianchuan/creative/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
