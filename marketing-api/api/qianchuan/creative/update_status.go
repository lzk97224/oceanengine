package creative

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/creative"
)

// UpdateStatus 更新创意状态
// 此接口用于批量修改创意状态；
// 不能操作已删除的创意状态
// 程序化创意不能进行单独操作，需要从计划纬度进行操作，否则会报错
func UpdateStatus(clt *core.SDKClient, accessToken string, req *creative.UpdateStatusRequest) (*creative.UpdateResponseData, error) {
	var resp creative.UpdateResponse
	err := clt.Post("v1.0/qianchuan/creative/status/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
