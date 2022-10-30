package asynctask

import (
	"github.com/lzk97224/oceanengine/marketing-api/core"
	"github.com/lzk97224/oceanengine/marketing-api/model/report/asynctask"
)

// Download 下载任务结果
// 任务结果只保存 15 天。
// range_from 和 range_to 之间的关系: range_from < file_size。 range_from <= range_to。
// range_to > file_size -1 是允许的。
// 以文件形式进行返回，指标参数同实时接口保持一致
// 如果结果为空，那么返回"empty content"
func Download(clt *core.SDKClient, accessToken string, req *asynctask.DownloadRequest) ([]byte, error) {
	return clt.GetBytes("2/async_task/download/", req, accessToken)
}
