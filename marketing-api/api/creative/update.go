package creative

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/creative"
)

// Update 修改创意信息
// 此接口用于全量更新广告创意：当计划下有创意时，需要通过此接口进行全量更新与创建创意（可以先通过创意详细信息接口获取全部信息在进行更新）
// 创意包括了首选位置、基础创意、创意分类、监测链接等信息，对于其中的概念解释可以参考：【广告创意】
// 如果创建创意遇到问题，可通过 常见问题 来解决；
// 如果计划ID下已有创意信息，需要使用update_v2接口进行修改或者新增创意素材，否则会报错
// 全量更新解释：
// 全量更新即为最新一次更新将覆盖前一次的所有信息，例如：如果已有三个创意A、B、C，更新提交了创意B、C、D，则意味着覆盖删除创意A，新增了D，保留（或修改）了B、C；
// 对于新建的创意，不需要上传创意ID，会生成新的创意ID；对于您希望保留的创意，不要忘记上传创意ID，否则将被覆盖；
// 注意：
// 每个计划下程序化创意和自定义创意为二选一，且无法修改；
// 程序化创意: 最多10个标题、12个图片素材和10个视频素材;如果创建的是程序化创意（程序化创意实际会按照传入的title_list和image_list进行组合，对于效果不好的组合无法通过审核，获取到的都是审核通过的创意），只有在审核之后才会获取到创意数据与创意id；
// 自定义创意限制计划下不能超过10个创意，且每日最多创建500个自定义创意。
// 素材类型：不同广告未要求素材类型不同,其中每一种素材类型都有自己的规格，请传入符合要求的素材，否则会报错！
// 其中视频的时长需要>=4s，否则会报错！
// 监测链接：当在计划纬度设置了转化id，如果在创建创意时不传监测链接，会自动获取转化id里监测链接；如果在创建（更新）创意时传入对应的监测链接，会取传入的监测链接，但是对于应用下载推广，即便主动传入点击监测链接，也会取转化id监测链接；
// 创建之后：素材类型、附加创意类型、品牌主页、数据发送方式不能进行修改；
// 对于不打算传的字段，不要传“”或者null，传了会校验!!!
// 触发审核
// 在修改用户端可见的内容包括标题、图片/视频、来源、附加创意、落地页链接等时会触发审核；
// 对于广告位的修改有以下情况也会触发审核：①选择了网盟广告位然后增加头条广告位；②选择了头条广告位然后增加抖音广告位；这两种情况都会触发审核。其他情况比如选了头条广告位再增加火山、西瓜不会触发审核。；
func Update(clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (*creative.CreateResponseData, error) {
	var resp creative.CreateResponse
	if err := clt.Post("2/creative/update_v2/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
