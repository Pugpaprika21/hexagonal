BEGIN WITH unique_columns AS (
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
            WHEN DATA_TYPE = 'date' THEN 'string'
            WHEN DATA_TYPE IN (
                'varchar',
                'text',
                'char',
                'longtext',
                'datetime',
                'timestamp'
            ) THEN 'string'
            WHEN DATA_TYPE = 'int' THEN 'int'
            WHEN DATA_TYPE = 'bigint' THEN 'int64'
            WHEN DATA_TYPE = 'decimal' THEN 'float64'
            ELSE DATA_TYPE
        END,
        ' `json:"',
        LOWER(COLUMN_NAME),
        '"`'
    ) AS request_struct,
    CONCAT(
        formatted_name,
        ' ',
        CASE
            DATA_TYPE
            WHEN DATA_TYPE = 'date' THEN 'string'
            WHEN DATA_TYPE IN (
                'varchar',
                'text',
                'char',
                'longtext',
                'datetime',
                'timestamp'
            ) THEN 'string'
            WHEN DATA_TYPE = 'int' THEN 'int'
            WHEN DATA_TYPE = 'bigint' THEN 'int64'
            WHEN DATA_TYPE = 'decimal' THEN 'float64'
            ELSE DATA_TYPE
        END,
        ' `json:"',
        LOWER(COLUMN_NAME),
        '"`'
    ) AS respone_struct,
    CONCAT(
        formatted_name,
        ' ',
        CASE
            DATA_TYPE
            WHEN DATA_TYPE = 'date' THEN 'string'
            WHEN DATA_TYPE IN (
                'varchar',
                'text',
                'char',
                'longtext',
                'datetime',
                'timestamp'
            ) THEN 'string'
            WHEN DATA_TYPE = 'int' THEN 'int'
            WHEN DATA_TYPE = 'bigint' THEN 'int64'
            WHEN DATA_TYPE = 'decimal' THEN 'float64'
            ELSE DATA_TYPE
        END,
        ' `gorm:"column:',
        LOWER(COLUMN_NAME),
        '"`'
    ) AS gorm_fetch_struct,
    CONCAT(
        formatted_name,
        ': res.',
        formatted_name,
        ','
    ) AS respone_result_struct
FROM
    formatted_columns
ORDER BY
    ORDINAL_POSITION;

END;