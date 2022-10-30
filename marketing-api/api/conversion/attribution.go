package conversion

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/conversion"
)

// Attribution 电话转化回传
func Attribution(clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsV1Post("attribution", req, &resp)
}
