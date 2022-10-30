package eventmanager

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/eventmanager"
)

// ShareCancel 事件管理资产取消共享
func ShareCancel(clt *core.SDKClient, accessToken string, req *eventmanager.ShareRequest) ([]eventmanager.ShareError, error) {
	var resp eventmanager.ShareResponse
	if err := clt.Post("v3.0/event_manager/share/cancel/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ErrorList, nil
}
