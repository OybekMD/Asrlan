package models

type (
	ContentFileCreate struct {
		ContentId int64  `json:"content_id"`
		SoundData string `json:"sound_data"`
		ImageData string `json:"image_data"`
		VideoData string `json:"video_data"`
	}

	ContentFileRequest struct {
		Id int64 `json:"id"`
	}

	ContentFileResponse struct {
		Id        int64  `json:"id"`
		ContentId int64  `json:"content_id"`
		SoundData string `json:"sound_data"`
		ImageData string `json:"image_data"`
		VideoData string `json:"video_data"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
