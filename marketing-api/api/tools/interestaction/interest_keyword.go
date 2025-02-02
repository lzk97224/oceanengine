package interestaction

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/interestaction"
)

// InterestKeyword 兴趣关键词查询
func InterestKeyword(clt *core.SDKClient, accessToken string, req *interestaction.InterestKeywordRequest) ([]interestaction.Object, error) {
	var resp interestaction.InterestKeywordResponse
	err := clt.Get("2/tools/interest_action/interest/keyword", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
