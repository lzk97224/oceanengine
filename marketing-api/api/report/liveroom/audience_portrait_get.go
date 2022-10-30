package liveroom

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/report/liveroom"
)

// AudiencePortraitGet 直播受众分析报表
// 此接口用于进行直播间的受众分析、获取直播间用户画像数据；
// 用户画像维度包含性别、年龄范围、省份、城市、用户设备平台；
// 支持获取分用户画像维度的直播间观看数、关注数、加入粉丝团数、打赏次数、商品订单数、客单价等；
func AudiencePortraitGet(clt *core.SDKClient, accessToken string, req *liveroom.Request) (*liveroom.ResponseData, error) {
	var resp liveroom.Response
	err := clt.Get("2/report/live_room/audience/portrait/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
