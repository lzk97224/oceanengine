package agent

import (
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// ChildAgentSelectRequest 二级代理商列表 API Request
type ChildAgentSelectRequest struct {
	// AdvertiserID 代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码.默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量.默认值: 100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ChildAgentSelectRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// ChildAgentSelectResponse 二级代理商列表 API Response
type ChildAgentSelectResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ChildAgentSelectResponseData `json:"data,omitempty"`
}

// ChildAgentSelectResponseData json返回值
type ChildAgentSelectResponseData struct {
	// ChildAgentIDs 二级代理商ID列表
	ChildAgentIDs []uint64 `json:"child_agent_ids,omitempty"`
}
