package creativeword

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/creativeword"
)

// Delete 删除动态创意词包
// 使用此接口可以删除已创建的动态创意词包。
func Delete(clt *core.SDKClient, accessToken string, req *creativeword.DeleteRequest) (uint64, error) {
	var resp creativeword.DeleteResponse
	err := clt.Post("2/tools/creative_word/delete/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CreativeWordID, nil
}
