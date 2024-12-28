CREATE PROCEDURE go_gorm_createfield(IN tableName VARCHAR(255)) BEGIN WITH unique_columns AS (
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
            ': ',
            'sqlx.Nil(rec.',
            formatted_name,
            '),'
        ) AS go_gorm_createfield
FROM
    formatted_columns
ORDER BY
    ORDINAL_POSITION;
END;