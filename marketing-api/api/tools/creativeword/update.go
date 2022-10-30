package creativeword

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/creativeword"
)

// Update 更新动态创意词包
// 针对已创建的动态创意词包进行修改，可修改词包名称、默认词、替换词三部分。
func Update(clt *core.SDKClient, accessToken string, req *creativeword.UpdateRequest) (uint64, error) {
	var resp creativeword.UpdateResponse
	err := clt.Post("2/tools/creative_word/update/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CreativeWordID, nil
}
