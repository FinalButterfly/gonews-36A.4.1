package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// The store which has the connection
type Store struct {
	db *pgxpool.Pool
}

// Constructor for Store
func New(connectionString string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

// Gets news where n = number of entries
func (s *Store) News(limit int) ([]storage.Post, error) {
	if limit == 0 {
		limit = 10
	}
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			title,
			content,
			pubtime,
			link
		FROM posts
		ORDER BY id
		LIMIT $1;
	`,
		limit,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.PubTime,
			&t.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, t)

	}
	return posts, rows.Err()
}

// Adds multiple posts to the database
func (s *Store) AddPosts(posts []storage.Post) error {
	var err error
	tx, err := s.db.Begin(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range posts {
		_, err = tx.Exec(context.Background(), `
		INSERT into posts(
			title,
			content,
			pubtime,
			link
		)
		values(
			$1,
			$2,
			$3,
			$4
		) RETURNING id;
		`,
			post.Title,
			post.Content,
			post.PubTime,
			post.Link,
		)

		if err == nil {
			tx.Commit(context.Background())
		}
	}
	return err
}
