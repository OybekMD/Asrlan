package storage

import (
	"asrlan-monolight/storage/postgres"
	"asrlan-monolight/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Login() repo.LoginStorageI
	User() repo.UserStorageI
	Admin() repo.AdminStorageI
	Social() repo.SocialStorageI
	Badge() repo.BadgeStorageI
	UserBadge() repo.UserBadgeStorageI
	Activity() repo.ActivityStorageI
	// Couser
	Language() repo.LanguageStorageI
	Level() repo.LevelStorageI
	Topic() repo.TopicStorageI
	Lesson() repo.LessonStorageI
	Content() repo.ContentStorageI
	ContentFile() repo.ContentFileStorageI
	Dashboard() repo.DashboardStorageI
	UserLanguage() repo.UserLanguageStorageI
	UserLevel() repo.UserLevelStorageI
	UserTopic() repo.UserTopicStorageI
	UserLesson() repo.UserLessonStorageI
	Book() repo.BookStorageI
	Profile() repo.ProfileStorageI
}

type storagePg struct {
	loginRepo        repo.LoginStorageI
	userRepo         repo.UserStorageI
	adminRepo        repo.AdminStorageI
	socialRepo       repo.SocialStorageI
	badgeRepo        repo.BadgeStorageI
	userBadgeRepo    repo.UserBadgeStorageI
	activityRepo     repo.ActivityStorageI
	languageRepo     repo.LanguageStorageI
	levelRepo        repo.LevelStorageI
	topicRepo        repo.TopicStorageI
	lessonRepo       repo.LessonStorageI
	contentRepo      repo.ContentStorageI
	contentFileRepo  repo.ContentFileStorageI
	dashboardRepo    repo.DashboardStorageI
	userLanguageRepo repo.UserLanguageStorageI
	userLevelRepo    repo.UserLevelStorageI
	userTopicRepo    repo.UserTopicStorageI
	userLessonRepo   repo.UserLessonStorageI
	bookRepo         repo.BookStorageI
	profileRepo      repo.ProfileStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		loginRepo:        postgres.NewLogin(db),
		userRepo:         postgres.NewUser(db),
		adminRepo:        postgres.NewAdmin(db),
		socialRepo:       postgres.NewSocial(db),
		badgeRepo:        postgres.NewBadge(db),
		userBadgeRepo:    postgres.NewUserBadge(db),
		activityRepo:     postgres.NewActivity(db),
		languageRepo:     postgres.NewLanguage(db),
		levelRepo:        postgres.NewLevel(db),
		topicRepo:        postgres.NewTopic(db),
		lessonRepo:       postgres.NewLesson(db),
		contentRepo:      postgres.NewContent(db),
		contentFileRepo:  postgres.NewContentFile(db),
		dashboardRepo:    postgres.NewDashboard(db),
		userLanguageRepo: postgres.NewUserLanguage(db),
		userLevelRepo:    postgres.NewUserLevel(db),
		userTopicRepo:    postgres.NewUserTopic(db),
		userLessonRepo:   postgres.NewUserLesson(db),
		bookRepo:         postgres.NewBook(db),
		profileRepo:      postgres.NewProfile(db),
	}
}

func (s *storagePg) Login() repo.LoginStorageI {
	return s.loginRepo
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Admin() repo.AdminStorageI {
	return s.adminRepo
}

func (s *storagePg) Social() repo.SocialStorageI {
	return s.socialRepo
}

func (s *storagePg) Badge() repo.BadgeStorageI {
	return s.badgeRepo
}

func (s *storagePg) UserBadge() repo.UserBadgeStorageI {
	return s.userBadgeRepo
}

func (s *storagePg) Activity() repo.ActivityStorageI {
	return s.activityRepo
}

func (s *storagePg) Language() repo.LanguageStorageI {
	return s.languageRepo
}

func (s *storagePg) Level() repo.LevelStorageI {
	return s.levelRepo
}

func (s *storagePg) Topic() repo.TopicStorageI {
	return s.topicRepo
}

func (s *storagePg) Lesson() repo.LessonStorageI {
	return s.lessonRepo
}

func (s *storagePg) Content() repo.ContentStorageI {
	return s.contentRepo
}

func (s *storagePg) ContentFile() repo.ContentFileStorageI {
	return s.contentFileRepo
}

func (s *storagePg) Dashboard() repo.DashboardStorageI {
	return s.dashboardRepo
}

func (s *storagePg) UserLanguage() repo.UserLanguageStorageI {
	return s.userLanguageRepo
}

func (s *storagePg) UserLevel() repo.UserLevelStorageI {
	return s.userLevelRepo
}

func (s *storagePg) UserTopic() repo.UserTopicStorageI {
	return s.userTopicRepo
}

func (s *storagePg) UserLesson() repo.UserLessonStorageI {
	return s.userLessonRepo
}

func (s *storagePg) Book() repo.BookStorageI {
	return s.bookRepo
}

func (s *storagePg) Profile() repo.ProfileStorageI {
	return s.profileRepo
}
