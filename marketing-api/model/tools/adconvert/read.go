package adconvert

import (
	"net/url"
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// ReadRequest 查询转化目标详细信息 API Request
type ReadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ConvertID 转化id，其中较小数值convert_id为预定义转化
	ConvertID uint64 `json:"convert_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReadRequest) Encode() string {
	values := &url.Values{}
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("convert_id", strconv.FormatUint(r.ConvertID, 10))
	return values.Encode()
}

// ReadResponse 查询转化目标详细信息 API Response
type ReadResponse struct {
	model.BaseResponse
	Data *Convert `json:"data,omitempty"`
}
