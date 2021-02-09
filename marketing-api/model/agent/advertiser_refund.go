package agent

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type AdvertiserRefundRequest struct {
	AdvertiserID uint64  `json:"advertiser_id,omitempty"` // 广告主ID
	AgentID      uint64  `json:"agent_id,omitempty"`      // 代理商ID
	TransferType string  `json:"transfer_type,omitempty"` // 转账类型（新增字段）. 枚举：CASH：现金，GRANT：赠款
	Amount       float64 `json:"amount,omitempty"`        // 金额,单位(元),最低转账金额500元
}

func (r AdvertiserRefundRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type AdvertiserRefundResponse struct {
	model.BaseResponse
	Data *AdvertiserRefundResponseData `json:"data,omitempty"`
}

type AdvertiserRefundResponseData struct {
	TransactionSeq string `json:"transaction_seq,omitempty"` // 交易序列号
}
