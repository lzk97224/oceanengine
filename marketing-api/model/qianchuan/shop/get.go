package shop

import (
	"encoding/json"
	"net/url"

	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// GetRequest 获取店铺账户信息 API Request
type GetRequest struct {
	// ShopIDs 店铺id
	ShopIDs []uint64 `json:"shop_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := &url.Values{}
	ids, _ := json.Marshal(r.ShopIDs)
	values.Set("shop_ids", string(ids))
	return values.Encode()
}

// GetResponse 获取店铺账户信息 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List shop账号列表
		List []Shop `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// Shop shop账号
type Shop struct {
	// ID 店铺id
	ID uint64 `json:"shop_id,omitempty"`
	// Name 店铺名称
	Name string `json:"shop_name,omitempty"`
}
