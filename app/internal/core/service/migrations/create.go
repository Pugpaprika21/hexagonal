package migrations

import (
	"context"
	"fmt"
	"os"
	"pugpaprika/app/pkg/sqlx"
)

func (m *migrationsService) CreateGoStructProcedure(ctx context.Context) error {
	var sql sqlx.Sqlx
	var pathstr = "../internal/infrastructure/migrations/mysql/"
	var files = []string{"go_gorm_create.sql", "go_gorm_get.sql", "go_request.sql", "go_respone_get.sql"}

	for _, file := range files {
		migrationPath := fmt.Sprintf("%s/%s", pathstr, file)
		if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
			return err
		}

		sqlBytes, err := os.ReadFile(migrationPath)
		if err != nil {
			return err
		}

		sql.Stmt = string(sqlBytes)

		if err := m.repos.CreateGoStructProcedure(ctx, sql); err != nil {
			return err
		}

		sql.Stmt = ""
	}

	return nil
}
