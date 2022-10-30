package eventmanager

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/eventmanager"
)

// AvailableEventsGet 获取可创建事件列表
// 此接口用于查询资产下支持创建的事件列表
func AvailableEventsGet(clt *core.SDKClient, accessToken string, req *eventmanager.AvailableEventsGetRequest) ([]eventmanager.EventConfig, error) {
	var resp eventmanager.AvailableEventsGetResponse
	err := clt.Get("2/event_manager/available_events/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
