package migrations

import (
	"context"
	"fmt"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/sqlx"
)

func (m *migrationsService) DeleteAllGoStructProcedure(ctx context.Context) error {
	var sql sqlx.Sqlx

	steps := []struct {
		label string
		proc  string
	}{
		{"schema_struct", constant.GO_GORM_CREATE},
		{"request_struct", constant.GO_REQUEST},
		{"response_struct", constant.GO_RESPONE_GET},
		{"response_gorm_struct", constant.GO_GORM_GET},
		{"create_dto", constant.GO_GORM_CREATEFIELD},
		{"response_dto", constant.GO_GORM_RESPONSEFIELD},
	}

	for _, step := range steps {
		sql.Stmt = fmt.Sprintf("DROP PROCEDURE IF EXISTS %s;", step.proc)

		if err := m.repos.DeleteAllGoStructProcedure(ctx, sql); err != nil {
			return err
		}
	}

	return nil
}
