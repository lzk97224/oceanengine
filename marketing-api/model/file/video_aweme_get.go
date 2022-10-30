package file

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// VideoAwemeGetRequest 获取抖音号下的视频 API Request
type VideoAwemeGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 需拉取视频的抖音号
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// Filtering 视频过滤条件
	Filtering *VideoAwemeGetFiltering `json:"filtering,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小， 默认值：10，取值范围1-100
	PageSize int `json:"page_size,omitempty"`
}

// VideoAwemeGetFiltering 筛选条件
type VideoAwemeGetFiltering struct {
	// AwemeURL 抖音视频链接，可在抖音端上通过【分享】-【复制链接】获取
	AwemeURL string `json:"aweme_url,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoAwemeGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	if r.Filtering != nil {
		filter, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filter))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// VideoAwemeGetResponse 获取抖音号下的视频 API Response
type VideoAwemeGetResponse struct {
	model.BaseResponse
	Data *VideoAwemeGetResponseData `json:"data,omitempty"`
}

// VideoAwemeGetResponseData json返回值
type VideoAwemeGetResponseData struct {
	// List 视频列表
	List []Video `json:"video_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
