package advertiser

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/advertiser"
)

// AvatarGet 获取广告主头像信息
func AvatarGet(clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.AvatarGetResponseData, error) {
	req := &advertiser.AvatarGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.AvatarGetResponse
	err := clt.Get("2/advertiser/avatar/get", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
