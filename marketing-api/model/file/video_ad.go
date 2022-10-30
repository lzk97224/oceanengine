package file

import (
	"strconv"

	"github.com/lzk97224/oceanengine/marketing-api/model"
)

// VideoAdRequest 上传广告视频 API Request
type VideoAdRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// VideoSignature 视频的md5值(用于服务端校验)
	VideoSignature string `json:"video_signature,omitempty"`
	// VideoFile 视频文件; 允许格式：mp4、mpeg、3gp、avi（10s超时限制）
	VideoFile *model.UploadField `json:"video_file,omitempty"`
	// Filename 素材的文件名，可自定义素材名，不传则默认取文件名，最长255个字符; 注：若同一素材已进行上传，重新上传不会改名
	Filename string `json:"filename,omitempty"`
}

// Encode implement UploadRequest interface
func (r VideoAdRequest) Encode() []model.UploadField {
	fields := []model.UploadField{
		{
			Key:   "advertiser_id",
			Value: strconv.FormatUint(r.AdvertiserID, 10),
		},
	}
	if r.VideoFile != nil {
		filename := r.VideoFile.Value
		if filename == "" {
			filename = "file"
		}
		fields = append(fields, model.UploadField{
			Key:   "video_signature",
			Value: r.VideoSignature,
		}, model.UploadField{
			Key:    "video_file",
			Value:  filename,
			Reader: r.VideoFile.Reader,
		})
	}
	if r.Filename != "" {
		fields = append(fields, model.UploadField{
			Key:   "filename",
			Value: r.Filename,
		})
	}
	return fields
}

// VideoAdResponse 上传广告视频 API Response
type VideoAdResponse struct {
	model.BaseResponse
	Data *Video `json:"data,omitempty"`
}
