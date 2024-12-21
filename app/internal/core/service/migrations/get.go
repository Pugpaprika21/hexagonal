package migrations

import (
	"context"
	"fmt"
	"pugpaprika/app/dto/request"
	"pugpaprika/app/pkg/sqlx"
	"strings"
)

func (m *migrationsService) GetGoStructProcedure(ctx context.Context, parmObj request.GetGoStructProcedure) (string, error) {
	var sql sqlx.Sqlx

	sql.Stmt += fmt.Sprintf("set @tablename = '%s';", parmObj.Tablename)
	sql.Stmt += fmt.Sprintf("call %s(@tablename);", parmObj.StName)

	result, err := m.repos.GetGoStructProcedure(ctx, sql)
	if err != nil {
		return "", err
	}

	return "\n" + strings.Join(result, "\n"), nil
}
