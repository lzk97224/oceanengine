package file

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/file"
)

// VideoAwemeGet 获取抖音号下的视频
// 获取抖音号下已有的视频素材，支持传入商品id，过滤拉取包含待推广商品的视频
// 1、仅自定义创意支持选择抖音号视频，程序化创意不支持
// 2、短视频带货场景下，仅支持选择抖音号下已关联相应推广商品的视频
// 3、由于素材库存在分钟级延迟，上传素材后请勿立即获取并创建计划
func VideoAwemeGet(clt *core.SDKClient, accessToken string, req *file.VideoAwemeGetRequest) (*file.VideoAwemeGetResponseData, error) {
	var resp file.VideoAwemeGetResponse
	err := clt.Get("2/file/video/aweme/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
