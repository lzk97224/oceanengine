package interestaction

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/interestaction"
)

// Id2Word 兴趣行为类目关键词id转词
func Id2Word(clt *core.SDKClient, accessToken string, req *interestaction.Id2WordRequest) (*interestaction.Id2WordResponseData, error) {
	var resp interestaction.Id2WordResponse
	err := clt.Get("2/tools/interest_action/id2word/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
