package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type User struct {
	_    struct{} `db:"users" json:"-"`
	ID   *int     `db:"id" json:"id"`
	Name string   `db:"name" json:"name"`
	Age  int      `db:"age" json:"age"`
}

// names returns the user names in result order
func names(users []User) []string {
	result := []string{}
	for _, user := range users {
		result = append(result, user.Name)
	}
	return result
}

// createUsersTable creates the users table with the given DDL, dropping any leftover from a previous run
func createUsersTable(t *testing.T, db *sql.DB, ddl string) {
	_, _ = db.Exec("DROP TABLE users")
	_, err := db.Exec(ddl)
	require.NoError(t, err)
}

func UnitTests(ctx context.Context, t *testing.T, repo Repository[User]) {
	// IDs by name, filled from the Post result so tests never assume ID values
	ids := map[string]string{}

	get := func(t *testing.T, where map[string]any, order map[string]any, limit int, skip int) []User {
		result, err := repo.Get(ctx, &where, &order, &limit, &skip)
		require.NoError(t, err)
		return result
	}

	t.Run("Post", func(t *testing.T) {
		users := []User{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 35},
			{Name: "Charlie", Age: 45},
			{Name: "Dave", Age: 55},
		}

		result, err := repo.Post(ctx, &users)
		require.NoError(t, err)
		require.Len(t, result, 4)

		for i, user := range result {
			require.NotNil(t, user.ID)
			assert.Equal(t, users[i].Name, user.Name)
			assert.Equal(t, users[i].Age, user.Age)
			ids[user.Name] = fmt.Sprint(*user.ID)
		}
		assert.Len(t, ids, 4, "IDs must be unique")
	})

	t.Run("PostSingle", func(t *testing.T) {
		users := []User{{Name: "Eve", Age: 65}}

		result, err := repo.Post(ctx, &users)
		require.NoError(t, err)
		require.Len(t, result, 1)
		require.NotNil(t, result[0].ID)
		assert.Equal(t, "Eve", result[0].Name)
		ids["Eve"] = fmt.Sprint(*result[0].ID)
	})

	t.Run("GetAll", func(t *testing.T) {
		assert.ElementsMatch(t, []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}, names(get(t, nil, nil, 0, 0)))
	})

	t.Run("GetNoMatch", func(t *testing.T) {
		assert.Empty(t, get(t, map[string]any{"age": map[string]any{"_gt": "1000"}}, nil, 0, 0))
	})

	t.Run("Operators", func(t *testing.T) {
		cases := []struct {
			name  string
			where map[string]any
			want  []string
		}{
			{"_eq", map[string]any{"id": map[string]any{"_eq": ids["Bob"]}}, []string{"Bob"}},
			{"_neq", map[string]any{"age": map[string]any{"_neq": "25"}}, []string{"Bob", "Charlie", "Dave", "Eve"}},
			{"_gt", map[string]any{"age": map[string]any{"_gt": "35"}}, []string{"Charlie", "Dave", "Eve"}},
			{"_gte", map[string]any{"age": map[string]any{"_gte": "35"}}, []string{"Bob", "Charlie", "Dave", "Eve"}},
			{"_lt", map[string]any{"age": map[string]any{"_lt": "35"}}, []string{"Alice"}},
			{"_lte", map[string]any{"age": map[string]any{"_lte": "35"}}, []string{"Alice", "Bob"}},
			{"_like", map[string]any{"name": map[string]any{"_like": "%li%"}}, []string{"Alice", "Charlie"}},
			{"_nlike", map[string]any{"name": map[string]any{"_nlike": "%li%"}}, []string{"Bob", "Dave", "Eve"}},
			{"_ilike", map[string]any{"name": map[string]any{"_ilike": "a%"}}, []string{"Alice"}},
			{"_nilike", map[string]any{"name": map[string]any{"_nilike": "a%"}}, []string{"Bob", "Charlie", "Dave", "Eve"}},
			{"_in", map[string]any{"id": map[string]any{"_in": []string{ids["Alice"], ids["Dave"]}}}, []string{"Alice", "Dave"}},
			{"_nin", map[string]any{"id": map[string]any{"_nin": []string{ids["Alice"], ids["Dave"]}}}, []string{"Bob", "Charlie", "Eve"}},
			{"multiple fields", map[string]any{"age": map[string]any{"_gt": "30"}, "name": map[string]any{"_like": "C%"}}, []string{"Charlie"}},
			{"_and", map[string]any{"_and": []any{
				map[string]any{"age": map[string]any{"_gt": "30"}},
				map[string]any{"age": map[string]any{"_lt": "50"}},
			}}, []string{"Bob", "Charlie"}},
			{"_or", map[string]any{"_or": []any{
				map[string]any{"age": map[string]any{"_lt": "30"}},
				map[string]any{"age": map[string]any{"_gt": "60"}},
			}}, []string{"Alice", "Eve"}},
			{"_not", map[string]any{"_not": map[string]any{"age": map[string]any{"_eq": "25"}}}, []string{"Bob", "Charlie", "Dave", "Eve"}},
			{"nested", map[string]any{"_or": []any{
				map[string]any{"_and": []any{
					map[string]any{"age": map[string]any{"_gt": "30"}},
					map[string]any{"age": map[string]any{"_lt": "40"}},
				}},
				map[string]any{"_not": map[string]any{"age": map[string]any{"_lt": "60"}}},
			}}, []string{"Bob", "Eve"}},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.ElementsMatch(t, c.want, names(get(t, c.where, nil, 0, 0)))
			})
		}
	})

	t.Run("Order", func(t *testing.T) {
		assert.Equal(t, []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}, names(get(t, nil, map[string]any{"age": "ASC"}, 0, 0)))
		assert.Equal(t, []string{"Eve", "Dave", "Charlie", "Bob", "Alice"}, names(get(t, nil, map[string]any{"name": "DESC"}, 0, 0)))
	})

	t.Run("Pagination", func(t *testing.T) {
		order := map[string]any{"age": "ASC"}
		assert.Equal(t, []string{"Alice", "Bob"}, names(get(t, nil, order, 2, 0)), "limit")
		assert.Equal(t, []string{"Bob", "Charlie"}, names(get(t, nil, order, 2, 1)), "limit and skip")
		assert.Equal(t, []string{"Dave", "Eve"}, names(get(t, nil, order, 0, 3)), "skip")
		assert.Equal(t, []string{"Charlie"}, names(get(t, map[string]any{"age": map[string]any{"_gt": "30"}}, order, 1, 1)), "where, limit and skip")
		assert.Len(t, get(t, nil, nil, 2, 1), 2, "limit and skip without order")
	})

	t.Run("Put", func(t *testing.T) {
		users := []User{{Name: "Bob Updated", Age: 36}, {Name: "Charlie Updated", Age: 46}}
		for i, name := range []string{"Bob", "Charlie"} {
			id := 0
			fmt.Sscan(ids[name], &id)
			users[i].ID = &id
		}

		result, err := repo.Put(ctx, &users)
		require.NoError(t, err)
		require.Len(t, result, 2)
		for i, user := range result {
			assert.Equal(t, *users[i].ID, *user.ID)
			assert.Equal(t, users[i].Name, user.Name)
			assert.Equal(t, users[i].Age, user.Age)
		}

		// The update is persisted
		stored := get(t, map[string]any{"id": map[string]any{"_eq": ids["Bob"]}}, nil, 0, 0)
		require.Len(t, stored, 1)
		assert.Equal(t, "Bob Updated", stored[0].Name)
		assert.Equal(t, 36, stored[0].Age)
	})

	t.Run("PutNotFound", func(t *testing.T) {
		id := -1
		users := []User{{ID: &id, Name: "Nobody", Age: 1}}

		result, err := repo.Put(ctx, &users)
		require.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("DeleteByID", func(t *testing.T) {
		where := map[string]any{"id": map[string]any{"_eq": ids["Alice"]}}
		result, err := repo.Delete(ctx, &where)
		require.NoError(t, err)
		assert.Equal(t, []string{"Alice"}, names(result), "returns the deleted rows")
		assert.Empty(t, get(t, where, nil, 0, 0), "row is gone")
	})

	t.Run("DeleteWithFilters", func(t *testing.T) {
		where := map[string]any{"age": map[string]any{"_gt": "50"}}
		result, err := repo.Delete(ctx, &where)
		require.NoError(t, err)
		assert.ElementsMatch(t, []string{"Dave", "Eve"}, names(result))
		assert.ElementsMatch(t, []string{"Bob Updated", "Charlie Updated"}, names(get(t, nil, nil, 0, 0)))
	})

	t.Run("DeleteNoMatch", func(t *testing.T) {
		where := map[string]any{"age": map[string]any{"_gt": "1000"}}
		result, err := repo.Delete(ctx, &where)
		require.NoError(t, err)
		assert.Empty(t, result)
	})
}
