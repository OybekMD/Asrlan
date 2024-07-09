package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

const socialsTableName = "socials"

type socialRepo struct {
	db *sqlx.DB
}

func NewSocial(db *sqlx.DB) repo.SocialStorageI {
	return &socialRepo{
		db: db,
	}
}

// Old
// func QueryBuildCreateSocial(reqData *repo.Social) (string, []interface{}) {
// 	args := []interface{}{}
// 	query := "INSERT INTO social("
// 	reply := " RETURNING "
// 	count := 1
// 	if reqData.LocationName != "" {
// 		query += ", location_name"
// 		reply += ", location_name"
// 		args = append(args, reqData.LocationName)
// 		count++
// 	}
// 	if reqData.LocationUrl != "" {
// 		query += ", location_ur;"
// 		reply += ", location_ur;"
// 		args = append(args, reqData.LocationUrl)
// 		count++
// 	}
// 	if reqData.EducationName != "" {
// 		query += ", education_name"
// 		reply += ", education_name"
// 		args = append(args, reqData.EducationName)
// 		count++
// 	}
// 	if reqData.EducationUrl != "" {
// 		query += ", education_url"
// 		reply += ", education_url"
// 		args = append(args, reqData.EducationUrl)
// 		count++
// 	}
// 	if reqData.TelegramName != "" {
// 		query += ", telegram_name"
// 		reply += ", telegram_name"
// 		args = append(args, reqData.TelegramName)
// 		count++
// 	}
// 	if reqData.TelegramUrl != "" {
// 		query += ", telegram_url"
// 		reply += ", telegram_url"
// 		args = append(args, reqData.TelegramUrl)
// 		count++
// 	}
// 	if reqData.TwitterName != "" {
// 		query += ", twitter_name"
// 		reply += ", twitter_name"
// 		args = append(args, reqData.TwitterName)
// 		count++
// 	}
// 	if reqData.TwitterUrl != "" {
// 		query += ", twitter_url"
// 		reply += ", twitter_url"
// 		args = append(args, reqData.TwitterUrl)
// 		count++
// 	}
// 	if reqData.InstagramName != "" {
// 		query += ", instagram_name"
// 		reply += ", instagram_name"
// 		args = append(args, reqData.InstagramName)
// 		count++
// 	}
// 	if reqData.InstagramUrl != "" {
// 		query += ", instagram_url"
// 		reply += ", instagram_url"
// 		args = append(args, reqData.InstagramUrl)
// 		count++
// 	}
// 	if reqData.YoutubeName != "" {
// 		query += ", youtube_name"
// 		reply += ", youtube_name"
// 		args = append(args, reqData.YoutubeName)
// 		count++
// 	}
// 	if reqData.YoutubeUrl != "" {
// 		query += ", youtube_url"
// 		reply += ", youtube_url"
// 		args = append(args, reqData.YoutubeUrl)
// 		count++
// 	}
// 	if reqData.LinkedinName != "" {
// 		query += ", linkedin_name"
// 		reply += ", linkedin_name"
// 		args = append(args, reqData.LinkedinName)
// 		count++
// 	}
// 	if reqData.LinkedinUrl != "" {
// 		query += ", linkedin_url"
// 		reply += ", linkedin_url"
// 		args = append(args, reqData.LinkedinUrl)
// 		count++
// 	}
// 	if reqData.WebsiteName != "" {
// 		query += ", website_name"
// 		reply += ", website_name"
// 		args = append(args, reqData.WebsiteName)
// 		count++
// 	}
// 	if reqData.WebsiteUrl != "" {
// 		query += ", website_url"
// 		reply += ", website_url"
// 		args = append(args, reqData.WebsiteUrl)
// 		count++
// 	}
// 	query += ", user_id"
// 	reply += ", user_id"
// 	args = append(args, reqData.UserId)
// 	count++
// 	query += ") VALUES ("
// 	for i := 1; i <= count; i++ {
// 		num := strconv.Itoa(i)
// 		if i == count {
// 			query = query + "$" + num + ")"
// 		} else {
// 			query = query + "$" + num + ", "
// 		}
// 	}
// 	query += reply
// 	return query, args
// }

// QueryBuildCreateSocial dynamically constructs an INSERT query based on the provided Social data
func QueryBuildCreateSocial(reqData *repo.Social) (string, []interface{}) {
	args := []interface{}{}
	columns := []string{}
	placeholders := []string{}

	addColumn := func(name string, value interface{}) {
		if value != "" {
			columns = append(columns, name)
			args = append(args, value)
			placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))
		}
	}

	addColumn("location_name", reqData.LocationName)
	addColumn("location_url", reqData.LocationUrl)
	addColumn("education_name", reqData.EducationName)
	addColumn("education_url", reqData.EducationUrl)
	addColumn("telegram_name", reqData.TelegramName)
	addColumn("telegram_url", reqData.TelegramUrl)
	addColumn("twitter_name", reqData.TwitterName)
	addColumn("twitter_url", reqData.TwitterUrl)
	addColumn("instagram_name", reqData.InstagramName)
	addColumn("instagram_url", reqData.InstagramUrl)
	addColumn("youtube_name", reqData.YoutubeName)
	addColumn("youtube_url", reqData.YoutubeUrl)
	addColumn("linkedin_name", reqData.LinkedinName)
	addColumn("linkedin_url", reqData.LinkedinUrl)
	addColumn("website_name", reqData.WebsiteName)
	addColumn("website_url", reqData.WebsiteUrl)

	// Always include the user_id column
	columns = append(columns, "user_id")
	args = append(args, reqData.UserId)
	placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))

	query := fmt.Sprintf("INSERT INTO socials (%s) VALUES (%s) RETURNING user_id",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args
}

// QueryBuildUpdateSocial dynamically constructs an UPDATE query based on the provided Social data
func QueryBuildUpdateSocial(reqData *repo.Social) (string, []interface{}) {
	args := []interface{}{}
	updates := []string{}

	addUpdate := func(columnName string, value interface{}) {
		if value != "" {
			updates = append(updates, fmt.Sprintf("%s = $%d", columnName, len(args)+1))
			args = append(args, value)
		}
	}

	addUpdate("location_name", reqData.LocationName)
	addUpdate("location_url", reqData.LocationUrl)
	addUpdate("education_name", reqData.EducationName)
	addUpdate("education_url", reqData.EducationUrl)
	addUpdate("telegram_name", reqData.TelegramName)
	addUpdate("telegram_url", reqData.TelegramUrl)
	addUpdate("twitter_name", reqData.TwitterName)
	addUpdate("twitter_url", reqData.TwitterUrl)
	addUpdate("instagram_name", reqData.InstagramName)
	addUpdate("instagram_url", reqData.InstagramUrl)
	addUpdate("youtube_name", reqData.YoutubeName)
	addUpdate("youtube_url", reqData.YoutubeUrl)
	addUpdate("linkedin_name", reqData.LinkedinName)
	addUpdate("linkedin_url", reqData.LinkedinUrl)
	addUpdate("website_name", reqData.WebsiteName)
	addUpdate("website_url", reqData.WebsiteUrl)

	// Always include the user_id for updating the correct record
	args = append(args, reqData.UserId)

	query := fmt.Sprintf("UPDATE socials SET %s WHERE user_id = $%d RETURNING user_id",
		strings.Join(updates, ", "),
		len(args),
	)

	return query, args
}

// This function create social in postgres
func (s *socialRepo) Create(ctx context.Context, social *repo.Social) (*repo.Social, error) {
	query, args := QueryBuildCreateSocial(social)

	// var respSocial repo.Social
	err := s.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&social.UserId,
	)
	if err != nil {
		log.Println("Eror creating user in postgres method", err.Error())
		return nil, err
	}

	return social, nil
}

// This function update social info from postgres
func (s *socialRepo) Update(ctx context.Context, newSocial *repo.Social) (*repo.Social, error) {
	query, args := QueryBuildUpdateSocial(newSocial)

	err := s.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&newSocial.UserId,
	)
	if err != nil {
		log.Println("Eror updating social in postgres method", err.Error())
		return nil, err
	}
	return newSocial, nil
}

// This function delete social info from postgres
func (s *socialRepo) Delete(ctx context.Context, userid string) (bool, error) {
	query := `
	DELETE
		socials
	WHERE
		user_id = $1
	AND
		deleted_at IS NULL
	`

	resutl, err := s.db.ExecContext(ctx, query, userid)
	if err != nil {
		log.Println("Error deleting social", err.Error())
		return false, err
	}

	rowEffect, err := resutl.RowsAffected()
	if err != nil {
		log.Println("Error getting row effect", err.Error())
		return false, err
	}

	if rowEffect == 0 {
		log.Println("Nothing updated, Social")
		return false, nil
	}

	return true, nil
}

// This function gets social in postgres
func (s *socialRepo) Get(ctx context.Context, id string) (*repo.Social, error) {
	query := `
	SELECT
		location_name,
		location_url,
		education_name,
		education_url,
		telegram_name,
		telegram_url,
		twitter_name,
		twitter_url,
		instagram_name,
		instagram_url,
		youtube_name,
		youtube_url,
		linkedin_name,
		linkedin_url,
		website_name,
		website_url,
		user_id
	FROM
		socials
	WHERE
		user_id = $1
	`

	var responseSocial repo.Social
	var (
		nullLocationName, nullLocationUrl,
		nullEducationName, nullEducationUrl,
		nullTelegramName, nullTelegramUrl,
		nullTwitterName, nullTwitterUrl,
		nullInstagramName, nullInstagramUrl,
		nullYoutubeName, nullYoutubeUrl,
		nullLinkedinName, nullLinkedinUrl,
		nullWebsiteName, nullWebsiteUrl sql.NullString
	)

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&nullLocationName, &nullLocationUrl,
		&nullEducationName, &nullEducationUrl,
		&nullTelegramName, &nullTelegramUrl,
		&nullTwitterName, &nullTwitterUrl,
		&nullInstagramName, &nullInstagramUrl,
		&nullYoutubeName, &nullYoutubeUrl,
		&nullLinkedinName, &nullLinkedinUrl,
		&nullWebsiteName, &nullWebsiteUrl,
		&responseSocial.UserId,
	)
	if err != nil {
		log.Println("Eror getting social in postgres method", err.Error())
		return nil, err
	}
	if nullLocationName.Valid {
		responseSocial.LocationName = nullLocationName.String
	}
	if nullLocationUrl.Valid {
		responseSocial.LocationUrl = nullLocationUrl.String
	}
	if nullEducationName.Valid {
		responseSocial.EducationName = nullEducationName.String
	}
	if nullEducationUrl.Valid {
		responseSocial.EducationUrl = nullEducationUrl.String
	}
	if nullTelegramName.Valid {
		responseSocial.TelegramName = nullTelegramName.String
	}
	if nullTelegramUrl.Valid {
		responseSocial.TelegramUrl = nullTelegramUrl.String
	}
	if nullTwitterName.Valid {
		responseSocial.TwitterName = nullTwitterName.String
	}
	if nullTwitterUrl.Valid {
		responseSocial.TwitterUrl = nullTwitterUrl.String
	}
	if nullInstagramName.Valid {
		responseSocial.InstagramName = nullInstagramName.String
	}
	if nullInstagramUrl.Valid {
		responseSocial.InstagramUrl = nullInstagramUrl.String
	}
	if nullYoutubeName.Valid {
		responseSocial.YoutubeName = nullYoutubeName.String
	}
	if nullYoutubeUrl.Valid {
		responseSocial.YoutubeUrl = nullYoutubeUrl.String
	}
	if nullLinkedinName.Valid {
		responseSocial.LinkedinName = nullLinkedinName.String
	}
	if nullLinkedinUrl.Valid {
		responseSocial.LinkedinUrl = nullLinkedinUrl.String
	}
	if nullWebsiteName.Valid {
		responseSocial.WebsiteName = nullWebsiteName.String
	}
	if nullWebsiteUrl.Valid {
		responseSocial.WebsiteUrl = nullWebsiteUrl.String
	}

	return &responseSocial, nil
}
