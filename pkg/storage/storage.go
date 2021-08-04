package storage

// Публикация, получаемая из RSS.
type Post struct {
	ID      int    // номер записи
	Title   string // заголовок публикации
	Content string // содержание публикации
	PubTime int64  // время публикации
	Link    string // ссылка на источник
}

// Database methods
type Interface interface {
	News(limit int) ([]Post, error) // получение всех публикаций
	AddPosts(posts []Post) error    // Adds multiple posts to the database
}
