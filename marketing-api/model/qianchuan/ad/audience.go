package ad

import (
	"github.com/lzk97224/oceanengine/marketing-api/enum"
	"github.com/lzk97224/oceanengine/marketing-api/enum/qianchuan"
)

// Audience 定向设置
type Audience struct {
	// AudienceMode 人群定向模式，当promotion_way为 SIMPLE时返回，枚举值：AUTO智能推荐、CUSTOM自定义
	AudienceMode enum.AudienceMode `json:"audience_mode,omitempty"`
	// District 地域定向类型，配合city字段使用，允许值：CITY：省市，COUNTY：区县，NONE：不限；默认值：NONE
	District string `json:"district,omitempty"`
	// City 具体定向的城市列表，当 district 为COUNTY，city 为必填，枚举值详见【附件-city.json】；省市传法：city: [12]，district: CITY；区县的传法：city: [130102]，district: COUNTY
	City []uint64 `json:"city,omitempty"`
	// LocationType 地域定向的用户状态类型，当 district 为COUNTY，CITY为必填，允许值：CURRENT：正在该地区的用户，HOME：居住在该地区的用户，TRAVEL；到该地区旅行的用户，ALL：该地区内的所有用户
	LocationType string `json:"location_type,omitempty"`
	// Gender 允许值: GENDER_FEMALE：女性，GENDER_MALE：男性，NONE： 不限
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄，详见【附录-受众年龄区间】；允许值：AGE_BETWEEN_18_23, AGE_BETWEEN_24_30, AGE_BETWEEN_31_40, AGE_BETWEEN_41_49, AGE_ABOVE_50
	Age enum.AudienceAge `json:"age,omitempty"`
	// AwemeFanBehaviors 抖音达人互动用户行为类型
	AwemeFanBehaviors []string `json:"aweme_fan_behaviors,omitempty"`
	// AwemeFanBehaviorDays 抖音达人互动用户行为天数
	AwemeFanBehaviorDays string `json:"aweme_fan_behavior_days,omitempty"`
	// AwemeFanCategories 抖音达人分类ID列表
	AwemeFanCategories []uint64 `json:"aweme_fan_categories,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表
	AwemeFanAccounts []uint64 `json:"aweme_fan_accounts,omitempty"`
	// AutoExtendEnabled 是否启用智能放量
	AutoExtendEnabled *int `json:"auto_extend_enabled,omitempty"`
	// AutoExtendTargets 可放开定向列表
	AutoExtendTargets []qianchuan.AutoExtendTarget `json:"auto_extend_targets,omitempty"`
	// Platform 投放平台列表
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// SmartInterestAction 行为兴趣意向定向模式
	SmartInterestAction string `json:"smart_interest_action,omitempty"`
	// ActionScene 行为场景
	ActionScene []enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 用户发生行为天数
	ActionDays int `json:"action_days,omitempty"`
	// ActionCategories 行为类目词
	ActionCategories []uint64 `json:"action_categories,omitempty"`
	// ActionWords 行为关键词
	ActionWords []uint64 `json:"action_words,omitempty"`
	// InterestCategories 兴趣类目词
	InterestCategories []uint64 `json:"interest_categories,omitempty"`
	// InterestWords 兴趣关键词
	InterestWords []uint64 `json:"interest_words,omitempty"`
	// Ac 网络类型
	Ac []string `json:"ac,omitempty"`
	// RetargetingTagsInclude 定向人群包id列表
	RetargetingTagsInclude []uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包id列表
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// LivePlatformTags 直播带货平台精选人群包
	LivePlatformTags []enum.LivePlatformTag `json:"live_platform_tags,omitempty"`
}
