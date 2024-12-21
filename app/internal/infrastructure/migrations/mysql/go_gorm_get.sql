CREATE PROCEDURE go_gorm_get(IN tableName VARCHAR(255)) BEGIN WITH unique_columns AS (
    SELECT
        COLUMN_NAME,
        ORDINAL_POSITION,
        DATA_TYPE,
        ROW_NUMBER() OVER (
            PARTITION BY COLUMN_NAME
            ORDER BY
                ORDINAL_POSITION
        ) AS row_num
    FROM
        INFORMATION_SCHEMA.COLUMNS
    WHERE
        TABLE_NAME = tableName
),
formatted_columns AS (
    SELECT
        COLUMN_NAME,
        ORDINAL_POSITION,
        DATA_TYPE,
        CASE
            WHEN COLUMN_NAME = 'id' THEN 'ID'
            WHEN COLUMN_NAME LIKE '%_id' THEN CONCAT(
                UPPER(LEFT(SUBSTRING_INDEX(COLUMN_NAME, '_', 1), 1)),
                LOWER(
                    SUBSTRING(SUBSTRING_INDEX(COLUMN_NAME, '_', 1), 2)
                ),
                CASE
                    WHEN LENGTH(COLUMN_NAME) - LENGTH(REPLACE(COLUMN_NAME, '_', '')) > 1 THEN CONCAT(
                        UPPER(
                            LEFT(
                                SUBSTRING_INDEX(SUBSTRING_INDEX(COLUMN_NAME, '_', 2), '_', -1),
                                1
                            )
                        ),
                        LOWER(
                            SUBSTRING(
                                SUBSTRING_INDEX(SUBSTRING_INDEX(COLUMN_NAME, '_', 2), '_', -1),
                                2
                            )
                        )
                    )
                    ELSE ''
                END,
                'ID'
            )
            ELSE CONCAT_WS(
                '',
                UPPER(LEFT(SUBSTRING_INDEX(COLUMN_NAME, '_', 1), 1)),
                LOWER(
                    SUBSTRING(SUBSTRING_INDEX(COLUMN_NAME, '_', 1), 2)
                ),
                CASE
                    WHEN LOCATE('_', COLUMN_NAME) > 0 THEN CONCAT(
                        UPPER(
                            LEFT(
                                SUBSTRING_INDEX(SUBSTRING_INDEX(COLUMN_NAME, '_', -1), '_', 1),
                                1
                            )
                        ),
                        LOWER(
                            SUBSTRING(
                                SUBSTRING_INDEX(SUBSTRING_INDEX(COLUMN_NAME, '_', -1), '_', 1),
                                2
                            )
                        )
                    )
                    ELSE ''
                END
            )
        END AS formatted_name
    FROM
        unique_columns
    WHERE
        row_num = 1
)
SELECT
    CONCAT(
        formatted_name,
        ' ',
        CASE
            DATA_TYPE
            WHEN 'date' THEN 'sql.NullString'
            WHEN 'varchar' THEN 'sql.NullString'
            WHEN 'text' THEN 'sql.NullString'
            WHEN 'char' THEN 'sql.NullString'
            WHEN 'longtext' THEN 'sql.NullString'
            WHEN 'datetime' THEN 'sql.NullString'
            WHEN 'timestamp' THEN 'sql.NullString'
            WHEN 'int' THEN 'sql.NullInt32'
            WHEN 'bigint' THEN 'sql.NullInt64'
            WHEN 'tinyint' THEN 'sql.NullBool'
            WHEN 'decimal' THEN 'sql.NullFloat64'
            ELSE DATA_TYPE
        END,
        ' `gorm:"column:',
        LOWER(COLUMN_NAME),
        '"`'
    ) AS go_gorm_get
FROM
    formatted_columns
ORDER BY
    ORDINAL_POSITION;

END;