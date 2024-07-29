package stars

import (
	"time"

	"github.com/rept0id/birthday-planetary-backend/apis/main/model"
)

func FindByISODate(dbs model.DBs, isoDate string) (res map[string]interface{}, err error) {
	var res_data []model.Star
	var res_metadata = make(map[string]interface{}) // Initialize res_metadata

	var star model.Star
	var ly float64

	// Define the query with PostgreSQL placeholder
	q := "SELECT id, name, ly FROM public.stars WHERE ly >= $1 ORDER BY ly ASC LIMIT 1"

	// Parse the ISO date string
	tISODate, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return nil, err
	}

	// Calculate the number of light years since the given date
	ly = time.Since(tISODate).Hours() / 24.0 / 365.0

	// Execute the query and scan the result into the star struct
	err = dbs.DBStarsConn.QueryRow(q, ly).Scan(&star.ID, &star.Name, &star.LY)
	if err != nil {
		return nil, err
	}

	// Append the star to the results data slice
	res_data = append(res_data, star)

	// Prepare the final result map
	res = map[string]interface{}{
		"data":     res_data,
		"metadata": res_metadata,
	}

	return res, nil
}
