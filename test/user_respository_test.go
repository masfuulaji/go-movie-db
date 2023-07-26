package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/masfuulaji/go-movie-db/internal/models"
	"github.com/masfuulaji/go-movie-db/internal/repositories"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(
		postgres.Config{
            DSN: "sqlmock_db_0",
            DriverName: "postgres",
			Conn: db,
            PreferSimpleProtocol: true,
		},
	), &gorm.Config{
    })
	require.NoError(t, err)

	return mock, gormDB
}

// i give up sad
func TestCreateUser(t *testing.T) {
	mock, db := setupMockDB(t)

	repo := repositories.NewUserRepository(db)
	user := models.User{
		Name:     "test",
		Email:    "test",
		Password: "test",
	}

	// Expect the BEGIN transaction call
	mock.ExpectBegin()

	// Expect the INSERT INTO users query with RETURNING "id" and ignore other fields
    sql := `INSERT INTO "users" ("created_at", "updated_at", "deleted_at", "name", "email", "password") VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id"`
	mock.ExpectExec(sql).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Name, user.Email, user.Password).
        WillReturnResult(sqlmock.NewResult(1, 1))

	// Expect the COMMIT transaction call
	mock.ExpectCommit()

	err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

