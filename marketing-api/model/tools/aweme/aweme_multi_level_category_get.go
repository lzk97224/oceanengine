package aweme

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/enum"
	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// AwemeMultiLevelCategoryGetRequest 查询抖音类目列表 API Request
type AwemeMultiLevelCategoryGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Behaviors 抖音用户行为类型
	Behaviors []enum.Behavior `json:"behaviors,omitempty"`
}

// Encode implement GetRequest interface
func (r AwemeMultiLevelCategoryGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.Behaviors) > 0 {
		buf, _ := json.Marshal(r.Behaviors)
		values.Set("behaviors", string(buf))
	}
	return values.Encode()
}

// AwemeMultiLevelCategoryGetResponse 查询抖音帐号和类目信息 API Response
type AwemeMultiLevelCategoryGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []Category `json:"categories,omitempty"`
	} `json:"data,omitempty"`
}
