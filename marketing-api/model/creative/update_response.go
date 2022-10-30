package creative

import "github.com/lzk97224/oceanengine/marketing-api/model"

// UpdateResponse 更新创意 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json 返回值
type UpdateResponseData struct {
	// Success 更新状态成功的创意ID列表
	Success []uint64 `json:"success,omitempty"`
	// Errors 更新失败的创意列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败信息
type UpdateError struct {
	// CreativeID 更新失败的创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// ErrorMessage 更新失败的原因
	ErrorMessage string `json:"error_message,omitempty"`
}
