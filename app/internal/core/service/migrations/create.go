package migrations

import (
	"context"
	"fmt"
	"os"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/sqlx"
)

func (m *migrationsService) CreateGoStructProcedure(ctx context.Context) error {
	var sql sqlx.Sqlx
	var pathstr = "../internal/infrastructure/migrations/mysql/"
	var files = []string{
		constant.GO_GORM_CREATE + ".sql",
		constant.GO_GORM_GET + ".sql",
		constant.GO_REQUEST + ".sql",
		constant.GO_RESPONE_GET + ".sql",
		constant.GO_GORM_CREATEFIELD + ".sql",
		constant.GO_GORM_RESPONSEFIELD + ".sql",
	}

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
