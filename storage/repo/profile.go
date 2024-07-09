package repo

import "context"

type ProfileUser struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	BirthDay  string `json:"birth_day"`
	Avatar    string `json:"avatar"`
	Coint     int64  `json:"coint"`
	Score     int64  `json:"score"`
	Rank      int64  `json:"rank"`
	Streak    int64  `json:"streak"`
	CreatedAt string `json:"created_at"`
}

// Statistic start
type Statistic struct {
	Score  int64  `json:"score"`
	Period string `json:"period"`
}

type StatisticMonth struct {
	Month []Statistic `json:"month"`
}

// Statistic end

type Profile struct {
	User          ProfileUser            `json:"user"`
	StatisticYear map[string][]Statistic `json:"statistic"`
	StatisticWMY  []*Statistic           `json:"statisticwmy"`
	Badge         []*Badge               `json:"badge"`
	Certificate   []*Certificate         `json:"certificate"`
}

type Certificate struct {
	Name   string `json:"name"`
	Pdfile string `json:"pdfile"`
}

type ProfileStorageI interface {
	GetProfile(ctx context.Context, username, year, period string) (*Profile, error)
	GetUser(ctx context.Context, username string) (*ProfileUser, string, error)
	GetStatisticYear(ctx context.Context, year, userid string) (map[string][]Statistic, error)
	GetStatisticWMY(ctx context.Context, period, userid string) ([]*Statistic, error)
	GetBadge(ctx context.Context, userid string) ([]*Badge, error)
	SetLevelBadge(ctx context.Context, userID string, score int64) error
	GetCertificate(ctx context.Context, user_id string) ([]*Certificate, error)
}
