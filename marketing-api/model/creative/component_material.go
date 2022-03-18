package creative

import "github.com/bububa/oceanengine/marketing-api/enum"

// ComponentMaterial 组件信息
type ComponentMaterial struct {
	// ComponentID 组件id，通过【查询组件列表】接口获取
	ComponentID string `json:"component_id,omitempty"`
	// ComponentType 组件类型
	ComponentType enum.ComponentType `json:"component_type,omitempty"`
	// MaterialID
	MaterialID string `json:"material_id,omitempty"`
}
