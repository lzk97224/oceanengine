package creativecomponent

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/enum"
	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// GetRequest 查询组件列表 API Request
type GetRequest struct {
	AdvertiserID uint64        `json:"advertiser_id,omitempty"`
	Page         int           `json:"page,omitempty"`
	PageSize     int           `json:"page_size,omitempty"`
	Filtering    *GetFiltering `json:"filtering,omitempty"`
}

// GetFiltering 过滤条件
type GetFiltering struct {
	// ComponentID 组件ID
	ComponentID uint64 `json:"component_id,omitempty"`
	// ComponentName 组件名称。支持模糊查询
	ComponentName string `json:"component_name,omitempty"`
	// ComponentTypes 组件类型，不传查全部。
	ComponentTypes []enum.ComponentType `json:"component_types,omitempty"`
	// Status 组件审核状态，不传查全部。
	Status []enum.ComponentStatus `json:"status,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// GetResponse 查询组件列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	List     []ComponentInfo `json:"list,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
