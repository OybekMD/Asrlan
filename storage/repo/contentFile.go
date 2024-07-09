package repo

import "context"

type ContentFile struct {
	Id        int64  `json:"id"`
	ContentId int64  `json:"content_id"`
	SoundData string `json:"sound_data"`
	ImageData string `json:"image_data"`
	VideoData string `json:"video_data"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ContentFileStorageI interface {
	Create(ctx context.Context, badge *ContentFile) (*ContentFile, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*ContentFile, error)
}
