package gocrud

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

var xdb *sql.DB

type User struct {
	_    struct{} `db:"users" json:"-"`
	ID   *int     `db:"id" json:"id" required:"false"`
	Name string   `db:"name" json:"name" required:"false" maxLength:"30" example:"David" doc:"User name"`
	Age  int      `db:"age" json:"age" required:"false" minimum:"1" maximum:"120" example:"25" doc:"User age from 1 to 120"`
}

func TestRegister(t *testing.T) {
	// Create a new in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:?cache=shared")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	xdb = db

	// Create the users table
	_, err = db.Exec("CREATE TABLE users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER)")
	if err != nil {
		panic(err)
	}

	// Create a new Huma API
	_, api := humatest.New(t)
	repo := NewSQLRepository[User](xdb)
	Register(api, repo, &Config[User]{})

	t.Run("POST single", func(t *testing.T) {
		// Create a new user
		user := User{Name: "David", Age: 25}
		resp := api.Post("/user/one", &user)
		assert.Equal(t, resp.Code, 200)

		// Check the response
		var result User
		assert.Empty(t, json.Unmarshal(resp.Body.Bytes(), &result))
		assert.NotEmpty(t, result.ID)
		assert.Equal(t, result.Name, user.Name)
		assert.Equal(t, result.Age, user.Age)
	})

	t.Run("POST bulk", func(t *testing.T) {
		users := []User{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 35},
			{Name: "Charlie", Age: 45},
		}
		resp := api.Post("/user", &users)
		assert.Equal(t, resp.Code, 200)

		var result []User
		assert.Empty(t, json.Unmarshal(resp.Body.Bytes(), &result))
		for i := range users {
			assert.NotEmpty(t, result[i].ID)
			assert.Equal(t, result[i].Name, users[i].Name)
			assert.Equal(t, result[i].Age, users[i].Age)
		}
	})
}
