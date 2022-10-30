package privativeword

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/privativeword"
)

// CampaignUpdate 设置组否定词
func CampaignUpdate(clt *core.SDKClient, accessToken string, req *privativeword.AdUpdateRequest) (uint64, error) {
	var resp privativeword.AdUpdateResponse
	err := clt.Post("2/privative_word/campaign/update", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.CampaignID, nil
}
