package ad

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/qianchuan/ad"
)

// RejectReason 获取计划审核建议
// 通过此接口用于获取广告计划纬度的审核意见
func RejectReason(clt *core.SDKClient, accessToken string, req *ad.RejectReasonRequest) ([]ad.RejectReasonList, error) {
	var resp ad.RejectReasonResponse
	err := clt.Get("v1.0/qianchuan/ad/reject_reason/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
