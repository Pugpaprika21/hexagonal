package router

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func (r *appRouter) test() {
	migration := r.server.Group("/migrations")

	migration.POST("/create_go_procedure", func(c echo.Context) error {
		migrationFiles := []string{
			"go_gorm_create.sql",
			"go_gorm_get.sql",
			"go_request.sql",
			"go_respone_get.sql",
		}

		pathstr := "../internal/infrastructure/migrations/mysql/"

		for _, migrationFile := range migrationFiles {
			migrationPath := fmt.Sprintf("%s/%s", pathstr, migrationFile)
			if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
				return c.String(400, fmt.Sprintf("Migration file not found: %s", migrationFile))
			}

			sqlBytes, err := os.ReadFile(migrationPath)
			if err != nil {
				return c.String(500, fmt.Sprintf("Error reading migration file: %v", err))
			}

			if err := r.dbTest.Exec(string(sqlBytes)).Error; err != nil {
				return c.String(500, fmt.Sprintf("Error running migration: %v", err))
			}
		}

		return c.String(200, "Migrations applied successfully")
	})

	migration.GET("/get_go_struct/:tablename/:stname", func(c echo.Context) error {
		setStmt := fmt.Sprintf("SET @tablename = '%s';", c.Param("tablename"))
		callStmt := fmt.Sprintf("CALL %s(@tablename);", c.Param("stname"))

		if err := r.dbTest.Exec(setStmt).Error; err != nil {
			return c.String(500, fmt.Sprintf("Error running SET SQL statement: %v", err))
		}

		rows, err := r.dbTest.Raw(callStmt).Rows()
		if err != nil {
			return c.String(500, fmt.Sprintf("Error running SQL statement: %v", err))
		}
		defer rows.Close()

		var result string

		for rows.Next() {
			var genGoStruct string
			if err := rows.Scan(&genGoStruct); err != nil {
				return c.String(500, fmt.Sprintf("Error scanning row: %v", err))
			}

			result += genGoStruct + "\n"
		}

		output := "\n" + result
		return c.String(200, output)
	})
}
