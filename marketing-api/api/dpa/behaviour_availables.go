package dpa

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/dpa"
)

// BehaviourAvailables 获取DPA可用行为
func BehaviourAvailables(clt *core.SDKClient, accessToken string, req *dpa.BehaviourAvailablesRequest) ([]dpa.Behaviour, error) {
	var resp dpa.BehaviourAvailablesResponse
	if err := clt.Get("2/dpa/behaviour/availables/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
