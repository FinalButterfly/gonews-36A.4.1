package postgres

import (
	"GoNews/pkg/storage"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	pass := "somepass"
	store, err := New("postgresql://postgres:" + pass + "@localhost/gonews")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		connectionString string
	}
	tests := []struct {
		name    string
		args    args
		want    *Store
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				"postgresql://postgres:" + pass + "@localhost/gonews",
			},
			want:    store,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.connectionString)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_News(t *testing.T) {
	pass := "somepass"
	store, err := New("postgresql://postgres:" + pass + "@localhost/gonews")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		limit int
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    []storage.Post
		wantErr bool
	}{
		{
			name: "default",
			s:    store,
			args: args{
				limit: 1,
			},
			want: []storage.Post{
				{
					ID:      2,
					Title:   "Анализ теста по Go с PHDays",
					Content: "<p>Думаю, многие из вас сталкивались с замысловатыми задачками, которые в реальной практике встретить почти невозможно, но которые очень любят давать во всяких тестах и на собеседованиях.</p><p>В конце мая прошла конференция PHDays, на которой был тест как раз с такими задачками. К моему сожалению, я провалила этот тест, но затем разобралась что, как и почему, и хочу поделиться с вами.</p><p>Итак, 5 картинок с кодом, к каждому дается 4 варианта ответа. </p> <a href=\"https://habr.com/ru/post/571002/?utm_campaign=571002&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut\">Читать далее</a>",
					PubTime: 0,
					Link:    "https://habr.com/ru/post/571002/?utm_campaign=571002&utm_source=habrahabr&utm_medium=rss",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.News(tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.News() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.News() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_AddPosts(t *testing.T) {
	pass := "somepass"
	store, err := New("postgresql://postgres:" + pass + "@localhost/gonews")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		posts []storage.Post
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s:    store,
			args: args{
				posts: []storage.Post{
					{
						Title:   "Анализ теста по Go с PHDays",
						Content: "<p>Думаю, многие из вас сталкивались с замысловатыми задачками, которые в реальной практике встретить почти невозможно, но которые очень любят давать во всяких тестах и на собеседованиях.</p><p>В конце мая прошла конференция PHDays, на которой был тест как раз с такими задачками. К моему сожалению, я провалила этот тест, но затем разобралась что, как и почему, и хочу поделиться с вами.</p><p>Итак, 5 картинок с кодом, к каждому дается 4 варианта ответа. </p> <a href=\"https://habr.com/ru/post/571002/?utm_campaign=571002&amp;utm_source=habrahabr&amp;utm_medium=rss#habracut\">Читать далее</a>",
						PubTime: 0,
						Link:    "https://habr.com/ru/post/571002/?utm_campaign=571002&utm_source=habrahabr&utm_medium=rss",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.AddPosts(tt.args.posts); (err != nil) != tt.wantErr {
				t.Errorf("Store.AddPosts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
