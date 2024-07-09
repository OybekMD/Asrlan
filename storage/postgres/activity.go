package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp/v3"
)

type activityRepo struct {
	db *sqlx.DB
}

func NewActivity(db *sqlx.DB) repo.ActivityStorageI {
	return &activityRepo{
		db: db,
	}
}

// This function create activity in postgres
func (s *activityRepo) Create(ctx context.Context, activity *repo.Activity) (*repo.Activity, error) {
	query := `
	INSERT INTO activitys(
		day,
		score,
		lesson_id,
		user_id
	)
	VALUES ($1, $2, $3, $4) 
	RETURNING 
		id`

	err := s.db.QueryRowContext(
		ctx,
		query,
		activity.Day,
		activity.Score,
		activity.LessonId,
		activity.UserId).Scan(&activity.Id)
	if err != nil {
		log.Println("Eror creating activity in postgres method", err.Error())
		return nil, err
	}

	return activity, nil
}

// // This function get all activity with page and limit posgtres
// func (s *activityRepo) GetAll(ctx context.Context, page, limit uint64) ([]*repo.Activity, int64, error) {
// 	query := ``

// 	offset := limit * (page - 1)
// 	rows, err := s.db.QueryContext(ctx, query, limit, offset)
// 	if err != nil {
// 		log.Println("Error selecting activitys with page and limit in postgres", err.Error())
// 		return nil, 0, err
// 	}
// 	defer rows.Close()

// 	var responseActivitys []*repo.Activity
// 	for rows.Next() {
// 		// var activity repo.Activity
// 		// err = rows.Scan(
// 		// 	&activity.CreatedAt,
// 		// 	&activity.UpdatedAt,
// 		// )
// 		// if err != nil {
// 		// 	log.Println("Error scanning activity in getall activity method of postgres", err.Error())
// 		// 	return nil, 0, err
// 		// }

// 		responseActivitys = append(responseActivitys, &activity)
// 	}

// 	count := len(responseActivitys)

// 	return responseActivitys, int64(count), nil
// }

func (s *activityRepo) GetAllGroupedByMonth(ctx context.Context) (map[string][]*repo.Activity, error) {
	query := `
    SELECT 
        TO_CHAR(day, 'YYYY-MM') AS month,
        json_agg(json_build_object(
            'id', id,
            'day', day,
            'score', score,
            'user_id', user_id
        )) AS activities
    FROM 
        activitys
    WHERE
        EXTRACT(YEAR FROM day) = 2024
        AND user_id = '678e9012-e89b-12d3-a456-426614174006'
    GROUP BY 
        month
    ORDER BY 
        month DESC
    `
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string][]*repo.Activity)

	for rows.Next() {
		var month string
		var activityJSON []byte
		var activities []*repo.Activity

		if err := rows.Scan(&month, &activityJSON); err != nil {
			pp.Println("error 1: ", err)
			return nil, err
		}

		if err := json.Unmarshal(activityJSON, &activities); err != nil {
			pp.Println("error 2: ", err)
			return nil, err
		}

		result[month] = activities
	}

	if err := rows.Err(); err != nil {
		pp.Println("error 3: ", err)
		return nil, err
	}

	return result, nil
}

func (s *activityRepo) GetAllGroupedByChoice(ctx context.Context, choise string) (map[string][]*repo.Activity, error) {
	var query string
	if choise == "year" {
		query = `
        SELECT 
        EXTRACT(YEAR FROM day) AS year,
        json_agg(json_build_object(
            'id', id,
            'day', day,
            'user_id', user_id
        )) AS activities
    FROM 
        activitys
    GROUP BY 
        year
    ORDER BY 
        year DESC
    `
	} else if choise == "month" {
		query = `
        SELECT 
        TO_CHAR(day, 'YYYY-MM') AS month,
        json_agg(json_build_object(
            'id', id,
            'day', day,
            'user_id', user_id
        )) AS activities
    FROM 
        activitys
    GROUP BY 
        month
    ORDER BY 
        month DESC
    `
	} else {
		query = `
        WITH current_week_data AS (
            SELECT 
                EXTRACT(YEAR FROM day) AS year,
                EXTRACT(WEEK FROM day) AS week,
                day,
                user_id
            FROM 
                activitys
            WHERE 
                EXTRACT(YEAR FROM day) = EXTRACT(YEAR FROM CURRENT_DATE) AND
                EXTRACT(WEEK FROM day) = EXTRACT(WEEK FROM CURRENT_DATE)
        )
        SELECT 
            EXTRACT(YEAR FROM CURRENT_DATE) AS year,
            EXTRACT(WEEK FROM CURRENT_DATE) AS week_number,
            json_agg(json_build_object(
                'day', day,
                'user_id', user_id
            )) AS activities
        FROM 
            current_week_data
        GROUP BY 
            year, week_number        
        `
	}

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string][]*repo.Activity)

	for rows.Next() {
		var month string
		var activityJSON []byte
		var activities []*repo.Activity

		if err := rows.Scan(&month, &activityJSON); err != nil {
			pp.Println("error 1: ", err)
			return nil, err
		}

		if err := json.Unmarshal(activityJSON, &activities); err != nil {
			pp.Println("error 2: ", err)
			return nil, err
		}

		result[month] = activities
	}

	if err := rows.Err(); err != nil {
		pp.Println("error 3: ", err)
		return nil, err
	}

	return result, nil
}
