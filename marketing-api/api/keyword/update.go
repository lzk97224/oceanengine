package keyword

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/keyword"
)

// Update 更新关键词
// 根据keyword_id更新搜索词，可批量更新。不可更改word，如需要，可通过删除接口与创建接口实现。
func Update(clt *core.SDKClient, accessToken string, req *keyword.UpdateRequest) (*keyword.ResponseData, error) {
	var resp keyword.Response
	err := clt.Post("2/keyword/update_v2/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
