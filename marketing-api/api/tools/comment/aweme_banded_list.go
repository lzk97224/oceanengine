package comment

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/comment"
)

// AwemeBandedList 获取屏蔽用户列表
// 获取屏蔽用户列表用于获取广告主绑定的抖音账号（第一个绑定的抖音账号）屏蔽的抖音id列表或昵称关键词列表。
func AwemeBandedList(clt *core.SDKClient, accessToken string, req *comment.AwemeBandedListRequest) (*comment.AwemeBandedListResponseData, error) {
	var resp comment.AwemeBandedListResponse
	err := clt.Get("2/tools/comment/aweme_banded/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
