package adpreview

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/adpreview"
)

// QrcodeGet 获取广告预览二维码
func QrcodeGet(clt *core.SDKClient, accessToken string, req *adpreview.QrcodeGetRequest) (*adpreview.QrcodeGetResponseData, error) {
	var resp adpreview.QrcodeGetResponse
	err := clt.Get("2/tools/ad_preview/qrcode_get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
