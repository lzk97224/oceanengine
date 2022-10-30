package tools

import (
	"github.com/lzk97224/oceanengine/marketing-api/enum"
	"github.com/lzk97224/oceanengine/marketing-api/model"
	"github.com/lzk97224/oceanengine/marketing-api/util/query"
)

// BidSuggestRequest 建议日预算及预期成本 API Request
type BidSuggestRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Pricing 出价类型，查看【附录-出价类型】
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// CampaignID 广告组ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AdID 广告ID，修改广告时需要传
	AdID uint64 `json:"ad_id,omitempty"`
	// BidMode 出价方式，手动&自动
	// 允许值："SUGGEST"、"AUTO_BID"
	// - 手动获取到的是建议出价
	// - 自动获取到的是建议日预算和预期成本
	BidMode string `json:"bid_mode,omitempty"`
	// BudgetMode 广告预算类型(创建后不可修改), 详见【附录-预算类型】
	// 允许值: "BUDGET_MODE_DAY","BUDGET_MODE_TOTAL"
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 广告预算(出价方式为CPC、CPM、CPV时，不少于100元；出价方式为OCPM、OCPC时，不少于300元。单次预算修改幅度不小于100元,日修改预算不超过20次)
	// 取值范围: "≥ 0"
	Budget float64 `json:"budget,omitempty"`
	// ScheduleType 广告投放时间类型, 详见【附录-广告投放时间类型】
	// 允许值:"SCHEDULE_FROM_NOW","SCHEDULE_START_END"
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// FlowControlMode 广告投放速度类型, 详见【附录-广告投放速度类型】
	// 允许值: "FLOW_CONTROL_MODE_FAST","FLOW_CONTROL_MODE_SMOOTH"
	FlowControlMode enum.FlowControlMode `json:"flow_control_mode,omitempty"`
	// ConvertID 转化id，可通过【工具模块-OCPC广告创建转化查询】查询可用id
	ConvertID uint64 `json:"convert_id,omitempty"`
	// RetargetingType 定向人群包类型,详见【附录-定向人群包类型】,即将下线
	// 允许值: "RETARGETING_INCLUDE", "RETARGETING_EXCLUDE","RETARGETING_NONE"
	RetargetingType enum.RetargetingType `json:"retargeting_type,omitempty"`
	// RetargetingTags 当选择使用人群包定向时填写，内容为人群包id，即将下线
	RetargetingTags []uint64 `json:"retargeting_tags,omitempty"`
	// RetargetingTagsInclude 定向人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsInclude []uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// Gender 性别
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄
	Age []enum.AudienceAge `json:"age,omitempty"`
	// AndroidOsv 最低安卓版本，当app_type为"APP_ANDROID"选填,其余情况不填,
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本，当app_type为"APP_IOS"选填,其余情况不填
	IosOsv string `json:"ios_osv,omitempty"`
	// Ac 网络类型
	Ac []string `json:"ac,omitempty"`
	// Carrier 运营商
	Carrier []enum.Carrier `json:"carrier,omitempty"`
	// DeviceBrand 手机品牌
	DeviceBrand []string `json:"device_brand,omitempty"`
	// ArticleCategory 文章分类
	ArticleCategory []string `json:"article_category,omitempty"`
	// ActivateType 用户首次激活时间, 详见【附录-用户首次激活时间】
	// 允许值:"WITH_IN_A_MONTH","ONE_MONTH_2_THREE_MONTH","THREE_MONTH_EAILIER"
	ActivateType []enum.ActivateType `json:"activate_type,omitempty"`
	// Platform 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载方式的时候，平台可多选。为保证投放效果,平台类型定向PC与移动端互斥
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// District 地域;取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限，省市传法："city": [12],"district": "CITY",区县的传法："city": [130102],"district": "COUNTY";暂不支持"海外"
	District string `json:"district,omitempty"`
	// City 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),当district为"CITY"或"COUNTY"时有值
	City []uint64 `json:"city,omitempty"`
	// BusinessIDs 商圈ID数组，district为"BUSINESS_DISTRICT"时有值
	BusinessIDs []uint64 `json:"business_ids,omitempty"`
	// LocationType 位置类型;取值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户;当city和district有值时返回值
	LocationType string `json:"location_type,omitempty"`
	// AdTags （老版行为兴趣）兴趣分类,如果传"空数组"表示不限，如果"数组传0"表示系统推荐,如果按兴趣类型传表示自定义
	AdTags []uint64 `json:"ad_tags,omitempty"`
	// InterestTags （老版行为兴趣）兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。
	InterestTags []uint64 `json:"interest_tags,omitempty"`
	// AppBehaviorTarget （老版行为兴趣）APP行为; 取值：NONE不限，CATEGORY按分类，APP按APP
	AppBehaviorTarget string `json:"app_behavior_target,omitempty"`
	// AppCategory 老版行为兴趣）APP行为定向——按分类
	AppCategory []uint64 `json:"app_category,omitempty"`
	// AppIDs （老版行为兴趣）APP行为定向——按APP（请注意如果投放的是"应用下载-IOS"不支持设置APP行为定向，请勿传值。）可通过【工具-查询工具-查询应用信息】获取。当app_behavior_target为APP时有值
	AppIDs []uint64 `json:"app_ids,omitempty"`
	// SuperiorPopularityType 媒体定向;
	SuperiorPopularityType string `json:"superior_popularity_type,omitempty"`
	// FlowPackage 定向逻辑
	FlowPackage []uint64 `json:"flow_package,omitempty"`
	// ExcludeFlowPackage 排除定向逻辑
	ExcludeFlowPackage []uint64 `json:"exclude_flow_package,omitempty"`
	// IncludeCustomActions 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},]， dpa_local_audience为1时有值; day可选范围:1, 7, 14, 28, 60, 90, 120, 180。
	IncludeCustomActions []interface{} `json:"include_custom_actions,omitempty"`
	// ExcludeCustomActions 排除人群包((DPA推广目的特有,格式举例{"days": 7, "code": 1002},]，day可选范围: 1, 7, 14, 28, 60, 90, 120, 180。
	ExcludeCustomActions []interface{} `json:"exclude_custom_actions,omitempty"`
}

// Encode implement GetRequest interface
func (r BidSuggestRequest) Encode() string {
	values, _ := query.Values(r)
	return values.Encode()
}

// BidSuggestResponse 建议日预算及预期成本 API Response
type BidSuggestResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *BidSuggest `json:"data,omitempty"`
}

// BidSuggest 建议日预算及预期成本
type BidSuggest struct {
	// SuggestedBid 建议出价
	SuggestedBid float64 `json:"suggested_bid,omitempty"`
	// BidHigh50 竞争力超过50%的计划（手动出价获取）
	BidHigh50 float64 `json:"bid_high_50,omitempty"`
	// BidHigh90 竞争力超过90%的计划（手动出价获取）
	BidHigh90 float64 `json:"bid_high_90,omitempty"`
	// SmartBidSuggestBudget 建议日预算（自动出价获取）
	SmartBidSuggestBudget float64 `json:"smart_bid_suggest_budget,omitempty"`
	// SmartBudgetRange 建议日预算范围（自动出价获取）
	SmartBudgetRange []float64 `json:"smart_budget_range,omitempty"`
	// SmartBidRange 预期成本范围（自动出价获取）
	SmartBidRange []float64 `json:"smart_bid_range,omitempty"`
}
