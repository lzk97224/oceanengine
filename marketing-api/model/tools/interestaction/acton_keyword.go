package interestaction

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/enum"
	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// ActionKeywordRequest 行为关键词查询 API Request
type ActionKeywordRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// QueryWords 关键词
	QueryWords string `json:"query_words,omitempty"`
	// ActionScene 行为场景，查询目标为行为时必填，兴趣不生效;允许值：E-COMMERCE、NEWS、APP
	ActionScene []enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 行为天数，查询目标为行为时必填，兴趣不生效; 允许值：7, 15, 30, 60, 90, 180, 365
	ActionDays int `json:"action_days,omitempty"`
}

// Encode implement GetRequest interface
func (r ActionKeywordRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("query_words", r.QueryWords)
	scene, _ := json.Marshal(r.ActionScene)
	values.Set("action_scene", string(scene))
	values.Set("action_days", strconv.Itoa(r.ActionDays))
	return values.Encode()
}

// ActionKeywordResponse 行为关键词查询 API Response
type ActionKeywordResponse struct {
	model.BaseResponse
	// Data json返回值
	Data ActionKeywordResponseData `json:"data,omitempty"`
}

// ActionKeywordResponseData json返回值
type ActionKeywordResponseData struct {
	List []Object `json:"list,omitempty"`
}
