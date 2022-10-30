package taskraise

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/tools/taskraise"
)

// StatusStop 关闭优选起量任务
func StatusStop(clt *core.SDKClient, accessToken string, req *taskraise.StatusStopRequest) error {
	return clt.Post("2/tools/task_raise/status/stop/", req, nil, accessToken)
}
