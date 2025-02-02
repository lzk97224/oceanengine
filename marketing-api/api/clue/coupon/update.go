package coupon

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/clue/coupon"
)

// Update 编辑卡券
// 编辑卡券接口。 只有deliver_start、deliver_end、global_limit、user_limit允许编辑。 卡券开始后，所有的数据都不允许修改。 status字段控制卡券的上下线，只支持：NORMAL：上线，OFFLINE：下线
func Update(clt *core.SDKClient, accessToken string, req *coupon.UpdateRequest) error {
	return clt.Post("2/clue/coupon/update/", req, nil, accessToken)
}
