package repo

import "context"

type Social struct {
	LocationName  string `json:"location_name"`
	LocationUrl   string `json:"location_url"`
	EducationName string `json:"education_name"`
	EducationUrl  string `json:"education_url"`
	TelegramName  string `json:"telegram_name"`
	TelegramUrl   string `json:"telegram_url"`
	TwitterName   string `json:"twitter_name"`
	TwitterUrl    string `json:"twitter_url"`
	InstagramName string `json:"instagram_name"`
	InstagramUrl  string `json:"instagram_url"`
	YoutubeName   string `json:"youtube_name"`
	YoutubeUrl    string `json:"youtube_url"`
	LinkedinName  string `json:"linkedin_name"`
	LinkedinUrl   string `json:"linkedin_url"`
	WebsiteName   string `json:"website_name"`
	WebsiteUrl    string `json:"website_url"`
	UserId        string `json:"user_id"`
}

type SocialStorageI interface {
	Create(ctx context.Context, user *Social) (*Social, error)
	Update(ctx context.Context, user *Social) (*Social, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, userid string) (*Social, error)
}
