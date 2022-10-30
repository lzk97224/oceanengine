package oauth

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/oauth"
)

// AccessToken 获取Access Token
// Access-Token是调用接口时，操作指定广告账户的身份凭证，有效期为24小时
// Refresh-Token用于生成新access_token和refresh_token并且刷新时效达到续期的目的
func AccessToken(clt *core.SDKClient, authCode string) (*oauth.AccessTokenResponseData, error) {
	req := &oauth.AccessTokenRequest{
		AppId:     clt.AppID,
		Secret:    clt.Secret,
		GrantType: "auth_code",
		AuthCode:  authCode,
	}
	var resp oauth.AccessTokenResponse
	err := clt.Post("oauth2/access_token/", req, &resp, "")
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
