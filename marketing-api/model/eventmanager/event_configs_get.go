package eventmanager

import (
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/enum"
	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// EventConfigsGetRequest 获取已创建事件列表 API Request
type EventConfigsGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetID 资产ID
	AssetID uint64 `json:"asset_id,omitempty"`
	// SortType 创建时间排序方式，允许值：DESC 降序、ASC 升序。默认：ASC
	SortType enum.OrderType `json:"sort_type,omitempty"`
}

// Encode implement GetRequest interface
func (r EventConfigsGetRequest) Encode() string {
	values := &url.Values{}
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("asset_id", strconv.FormatUint(r.AssetID, 10))
	if r.SortType != "" {
		values.Add("sort_type", string(r.SortType))
	}
	return values.Encode()
}

// EventConfigsGetResponse 获取已创建事件列表 API Response
type EventConfigsGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 已创建事件列表
		List []EventConfig `json:"event_configs,omitempty"`
	} `json:"data,omitempty"`
}
