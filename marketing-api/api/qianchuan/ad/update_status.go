package ad

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/ad"
)

// UpdateStatus 更新计划状态
// 通过此接口可对计划做启用/暂停/删除操作，一次可以处理10个计划
// 对于删除的计划不能再进行状态操作，否则会报错
// 可能存在部分计划更新成功的情况，更新失败的计划id及原因将会在返回体中返回
func UpdateStatus(clt *core.SDKClient, accessToken string, req *ad.UpdateStatusRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post("v1.0/qianchuan/ad/status/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
