package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

type Video struct {
	Id          string    // video_id
	Title       string    // title
	PublishTime time.Time // publish_time
	Tags        []string  // tags
	Views       int       // views
}

func readVideoCSV(csvFile string) ([]Video, error) {
	// открываем csv файл
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var videos []Video

	// определим индексы нужных полей
	const (
		Id          = 0 // video_id
		Title       = 2 // title
		PublishTime = 5 // publish_time
		Tags        = 6 // tags
		Views       = 7 // views
	)

	// конструируем Reader из пакета encoding/csv
	// он умеет читать строки csv-файла
	r := csv.NewReader(file)
	// пропустим первую строку с именами полей
	if _, err := r.Read(); err != nil {
		return nil, err
	}

	for {
		// csv.Reader за одну операцию Read() считывает одну csv-запись
		l, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// инициализируем целевую структуру,
		// в которую будем делать разбор csv-записи
		v := Video{
			Id:    l[Id],
			Title: l[Title],
		}
		// парсинг строковых записей в типизированные поля структуры
		if v.PublishTime, err = time.Parse(time.RFC3339, l[PublishTime]); err != nil {
			return nil, err
		}
		tags := strings.Split(l[Tags], "|")
		for i, v := range tags {
			tags[i] = strings.Trim(v, `"`)
		}
		v.Tags = tags
		if v.Views, err = strconv.Atoi(l[Views]); err != nil {
			return nil, err
		}
		// добавляем полученную структуру в слайс
		videos = append(videos, v)
	}
	return videos, nil
}

func insertVideos(ctx context.Context, db *sql.DB, videos []Video) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, v := range videos {
		// в этом случае возвращаемое значение не несёт полезной информации,
		// поэтому его можно игнорировать
		_, err := tx.ExecContext(ctx,
			"INSERT INTO videos (video_id, title, publish_time, tags, views)"+
				" VALUES(?,?,?,?,?)", v.Id, v.Title, v.PublishTime,
			strings.Join(v.Tags, `|`), v.Views)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func main() {
	// открываем соединение с БД
	db, err := sql.Open("sqlite", "newvideo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// читаем записи из файла в слайс []Video вспомогательной функцией
	videos, err := readVideoCSV("USvideos.csv")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	_, err = db.ExecContext(ctx, `CREATE TABLE videos (
		"video_id"  TEXT,
		"title" TEXT,
		"publish_time" TEXT,
		"tags" TEXT,
		"views" INTEGER NOT NULL DEFAULT 0
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// записываем []Video в базу данных
	// тоже вспомогательной функцией
	err = insertVideos(ctx, db, videos)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Всего csv-записей %v\n", len(videos))
}
