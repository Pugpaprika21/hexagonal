package sqlx

type Sqlx struct {
	Stmt        string
	Args        []interface{}
	WhereClause []string
	Each        Each // fetch rows with attr {Each} .Take(&sql.Each.Rows)
}

type Each struct {
	Rows []map[string]interface{}
}

func Nil[T comparable](value *T) *T {
	if value == nil {
		return nil
	}

	var zero T
	if *value == zero {
		return nil
	}
	return value
}
