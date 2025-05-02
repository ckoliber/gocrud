package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	_    struct{} `db:"users" json:"-"`
	ID   *int     `db:"id" json:"id"`
	Name string   `db:"name" json:"name"`
	Age  int      `db:"age" json:"age"`
}

func UnitTests(ctx context.Context, t *testing.T, repo Repository[User]) {
	t.Run("Post", func(t *testing.T) {
		users := []User{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 35},
			{Name: "Charlie", Age: 45},
		}

		result, err := repo.Post(ctx, &users)
		assert.NoError(t, err)
		assert.Len(t, result, 3)

		for i, user := range result {
			assert.NotNil(t, user.ID)
			assert.Equal(t, users[i].Name, user.Name)
			assert.Equal(t, users[i].Age, user.Age)
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		where := map[string]any{"id": map[string]any{"_eq": "1"}}
		result, err := repo.Get(ctx, &where, nil, nil, nil)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
	})

	t.Run("GetWithFilters", func(t *testing.T) {
		where := map[string]any{"age": map[string]any{"_gt": "25"}}
		result, err := repo.Get(ctx, &where, nil, nil, nil)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("GetPagination", func(t *testing.T) {
		limit := 5
		skip := 0
		result, err := repo.Get(ctx, nil, nil, &limit, &skip)
		assert.NoError(t, err)
		assert.LessOrEqual(t, len(result), limit)
	})

	t.Run("Put", func(t *testing.T) {
		users := []User{
			{ID: &[]int{2}[0], Name: "Alice Updated", Age: 26},
			{ID: &[]int{3}[0], Name: "Bob Updated", Age: 36},
			{ID: &[]int{4}[0], Name: "Charlie Updated", Age: 46},
		}

		result, err := repo.Put(ctx, &users)
		assert.NoError(t, err)
		assert.Len(t, result, 3)

		for i, user := range result {
			assert.Equal(t, users[i].ID, user.ID)
			assert.Equal(t, users[i].Name, user.Name)
			assert.Equal(t, users[i].Age, user.Age)
		}
	})

	t.Run("DeleteByID", func(t *testing.T) {
		where := map[string]any{"id": map[string]any{"_eq": "1"}}
		result, err := repo.Delete(ctx, &where)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
	})

	t.Run("DeleteWithFilters", func(t *testing.T) {
		where := map[string]any{"age": map[string]any{"_lt": "30"}}
		result, err := repo.Delete(ctx, &where)
		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}
