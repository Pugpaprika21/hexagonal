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

// GetAllGoStructProcedure *
func (m *migrationsService) GetAllGoStructProcedure(ctx context.Context, parmObj request.GetAllGoStructProcedure) ([]string, error) {
	var sql sqlx.Sqlx
	var procedures = []string{
		constant.GO_GORM_CREATE,
		constant.GO_GORM_GET,
		constant.GO_REQUEST,
		constant.GO_RESPONE_GET,
		constant.GO_GORM_CREATEFIELD,
		constant.GO_GORM_RESPONSEFIELD,
	}

	var goStructResult []string

	for _, stName := range procedures {
		sql.Stmt = fmt.Sprintf("set @tablename = '%s';", parmObj.Tablename)
		_, err := m.repos.GetGoStructProcedure(ctx, sql)
		if err != nil {
			return nil, err
		}

		sql.Stmt = fmt.Sprintf("call %s(@tablename);", stName)

		result, err := m.repos.GetGoStructProcedure(ctx, sql)
		if err != nil {
			return nil, err
		}

		structResult := strings.Join(result, "\n")

		goStructResult = append(goStructResult, structResult)
	}

	return goStructResult, nil
}
