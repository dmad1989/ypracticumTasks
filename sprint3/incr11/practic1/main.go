package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "video.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	row := db.QueryRowContext(context.Background(),
		"SELECT title, likes, comments_disabled "+
			"FROM videos ORDER BY likes DESC LIMIT 1")
	var (
		title  string
		likes  int
		comdis bool
	)
	// порядок переменных должен соответствовать порядку колонок в запросе
	err = row.Scan(&title, &likes, &comdis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s | %d | %t \r\n", title, likes, comdis)
	res, err := getDesc(context.Background(), db, "0EbFotkXOiA")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// videos, err := QueryVideos(context.Background(), db, limit)
	// if err != nil {
	// 	panic(err)
	// }

	// for i, v := range videos {
	// 	println(i, " ", v.Id, " ", v.Title, " ", v.Views)
	// }

}

func getDesc(ctx context.Context, db *sql.DB, id string) (string, error) {
	row := db.QueryRowContext(ctx,
		"SELECT description FROM videos WHERE video_id = ?", id)
	var desc sql.NullString

	err := row.Scan(&desc)
	if err != nil {
		return "", err
	}
	if desc.Valid {
		return desc.String, nil
	}

	list, err := QueryTagVideos(context.Background(), db, 5)
	if err != nil {
		panic(err)
	}
	// для теста проверим, какие строки содержит v.Tags
	// выведем по 4 первых тега
	for _, v := range list {
		length := 4
		if len(v.Tags) < length {
			length = len(v.Tags)
		}
		fmt.Println(strings.Join(v.Tags[:length], " # "))
	}

	return "-----", nil
}

type Tags []string

type Video struct {
	Id    string
	Title string
	Tags  Tags
}

// Value — функция реализующая интерфейс driver.Valuer
func (tags Tags) Value() (driver.Value, error) {
	// преобразуем []string в string
	if len(tags) == 0 {
		return "", nil
	}
	return strings.Join(tags, "|"), nil
}

func (tags *Tags) Scan(value interface{}) error {
	// если `value` равен `nil`, будет возвращён пустой массив
	if value == nil {
		*tags = Tags{}
		return nil
	}

	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return fmt.Errorf("cannot scan value. %w", err)
	}

	v, ok := sv.(string)
	if !ok {
		return errors.New("cannot scan value. cannot convert value to string")
	}
	*tags = strings.Split(v, "|")

	// удаляем кавычки у тегов
	for i, v := range *tags {
		(*tags)[i] = strings.Trim(v, `"`)
	}
	return nil
}

func QueryTagVideos(ctx context.Context, db *sql.DB, limit int) ([]Video, error) {
	videos := make([]Video, 0, limit)

	rows, err := db.QueryContext(ctx, "SELECT video_id, title, tags from videos "+
		"GROUP BY video_id ORDER BY views LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var v Video
		// все теги должны автоматически преобразоваться в слайс v.Tags
		err = rows.Scan(&v.Id, &v.Title, &v.Tags)
		if err != nil {
			return nil, err
		}
		videos = append(videos, v)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// Video — структура видео.
// type Video struct {
// 	Id    string
// 	Title string
// 	Views int64
// }

// // limit — максимальное количество записей.
// const limit = 20

// func QueryVideos(ctx context.Context, db *sql.DB, limit int) ([]Video, error) {
// 	videos := make([]Video, 0, limit)

// 	rows, err := db.QueryContext(ctx, "SELECT video_id, title, views from videos ORDER BY views LIMIT ?", limit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// обязательно закрываем перед возвратом функции
// 	defer rows.Close()

// 	// пробегаем по всем записям
// 	for rows.Next() {
// 		var v Video
// 		err = rows.Scan(&v.Id, &v.Title, &v.Views)
// 		if err != nil {
// 			return nil, err
// 		}

// 		videos = append(videos, v)
// 	}

// 	// проверяем на ошибки
// 	err = rows.Err()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return videos, nil
// }
