package repo

import "context"

type UserBadge struct {
	UserId  string `json:"user_id"`
	BadgeId string `json:"badge_id"`
}

type UserBadgeStorageI interface {
	Create(ctx context.Context, badge *UserBadge) (bool, error)
	Delete(ctx context.Context, user_id, badge_id string) (bool, error)
	AllUsersBadgeByUserId(ctx context.Context, page, limit uint64) ([]*Badge, int64, error)
	YearlyUsersBadgesByYear(ctx context.Context, year int) ([]*Badge, error)
}
