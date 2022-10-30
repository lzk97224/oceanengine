package interestaction

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/interestaction"
)

// KeywordSuggest 获取行为兴趣推荐关键词
func KeywordSuggest(clt *core.SDKClient, accessToken string, req *interestaction.KeywordSuggestRequest) ([]interestaction.Object, error) {
	var resp interestaction.KeywordSuggestResponse
	err := clt.Get("2/tools/interest_action/keyword/suggest/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.Keywords, nil
}
