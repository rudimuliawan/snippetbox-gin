package models

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Snippet struct {
	ID      int
	Title   *string
	Content *string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *gorm.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	snippet := Snippet{
		Title:   &title,
		Content: &content,
		Created: time.Now(),
		Expires: time.Now().AddDate(0, 0, expires),
	}

	ctx := context.Background()
	result := gorm.WithResult()
	err := gorm.G[Snippet](m.DB, result).Create(ctx, &snippet)
	if err != nil {
		return -1, err
	}

	return snippet.ID, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	ctx := context.Background()

	snippet, err := gorm.G[Snippet](m.DB).Where("id = ?", id).First(ctx)
	if err != nil {
		return Snippet{}, err
	}

	return snippet, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	ctx := context.Background()

	snippets, err := gorm.G[Snippet](m.DB).
		Order("ID DESC").
		Limit(10).
		Find(ctx)

	if err != nil {
		return nil, err
	}

	return snippets, nil
}
