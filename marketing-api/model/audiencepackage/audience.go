package audiencepackage

import (
	"github.com/lzk97224/oceanengine/marketing-api/enum"
	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// Audience 受众
type Audience struct {
	// RetargetingTags 定向人群包列表，内容为人群包id
	RetargetingTags []uint64 `json:"retargeting_tags,omitempty"`
	// RetargetingTagsExclude 排除人群包列表，内容为人群包id
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// Gender 性别
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄
	Age []enum.AudienceAge `json:"age,omitempty"`
	// AndroidOsv 最低安卓版本，当app_type为"APP_ANDROID"选填,其余情况不填,
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本，当app_type为"APP_IOS"选填,其余情况不填
	IosOsv string `json:"ios_osv,omitempty"`
	// Carrier 运营商
	Carrier []enum.Carrier `json:"carrier,omitempty"`
	// Ac 网络类型
	Ac []string `json:"ac,omitempty"`
	// DeviceBrand 手机品牌
	DeviceBrand []string `json:"device_brand,omitempty"`
	// ArticleCategory 文章分类
	ArticleCategory []string `json:"article_category,omitempty"`
	// ActivateType 新用户(新用户使用头条的时间)
	ActivateType []enum.ActivateType `json:"activate_type,omitempty"`
	// Platform 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载方式的时候，平台可多选。为保证投放效果,平台类型定向PC与移动端互斥
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// AppCategory 老版行为兴趣）APP行为定向——按分类
	AppCategory []uint64 `json:"app_category,omitempty"`
	// AppIDs （老版行为兴趣）APP行为定向——按APP（请注意如果投放的是"应用下载-IOS"不支持设置APP行为定向，请勿传值。）可通过【工具-查询工具-查询应用信息】获取。当app_behavior_target为APP时有值
	AppIDs []uint64 `json:"app_ids,omitempty"`
	// LauchPrice 手机价格,传入价格区间，最高传入11000（表示1w以上）;传值示例 "launch_price": [2000, 11000]，表示2000元以上;
	LaunchPrice []int `json:"launch_price,omitempty"`
	// InterestActionMode 行为兴趣;取值："UNLIMITED"不限,"CUSTOM"自定义,"RECOMMEND"系统推荐。若与自定义人群同时使用，系统推荐("RECOMMEND")不生效;仅推广范围为默认时可填，且不可与老版行为兴趣定向同时填写，否则会报错
	InterestActionMode string `json:"interest_action_mode,omitempty"`
	// ActionScene 行为场景
	ActionScene []enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 用户发生行为天数，当interest_action_mode传CUSTOM时有效
	ActionDays uint `json:"action_days,omitempty"`
	// ActionCategories 行为类目词，当interest_action_mode传CUSTOM时有效
	ActionCategories []uint64 `json:"action_categories,omitempty"`
	// ActionWords 行为关键词，当interest_action_mode传CUSTOM时有效
	ActionWords []uint64 `json:"action_words,omitempty"`
	// InterestCategories 兴趣类目词，当interest_action_mode传CUSTOM时有效
	InterestCategories []uint64 `json:"Interest_categories,omitempty"`
	// InterestWords 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	InterestWords []uint64 `json:"interest_words,omitempty"`
	// BusinessIDs 商圈ID数组，district为"BUSINESS_DISTRICT"时有值
	BusinessIDs []uint64 `json:"business_ids,omitempty"`
	// District 地域;取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限，省市传法："city": [12],"district": "CITY",区县的传法："city": [130102],"district": "COUNTY";暂不支持"海外"
	District string `json:"district,omitempty"`
	// RegionVersion 行政区域版本号
	RegionVersion string `json:"region_version,omitempty"`
	// City 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),当district为"CITY"或"COUNTY"时有值
	City []uint64 `json:"city,omitempty"`
	// LocationType 位置类型;取值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户;当city和district有值时返回值
	LocationType string `json:"location_type,omitempty"`
	// SuperiorPopularityType 媒体定向;
	SuperiorPopularityType string `json:"superior_popularity_type,omitempty"`
	// FlowPackage 定向逻辑
	FlowPackage []uint64 `json:"flow_package,omitempty"`
	// ExcludeFlowPackage 排除定向逻辑
	ExcludeFlowPackage []uint64 `json:"exclude_flow_package,omitempty"`
	// Career 职业选项，详见【附录-职业】
	Career []string `json:"career,omitempty"`
	// DeviceType 设备类型;取值是："MOBILE", "PAD"。缺省表示不限设备类型。穿山甲已经全量，投放范围为默认时需要有白名单权限才可以
	DeviceType []string `json:"device_type,omitempty"`
	// Geolocation 从地图添加(地图位置)
	Geolocation []model.Geolocation `json:"geolocation,omitempty"`
	// AutoExtendEnabled 是否启用智能放量。取值是：0、1。缺省为 0。
	AutoExtendEnabled int `json:"auto_extend_enabled,omitempty"`
	// AutoExtendTarget 可放开定向。当auto_extend_enabled=1 时选填。详见：【附录-可开放定向】。缺省为全不选。
	AutoExtendTarget []string `json:"auto_extend_targets,omitempty"`
	// FilterAwemeAbnormalActive （抖音号推广特有）过滤高活跃用户; 取值：0表示不过滤，1表示过滤
	FilterAwemeAbnormalActive int `json:"filter_aweme_abnormal_active,omitempty"`
	// FilterAwemeFansCount （抖音号推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterAwemeFansCount int64 `json:"filter_aweme_fans_count,omitempty"`
	// AwemeFansNumbers （抖音号推广特有）账号粉丝相似人群（添加抖音账号，会将广告投放给对应账号的相似人群粉丝）
	AwemeFansNumbers []int64 `json:"aweme_fans_numbers,omitempty"`
	// FilterOwnAwemeFans （抖音号推广特有）过滤自己的粉丝; 取值：0表示不过滤，1表示过滤
	FilterOwnAwemeFans int `json:"filter_own_aweme_fans,omitempty"`
	// AwemeFanTimeScope 粉丝互动行为时间范围
	AwemeFanTimeScope string `json:"aweme_fan_time_scope,omitempty"`
	// AwemeFanBehavior 抖音达人互动用户行为类型
	AwemeFanBehavior []enum.Behavior `json:"aweme_fan_behavior,omitempty"`
	// AwemeFanCategories 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanCategories []uint64 `json:"aweme_fan_categories,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanAccounts []uint64 `json:"aweme_fan_accounts,omitempty"`
	// BindInfo 定向包绑定计划信息
	BindInfo []struct {
		// AdID 计划ID
		AdID         uint64 `json:"ad_id,omitempty"`
		AdName       string `json:"ad_name,omitempty"`
		CampaignName string `json:"campaign_name,omitempty"`
		Status       string `json:"status,omitempty"`
	} `json:"bind_info,omitempty"`
}
