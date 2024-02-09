package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "video.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ts, err := TrendingCount(db)

	if err != nil {
		panic(err)
	}
	for _, v := range ts {
		fmt.Printf("%s ::: %d", v.T.String(), v.Count)
		fmt.Println()
	}

}

type Trend struct {
	T     time.Time
	Count int
}

// func (t Trend) Value(driver.Value, error){

// }

func TrendingCount(db *sql.DB) ([]Trend, error) {
	limit := 30
	trends := make([]Trend, 0, limit)
	rows, err := db.QueryContext(context.Background(), "select v.trending_date, count(v.trending_date) from videos v group by v.trending_date ORDER by v.trending_date desc limit ?", limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t Trend
		var strTime string
		err = rows.Scan(&strTime, &t.Count)
		if err != nil {
			return nil, err
		}
		t.T, err = time.Parse("06.02.01", strTime)
		if err != nil {
			return nil, err
		}
		trends = append(trends, t)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return trends, nil

}
