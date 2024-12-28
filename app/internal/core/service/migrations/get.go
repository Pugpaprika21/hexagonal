package migrations

import (
	"context"
	"fmt"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/constant"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (m *migrationsService) GetGoStructProcedure(ctx context.Context, parmObj request.GetGoStructProcedure) (string, error) {
	var sql sqlx.Sqlx

	sql.Stmt = fmt.Sprintf("set @tablename = '%s';", parmObj.Tablename)
	_, err := m.repos.GetGoStructProcedure(ctx, sql)
	if err != nil {
		return "", err
	}
	sql.Stmt = fmt.Sprintf("call %s(@tablename);", parmObj.StName)

	result, err := m.repos.GetGoStructProcedure(ctx, sql)
	if err != nil {
		return "", err
	}

	return "\n" + strings.Join(result, "\n"), nil
}

func (m *migrationsService) GetAllGoStructProcedure(ctx context.Context, parmObj request.GetAllGoStructProcedure) (string, error) {
	var sql sqlx.Sqlx
	var formatStrResult strings.Builder

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
		sql.Stmt = fmt.Sprintf("set @tablename = '%s';", parmObj.Tablename)
		if _, err := m.repos.GetAllGoStructProcedure(ctx, sql); err != nil {
			return "", err
		}

		sql.Stmt = fmt.Sprintf("call %s(@tablename);", step.proc)
		results, err := m.repos.GetAllGoStructProcedure(ctx, sql)
		if err != nil {
			return "", err
		}

		formatStrResult.WriteString(fmt.Sprintf("\n\n%s\n\n%s", step.label, strings.Join(results, "\n")))
	}

	return formatStrResult.String(), nil
}
