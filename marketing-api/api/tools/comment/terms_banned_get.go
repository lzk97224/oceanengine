package comment

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/comment"
)

// TermsBannedGet 获取屏蔽词列表
// 本接口用于获取广告主下的屏蔽词列表。该屏蔽词针对该广告主下的所有评论生效，命中屏蔽词的抖音评论将直接隐藏，不对外进行展示。屏蔽词管理模块目前只针对“抖音”广告位生效。屏蔽词数量最多500个。
func TermsBannedGet(clt *core.SDKClient, accessToken string, req *comment.TermsBannedGetRequest) (*comment.TermsBannedGetResponseData, error) {
	var resp comment.TermsBannedGetResponse
	err := clt.Get("2/tools/comment/terms_banned/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
