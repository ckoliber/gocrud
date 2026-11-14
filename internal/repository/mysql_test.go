package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestMySQLRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping container test in short mode")
	}
	ctx := context.Background()

	container, err := mysql.Run(ctx, "mysql:8.4")
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)

	dsn, err := container.ConnectionString(ctx)
	require.NoError(t, err)
	db, err := sql.Open("mysql", dsn)
	require.NoError(t, err)
	defer db.Close()

	createUsersTable(t, db, `CREATE TABLE users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT)`)

	UnitTests(ctx, t, NewMySQLRepository[User](db))
}
