package v3

import (
	"time"
)

// Dimensions 维度数据
type Dimensions struct {
	// ProjectID 项目ID，分组条件包含 STAT_GROUP_BY_APP_PROJECT_ID 时返回
	ProjectID uint64 `json:"project_id,omitempty"`
	// PromotionID 广告ID，分组条件包含 STAT_GROUP_BY_APP_PROMOTION_ID 时返回
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// StatTimeHour 细分到每个小时的数据
	StatTimeHour int `json:"stat_time_hour,omitempty"`
	// StatTimeDay 细分到自然天的数据
	StatTimeDay int `json:"stat_time_day,omitempty"`
	// StatTimeWeek 细分到自然周的数据，周一到周日为一整周
	StatTimeWeek int `json:"stat_time_week,omitempty"`
	// StatTimeMonth 细分到自然月的数据，例如2022-3，表示是3月1日到3月31日
	StatTimeMonth time.Month `json:"stat_time_month,omitempty"`
	// CdpProjectID 项目ID
	CdpProjectID uint64 `json:"cdp_project_id,omitempty"`
	// CdpProjectName 项目的名称
	CdpProjectName string `json:"cdp_project_name,omitempty"`
	// LandingType 对应项目的推广目的
	LandingType int `json:"landing_type,omitempty"`
	// ExternalAction 对应项目的转化目标
	ExternalAction string `json:"external_action,omitempty"`
	// Pricing 对应项目的计费类型
	Pricing int `json:"pricing,omitempty"`
	// DeepExternalAction 对应项目的深度转化目标
	DeepExternalAction int `json:"deep_external_action,omitempty"`
	// AdPlatformCdpProjectDownloadType 在项目中设置的下载方式
	AdPlatformCdpProjectDownloadType int `json:"ad_platform_cdp_project_download_type,omitempty"`
	// AdPlatformCdpProjectDownloadURL 在项目中设置的下载链接
	AdPlatformCdpProjectDownloadURL string `json:"ad_platform_cdp_project_download_url,omitempty"`
	// AdPlatformCdpProjectDownloadActionTrackURL 在项目中设置的监测链接
	AdPlatformCdpProjectDownloadActionTrackURL string `json:"ad_platform_cdp_project_action_track_url,omitempty"`
	// CdpPromotionID 广告ID
	CdpPromotionID uint64 `json:"cdp_promotion_id,omitempty"`
	// CdpPromotionName 对应广告的名称
	CdpPromotionName string `json:"cdp_promotion_name,omitempty"`
	// AdPlatformCdpPromotionBid 在广告中设置的出价
	AdPlatformCdpPromotionBid float64 `json:"ad_platform_cdp_promotion_bid,omitempty"`
	// AdPlatformCdpPromotionDeepCpaBid 在广告中设置的深度转化出价
	AdPlatformCdpPromotionDeepCpaBid float64 `json:"ad_platform_cdp_promotion_roi_goal,omitempty"`
	// AppCode 您所投放的广告数据中，对应的首选广告位
	AppCode int `json:"app_code,omitempty"`
	// PackageName 您在项目中设置的应用包包名
	PackageName string `json:"package_name,omitempty"`
	// Gender 您所投放的广告数据中，对应的用户性别。无法识别的性别数据会显示为“其他”
	Gender int `json:"gender,omitempty"`
	// Age 您所投放的广告数据中，对应的用户年龄。无法识别的年龄数据会显示为“其他”
	Age string `json:"age,omitempty"`
	// Ac 您所投放的广告数据中，对应的用户网络环境。无法识别的网络数据会显示为“其他”
	Ac int `json:"ac,omitempty"`
	// Platform 您所投放的广告数据，对应的操作系统平台，无法识别的操作系统会显示为“其他”
	Platform string `json:"platform,omitempty"`
	// ProvinceName 您所投放的广告数据，对应的用户省份信息
	ProvinceName string `json:"province_name,omitempty"`
	// CityName 您所投放的广告数据，对应的用户城市信息
	CityName string `json:"city_name,omitempty"`
	// MaterialID 素材的ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// AdPlatformMaterialContent 对应的素材标题、视频及图片内容
	AdPlatformMaterialContent string `json:"ad_platform_material_content,omitempty"`
	// ImageMode 对应的素材类型，包括标题、大图横图、竖版视频等
	ImageMode int `json:"image_mode,omitempty"`
	// AdPlatformMaterialName 素材为视频素材时，对应的视频名称
	AdPlatformMaterialName string `json:"ad_platform_material_name,omitempty"`
}
