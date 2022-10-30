package comment

import (
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// TermsBannedGetRequest 获取屏蔽词列表 API Request
type TermsBannedGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r TermsBannedGetRequest) Encode() string {
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

// TermsBannedGetResponse 获取屏蔽词列表 API Response
type TermsBannedGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *TermsBannedGetResponseData `json:"data,omitempty"`
}

// TermsBannedGetResponseData json 返回值
type TermsBannedGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Terms 屏蔽词列表
	Terms []string `json:"terms,omitempty"`
}
