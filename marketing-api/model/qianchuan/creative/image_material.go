package creative

import "github.com/lzk97224/oceanengine/marketing-api/enum/qianchuan"

// ImageMaterial 图片素材
type ImageMaterial struct {
	// ID 素材唯一标识
	ID uint64 `json:"id,omitempty"`
	// ImageIDs 图片ID列表
	ImageIDs []string `json:"image_ids,omitempty"`
	// ImageID 图片ID
	ImageID string `json:"image_id,omitempty"`
	// ImageMode 素材类型，见附录-枚举值
	ImageMode qianchuan.ImageMode `json:"image_mode,omitempty"`
	// IsAutoGenerate 是否为派生创意标识，1：是，0：不是
	IsAutoGenerate int `json:"is_auto_generate,omitempty"`
}
