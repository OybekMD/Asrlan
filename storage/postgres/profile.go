package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type profileRepo struct {
	db *sqlx.DB
}

func NewProfile(db *sqlx.DB) repo.ProfileStorageI {
	return &profileRepo{
		db: db,
	}
}

func (s *profileRepo) GetProfile(ctx context.Context, username, year, period string) (*repo.Profile, error) {
	var profile repo.Profile

	userProfile, userId, err := s.GetUser(ctx, username)
	if err != nil {
		log.Println("Error getting GetUser in postgres method", err.Error())
		return nil, err
	}

	statisticYear, err := s.GetStatisticYear(ctx, year, userId)
	if err != nil {
		log.Println("Error getting GetStatisticYear in postgres method", err.Error())
		return nil, err
	}

	statisticWMY, err := s.GetStatisticWMY(ctx, period, userId)
	if err != nil {
		log.Println("Error getting GetStatisticWMY in postgres method", err.Error())
		return nil, err
	}

	badge, err := s.GetBadge(ctx, userId)
	if err != nil {
		log.Println("Error getting GetBadge in postgres method", err.Error())
		return nil, err
	}

	certificate, err := s.GetCertificate(ctx, userId)
	if err != nil {
		log.Println("Error getting GetCertificate in postgres method", err.Error())
		return nil, err
	}

	// Makes a sertificate
	err = s.SetLevelBadge(ctx, userId, userProfile.Score)
	if err != nil {
		log.Println("Error getting SetLevelBadge in postgres method", err.Error())
		return nil, err
	}

	profile.User = *userProfile
	profile.StatisticYear = statisticYear
	profile.StatisticWMY = statisticWMY
	profile.Badge = badge
	profile.Certificate = certificate

	return &profile, nil
}

// This function gets profile in postgres
func (s *profileRepo) GetStatisticWMY(ctx context.Context, period, userid string) ([]*repo.Statistic, error) {
	var query string
	switch period {
	case "1": // Week
	// SELECT 
    //             DATE(created_at) AS period, 
    //             SUM(score) AS score 
    //         FROM 
    //             activitys 
    //         WHERE
    //             user_id = $1
    //             AND DATE_TRUNC('week', created_at) = DATE_TRUNC('week', CURRENT_DATE)
    //         GROUP BY 
    //             DATE(created_at)
    //         ORDER BY 
    //             period
		query = `
			WITH week_dates AS (
				SELECT generate_series(
					date_trunc('week', CURRENT_DATE),
					date_trunc('week', CURRENT_DATE) + '6 days'::interval,
					'1 day'::interval
				) AS period
			)
			SELECT 
				wd.period AS period, 
				COALESCE(SUM(a.score), 0) AS score 
			FROM 
				week_dates wd
			LEFT JOIN 
				activitys a
			ON 
				DATE(a.created_at) = wd.period
				AND a.user_id = $1
			GROUP BY 
				wd.period
			ORDER BY 
				wd.period;
        `
	case "2": // Month
		query = `
            SELECT 
                DATE(created_at) AS period, 
                SUM(score) AS score 
            FROM 
                activitys 
            WHERE 
                user_id = $1
                AND DATE_TRUNC('month', created_at) = DATE_TRUNC('month', CURRENT_DATE)
            GROUP BY 
                DATE(created_at)
            ORDER BY 
                period;
        `
	case "3": // Year
		query = `
            SELECT 
                DATE(created_at) AS period, 
                SUM(score) AS score 
            FROM 
                activitys 
            WHERE 
                user_id = $1
                AND DATE_TRUNC('year', created_at) = DATE_TRUNC('year', CURRENT_DATE)
            GROUP BY 
                DATE(created_at)
            ORDER BY 
                period;
        `
	default:
		return nil, fmt.Errorf("invalid period: %s", period)
	}

	rows, err := s.db.QueryContext(ctx, query, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statistics []*repo.Statistic
	for rows.Next() {
		var stat repo.Statistic
		if err := rows.Scan(&stat.Period, &stat.Score); err != nil {
			return nil, err
		}
		statistics = append(statistics, &stat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return statistics, nil
}

func (s *profileRepo) GetStatisticYear(ctx context.Context, year, userid string) (map[string][]repo.Statistic, error) {
	// Initialize a map with all months
	monthlyStatistics := make(map[string][]repo.Statistic)
	monthsOrder := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"}
	for _, month := range monthsOrder {
		monthlyStatistics[month] = make([]repo.Statistic, 0)
	}

	query := `
        SELECT
            TO_CHAR(created_at, 'MM') AS month_number, 
            DATE(created_at) AS period, 
            SUM(score) AS score 
        FROM 
            activitys 
        WHERE 
            user_id = $1
            AND EXTRACT(YEAR FROM created_at) = $2
        GROUP BY 
            TO_CHAR(created_at, 'MM'), DATE(created_at)
        ORDER BY 
            MIN(created_at);
    `

	// query := `
	// WITH all_dates AS (
	// 	SELECT 
	// 		generate_series(
	// 			'$2-01-01'::date, 
	// 			'$2-12-31'::date, 
	// 			'1 day'::interval
	// 		)::date AS period
	// ),
	// daily_scores AS (
	// 	SELECT
	// 		DATE(created_at) AS period, 
	// 		SUM(score) AS score
	// 	FROM 
	// 		activitys 
	// 	WHERE 
	// 		user_id = $1
	// 		AND EXTRACT(YEAR FROM created_at) = $2
	// 	GROUP BY 
	// 		DATE(created_at)
	// )
	// SELECT
	// 	TO_CHAR(d.period, 'MM') AS month_number, 
	// 	d.period, 
	// 	COALESCE(ds.score, 0) AS score
	// FROM 
	// 	all_dates d
	// LEFT JOIN 
	// 	daily_scores ds ON d.period = ds.period
	// ORDER BY 
	// 	d.period;	
	// `

	rows, err := s.db.QueryContext(ctx, query, userid, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Populate the map with the retrieved data
	for rows.Next() {
		var month string
		var stat repo.Statistic
		if err := rows.Scan(&month, &stat.Period, &stat.Score); err != nil {
			return nil, err
		}
		// month = strings.TrimSpace(month) // Remove any trailing spaces from month name
		monthlyStatistics[month] = append(monthlyStatistics[month], stat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Remove empty slices from the map
	for month, stats := range monthlyStatistics {
		if len(stats) == 0 {
			delete(monthlyStatistics, month)
		}
	}

	return monthlyStatistics, nil
}

func (s *profileRepo) GetBadge(ctx context.Context, userid string) ([]*repo.Badge, error) {
	query := `
	SELECT
		b.id,
		b.name, 
		b.badge_date, 
		b.badge_type, 
		b.picture
	FROM
		badges b
	JOIN
		user_badge ub ON ub.badge_id = b.id
	WHERE
		ub.user_id = $1
	`

	rows, err := s.db.QueryContext(ctx, query, userid)
	if err != nil {
		log.Println("Error selecting profile badges with in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseBadges []*repo.Badge
	for rows.Next() {
		var badge repo.Badge
		err = rows.Scan(
			&badge.Id,
			&badge.Name,
			&badge.BadgeDate,
			&badge.BadgeType,
			&badge.Picture,
		)
		if err != nil {
			log.Println("Error scanning badge in getall badge method of postgres", err.Error())
			return nil, err
		}

		responseBadges = append(responseBadges, &badge)
	}

	return responseBadges, nil
}

func (s *profileRepo) SetLevelBadge(ctx context.Context, userID string, score int64) error {
	// Determine the level based on score
	var badgeId int
	switch {
	case score >= 5000000:
		badgeId = 6
	case score >= 2000000:
		badgeId = 7
	case score >= 1000000:
		badgeId = 8
	case score >= 500000:
		badgeId = 9
	case score >= 250000:
		badgeId = 10
	case score >= 100000:
		badgeId = 11
	case score >= 50000:
		badgeId = 12
	case score >= 10000:
		badgeId = 13
	case score >= 1000:
		badgeId = 14
	case score >= 100:
		badgeId = 15
	default:
		return nil // No badge for scores less than 1000
	}

	queryExist := `
		SELECT
			badge_id
		FROM 
			user_badge ub
		WHERE
			ub.user_id = $1 
			AND badge_id = $2
		`

	var isExists int

	row := s.db.QueryRowContext(ctx, queryExist, userID, badgeId)
	err := row.Scan(&isExists)
	if err != sql.ErrNoRows {
		return err
	}
	if isExists != 0 {
		return nil // User already has this badge, no need to reassign
	}

	// // Get the badge ID based on the image link
	// queryBadgeID := `
	// 	SELECT
	// 		id
	// 	FROM
	// 		badges
	// 	WHERE
	// 		picture = $1`

	// var badgeID int

	// row = s.db.QueryRowContext(ctx, queryBadgeID)
	// err = row.Scan(&badgeID)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("Badge not found:", err)
	// 		return err
	// 	}
	// 	log.Println("Error fetching badge ID:", err)
	// 	return err
	// }

	// var badge_id string
	// if nullbadge_id.Valid {
	// 	badge_id = nullbadge_id.String
	// }

	// Insert into user_badge
	insertQuery := `
		INSERT INTO user_badge (user_id, badge_id)
		VALUES ($1, $2)
	`

	_, err = s.db.ExecContext(ctx, insertQuery, userID, badgeId)
	if err != nil {
		log.Println("Error inserting user badge:", err)
		return err
	}

	return nil
}

func (s *profileRepo) GetUser(ctx context.Context, username string) (*repo.ProfileUser, string, error) {
	query := `
	SELECT
		id,
		name, 
		username,
		bio, 
		birth_day, 
		avatar, 
		coint,
		created_at
	FROM
		users
	WHERE
		username = $1
	AND
		deleted_at IS NULL
	`

	var responseUser repo.ProfileUser
	var user_id string
	var nullBio, nullBirthDay, nullAvatar sql.NullString
	err := s.db.QueryRowContext(ctx, query, username).Scan(
		&user_id,
		&responseUser.Name,
		&responseUser.Username,
		&nullBio,
		&nullBirthDay,
		&nullAvatar,
		&responseUser.Coint,
		&responseUser.CreatedAt,
	)
	if err != nil {
		log.Println("Eror getting user in postgres method", err.Error())
		return nil, "", err
	}

	if nullBio.Valid {
		responseUser.Bio = nullBio.String
	}
	if nullBirthDay.Valid {
		responseUser.BirthDay = nullBirthDay.String
	}
	if nullAvatar.Valid {
		responseUser.Avatar = nullAvatar.String
	}

	queryScore := `
	SELECT 
		SUM(score) AS total_score
	FROM 
		activitys 
	WHERE 
		user_id = $1
	`
	err = s.db.QueryRowContext(ctx, queryScore, user_id).Scan(
		&responseUser.Score,
	)

	if err != nil {
		log.Println("Eror getting user score in postgres method", err.Error())
		return nil, "", err
	}

	queryStreak := `
	WITH activity_dates AS (
		SELECT DISTINCT DATE(created_at) AS activity_date
		FROM activitys
		WHERE user_id = $1
	),
	ranked_dates AS (
		SELECT 
			activity_date,
			RANK() OVER (ORDER BY activity_date) AS rank_date
		FROM activity_dates
	),
	streaks AS (
		SELECT 
			activity_date,
			rank_date,
			activity_date - INTERVAL '1 day' * (rank_date - 1) AS streak_group
		FROM ranked_dates
	),
	current_streak AS (
		SELECT 
			COUNT(*) AS streak_length
		FROM streaks
		WHERE streak_group = (
			SELECT streak_group
			FROM streaks
			ORDER BY activity_date DESC
			LIMIT 1
		)
	)
	SELECT streak_length
	FROM current_streak;
	`
	err = s.db.QueryRowContext(ctx, queryStreak, user_id).Scan(
		&responseUser.Streak,
	)
	if err != nil {
		log.Println("Eror getting user streak in postgres method", err.Error())
		return nil, "", err
	}

	queryRank := `
	WITH ranked_scores AS (
		SELECT
			a.user_id,
			RANK() OVER (ORDER BY SUM(a.score) DESC) AS rank
		FROM
			activitys a
		GROUP BY
			a.user_id
	)
	SELECT
		rs.rank
	FROM
		ranked_scores rs
	WHERE
		rs.user_id = $1
	`
	err = s.db.QueryRowContext(ctx, queryRank, user_id).Scan(
		&responseUser.Rank,
	)
	if err != nil {
		log.Println("Eror getting user streak in postgres method", err.Error())
		return nil, "", err
	}

	return &responseUser, user_id, nil
}

func (s *profileRepo) GetCertificate(ctx context.Context, user_id string) ([]*repo.Certificate, error) {
	query := `
	SELECT
		name,
		pdfile
	FROM
		certificates
	WHERE
		user_id = $1
	`

	rows, err := s.db.QueryContext(ctx, query, user_id)
	if err != nil {
		fmt.Println("Error1:", err)
		log.Println("Eror getting certificate in postgres method", err.Error())
		return nil, err
	}
	defer rows.Close()
	var responseCertificate []*repo.Certificate
	for rows.Next() {
		var certificate repo.Certificate
		err = rows.Scan(
			&certificate.Name,
			&certificate.Pdfile,
		)
		if err != nil {
			fmt.Println("Error2:", err)
			log.Println("Error scanning certificate method of postgres", err.Error())
			return nil, err
		}

		responseCertificate = append(responseCertificate, &certificate)
	}

	return responseCertificate, nil
}
