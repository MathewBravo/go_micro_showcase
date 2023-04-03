package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const TIMEOUT = time.Second * 5

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		Book: Book{},
	}
}

type Models struct {
	Book Book
}

type Book struct {
	ID          int64          `json:"ID"`
	Title       string         `json:"title"`
	Authors     pq.StringArray `json:"authors"`
	Description string         `json:"despcription"`
	IsReading   bool           `json:"isReading"`
	HasRead     bool           `json:"hasRead"`
	Thumbnail   string         `json:"thumbnail"`
}

func (b *Book) GetAll() ([]*Book, error) {
	ctx, timedout := context.WithTimeout(context.Background(), TIMEOUT)
	defer timedout()

	query := `select id, title, authors, description, isreading, hasread, thumbnail from books`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*Book

	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Authors,
			&book.Description,
			&book.IsReading,
			&book.HasRead,
			&book.Thumbnail,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}

func (b *Book) GetBookByID(id string) (Book, error) {
	ctx, timedout := context.WithTimeout(context.Background(), TIMEOUT)
	defer timedout()

	var book Book

	query := `select id, title, authors, description, isreading, hasread, thumbnail from books where id=$1`
	err := db.QueryRowContext(ctx, query, id).Scan(&book.Title,
		&book.Authors,
		&book.Description,
		&book.IsReading,
		&book.HasRead,
		&book.Thumbnail,
	)
	if err != nil {
		return Book{}, err
	}

	return book, nil

}

func (b *Book) Update() error {
	ctx, timedout := context.WithTimeout(context.Background(), TIMEOUT)
	defer timedout()

	statement := `update books set 
				title = $1,
				authors = $2,
				description = $3
				isReading = $4
				hasRead = $5
				thumbnail = $6
	`

	_, err := db.ExecContext(ctx, statement,
		b.Title,
		b.Authors,
		b.Description,
		b.IsReading,
		b.HasRead,
		b.Thumbnail,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) Insert(book Book) (int, error) {
	ctx, timedout := context.WithTimeout(context.Background(), TIMEOUT)
	defer timedout()

	var idNewRow int

	statement := `insert into books (title, authors, description, isreading, hasRead, thumbnail)
				values($1,$2,$3,$4,$5)
				return id
	`

	err := db.QueryRowContext(ctx, statement,
		book.Title,
		book.Authors,
		book.Description,
		book.IsReading,
		book.HasRead,
		book.Thumbnail,
	).Scan(&idNewRow)
	if err != nil {
		return 0, err
	}

	return idNewRow, nil
}
