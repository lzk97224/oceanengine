package tools

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools"
)

// SearchBidRatioGet 获取快投推荐出价系数
func SearchBidRatioGet(clt *core.SDKClient, accessToken string, req *tools.SearchBidRatioGetRequest) (float64, error) {
	var resp tools.SearchBidRatioGetResponse
	if err := clt.Get("2/tools/search_bid_ratio/get/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.Ratio, nil
}
