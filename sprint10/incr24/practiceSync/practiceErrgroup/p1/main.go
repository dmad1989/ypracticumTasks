package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type PostID string

type Post struct {
	ID      PostID
	Content string
}

type PostProvider interface {
	GetPost(ctx context.Context, postID PostID) (*Post, error)
}
type batchPostProvider struct {
	provider PostProvider
}

// строка ниже не несёт функциональной нагрузки
// её можно убрать без последствий для работы программы
// это отладочная строка
// в этой строке приведением типов проверяем,
// реализует ли *batchPostProvider интерфейс BatchPostProvider —
// если нет, если методы прописаны неверно,
// то компилятор выдаст на этой строке ошибку типизации
// var _ BatchPostProvider = (*batchPostProvider)(nil)

func (p *batchPostProvider) GetPosts(ctx context.Context,
	postIDs []PostID) (map[PostID]*Post, error) {
	grp, ctx := errgroup.WithContext(ctx)

	posts := make([]*Post, len(postIDs))
	for i, id := range postIDs {
		// нельзя использовать переменные цикла i и id в горутине,
		// так как, вероятнее всего, горутина получит только их конечные значения,
		// поэтому определяем переменные внутри цикла
		index, idx := i, id

		grp.Go(func() error {
			post, err := p.provider.GetPost(ctx, idx)
			if err != nil {
				return err
			}

			posts[index] = post
			return nil
		})
	}

	if err := grp.Wait(); err != nil {
		return nil, err
	}

	return p.collectPostsByIDs(postIDs, posts), nil
}

func (p *batchPostProvider) collectPostsByIDs(ids []PostID,
	posts []*Post) map[PostID]*Post {
	result := make(map[PostID]*Post, len(ids))
	for i, id := range ids {
		result[id] = posts[i]
	}

	return result
}
