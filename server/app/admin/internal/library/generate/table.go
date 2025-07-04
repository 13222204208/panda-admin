package generate

import (
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// TableMeta 表元数据信息结构体
type TableMeta struct {
	Name       string `json:"name"`        // 表名
	Comment    string `json:"comment"`     // 表注释/描述
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

// TableMetaResult 表元数据查询结果
type TableMetaResult struct {
	List  []TableMeta `json:"list"`  // 表列表
	Total int         `json:"total"` // 总数量
}

// ColumnMeta 字段元数据信息结构体
type ColumnMeta struct {
	ColumnName      string `json:"column_name"`      // 字段名
	DataType        string `json:"data_type"`        // 数据类型
	ColumnType      string `json:"column_type"`      // 完整字段类型
	IsNullable      string `json:"is_nullable"`      // 是否可为空
	ColumnDefault   string `json:"column_default"`   // 默认值
	ColumnComment   string `json:"column_comment"`   // 字段注释
	ColumnKey       string `json:"column_key"`       // 键类型(PRI, UNI, MUL)
	Extra           string `json:"extra"`            // 额外信息(auto_increment等)
	OrdinalPosition int    `json:"ordinal_position"` // 字段位置
}

// GetTablesMeta 获取表元数据信息(支持分页)
// dbType: mysql, pgsql, sqlite, mssql
// page: 当前页码(从1开始)
// pageSize: 每页条数
func GetTablesMeta(dbType string, page, pageSize int) (TableMetaResult, error) {
	ctx := gctx.New()
	db := g.DB()
	result := TableMetaResult{
		List:  make([]TableMeta, 0),
		Total: 0,
	}

	// 获取当前数据库名称
	currentDb, err := db.GetValue(ctx, "SELECT DATABASE()")
	if err != nil {
		fmt.Println("获取当前数据库名称出错:", err)
		return result, err
	}

	// 计算总数量
	countQuery := `
        SELECT 
            COUNT(*) as count
        FROM
            information_schema.TABLES
        WHERE
            TABLE_SCHEMA = ?
    `
	countResult, err := db.GetValue(ctx, countQuery, currentDb.String())
	if err != nil {
		fmt.Println("查询总数出错:", err)
		return result, err
	}
	result.Total = countResult.Int()

	// 计算分页参数
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	// 查询当前数据库的表信息（带分页）
	query := `
        SELECT
            TABLE_NAME AS table_name,
            TABLE_COMMENT AS table_comment,
            CREATE_TIME AS create_time,
            UPDATE_TIME AS update_time
        FROM
            information_schema.TABLES
        WHERE
            TABLE_SCHEMA = ?
        ORDER BY
            CREATE_TIME DESC, TABLE_NAME ASC
        LIMIT ?, ?
    `
	rows, err := db.GetAll(ctx, query, currentDb.String(), offset, pageSize)
	if err != nil {
		fmt.Println("查询出错:", err)
		return result, err
	}

	// 处理查询结果
	for _, row := range rows {
		tableMeta := TableMeta{
			Name:       row["table_name"].String(),
			Comment:    row["table_comment"].String(),
			CreateTime: row["create_time"].Time().String(),
			UpdateTime: row["update_time"].Time().String(),
		}
		result.List = append(result.List, tableMeta)
	}

	return result, nil
}

// GetTableColumns 获取指定表的字段信息
func GetTableColumns(tableName string) ([]ColumnMeta, error) {
	ctx := gctx.New()
	db := g.DB()
	columns := make([]ColumnMeta, 0)

	// 获取当前数据库名称
	currentDb, err := db.GetValue(ctx, "SELECT DATABASE()")
	if err != nil {
		return nil, fmt.Errorf("获取当前数据库名称出错: %v", err)
	}

	// 查询表字段信息
	query := `
        SELECT
            COLUMN_NAME as column_name,
            DATA_TYPE as data_type,
            COLUMN_TYPE as column_type,
            IS_NULLABLE as is_nullable,
            IFNULL(COLUMN_DEFAULT, '') as column_default,
            IFNULL(COLUMN_COMMENT, '') as column_comment,
            IFNULL(COLUMN_KEY, '') as column_key,
            IFNULL(EXTRA, '') as extra,
            ORDINAL_POSITION as ordinal_position
        FROM
            information_schema.COLUMNS
        WHERE
            TABLE_SCHEMA = ? AND TABLE_NAME = ?
        ORDER BY
            ORDINAL_POSITION ASC
    `

	rows, err := db.GetAll(ctx, query, currentDb.String(), tableName)
	if err != nil {
		return nil, fmt.Errorf("查询表 %s 字段信息出错: %v", tableName, err)
	}

	// 处理查询结果
	for _, row := range rows {
		columnMeta := ColumnMeta{
			ColumnName:      row["column_name"].String(),
			DataType:        row["data_type"].String(),
			ColumnType:      row["column_type"].String(),
			IsNullable:      row["is_nullable"].String(),
			ColumnDefault:   row["column_default"].String(),
			ColumnComment:   row["column_comment"].String(),
			ColumnKey:       row["column_key"].String(),
			Extra:           row["extra"].String(),
			OrdinalPosition: row["ordinal_position"].Int(),
		}
		columns = append(columns, columnMeta)
	}

	return columns, nil
}

// GetTableColumnsWithDetail 获取指定表的详细字段信息（包含更多元数据）
func GetTableColumnsWithDetail(tableName string) ([]ColumnMeta, error) {
	ctx := gctx.New()
	db := g.DB()
	columns := make([]ColumnMeta, 0)

	// 获取当前数据库名称
	currentDb, err := db.GetValue(ctx, "SELECT DATABASE()")
	if err != nil {
		return nil, fmt.Errorf("获取当前数据库名称出错: %v", err)
	}

	// 查询表字段详细信息
	query := `
        SELECT
            c.COLUMN_NAME as column_name,
            c.DATA_TYPE as data_type,
            c.COLUMN_TYPE as column_type,
            c.IS_NULLABLE as is_nullable,
            IFNULL(c.COLUMN_DEFAULT, '') as column_default,
            IFNULL(c.COLUMN_COMMENT, '') as column_comment,
            IFNULL(c.COLUMN_KEY, '') as column_key,
            IFNULL(c.EXTRA, '') as extra,
            c.ORDINAL_POSITION as ordinal_position,
            IFNULL(c.CHARACTER_MAXIMUM_LENGTH, 0) as character_maximum_length,
            IFNULL(c.NUMERIC_PRECISION, 0) as numeric_precision,
            IFNULL(c.NUMERIC_SCALE, 0) as numeric_scale
        FROM
            information_schema.COLUMNS c
        WHERE
            c.TABLE_SCHEMA = ? AND c.TABLE_NAME = ?
        ORDER BY
            c.ORDINAL_POSITION ASC
    `

	rows, err := db.GetAll(ctx, query, currentDb.String(), tableName)
	if err != nil {
		return nil, fmt.Errorf("查询表 %s 详细字段信息出错: %v", tableName, err)
	}

	// 处理查询结果
	for _, row := range rows {
		columnMeta := ColumnMeta{
			ColumnName:      row["column_name"].String(),
			DataType:        row["data_type"].String(),
			ColumnType:      row["column_type"].String(),
			IsNullable:      row["is_nullable"].String(),
			ColumnDefault:   row["column_default"].String(),
			ColumnComment:   row["column_comment"].String(),
			ColumnKey:       row["column_key"].String(),
			Extra:           row["extra"].String(),
			OrdinalPosition: row["ordinal_position"].Int(),
		}
		columns = append(columns, columnMeta)
	}

	return columns, nil
}

// ValidateTableExists 验证表是否存在
func ValidateTableExists(tableName string) (bool, error) {
	ctx := gctx.New()
	db := g.DB()

	// 获取当前数据库名称
	currentDb, err := db.GetValue(ctx, "SELECT DATABASE()")
	if err != nil {
		return false, fmt.Errorf("获取当前数据库名称出错: %v", err)
	}

	// 检查表是否存在
	query := `
        SELECT COUNT(*) as count
        FROM information_schema.TABLES
        WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?
    `

	count, err := db.GetValue(ctx, query, currentDb.String(), tableName)
	if err != nil {
		return false, fmt.Errorf("检查表 %s 是否存在时出错: %v", tableName, err)
	}

	return count.Int() > 0, nil
}

// GetAllTablesWithFilteredColumns 获取所有表和过滤后的字段信息
func GetAllTablesWithFilteredColumns() ([]TableWithColumnsInfo, error) {
	ctx := gctx.New()
	db := g.DB()
	result := make([]TableWithColumnsInfo, 0)

	// 获取当前数据库名称
	currentDb, err := db.GetValue(ctx, "SELECT DATABASE()")
	if err != nil {
		return nil, fmt.Errorf("获取当前数据库名称出错: %v", err)
	}

	// 定义需要过滤的表名称
	excludedTables := map[string]bool{
		"code_gen_record": true,
	}

	// 获取所有表信息
	tablesQuery := `
        SELECT
            TABLE_NAME AS table_name,
            TABLE_COMMENT AS table_comment
        FROM
            information_schema.TABLES
        WHERE
            TABLE_SCHEMA = ?
        ORDER BY
            TABLE_NAME ASC
    `
	tableRows, err := db.GetAll(ctx, tablesQuery, currentDb.String())
	if err != nil {
		return nil, fmt.Errorf("查询表信息出错: %v", err)
	}

	// 定义需要过滤的时间字段
	timeFields := map[string]bool{
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
	}

	// 为每个表获取字段信息
	for _, tableRow := range tableRows {
		tableName := tableRow["table_name"].String()
		tableComment := tableRow["table_comment"].String()

		// 过滤掉不需要的表
		if excludedTables[tableName] {
			continue
		}

		// 获取表字段信息
		columnsQuery := `
            SELECT
                COLUMN_NAME as column_name,
                DATA_TYPE as data_type,
                COLUMN_TYPE as column_type,
                IS_NULLABLE as is_nullable,
                IFNULL(COLUMN_DEFAULT, '') as column_default,
                IFNULL(COLUMN_COMMENT, '') as column_comment,
                IFNULL(COLUMN_KEY, '') as column_key,
                IFNULL(EXTRA, '') as extra
            FROM
                information_schema.COLUMNS
            WHERE
                TABLE_SCHEMA = ? AND TABLE_NAME = ?
            ORDER BY
                ORDINAL_POSITION ASC
        `

		columnRows, err := db.GetAll(ctx, columnsQuery, currentDb.String(), tableName)
		if err != nil {
			continue // 跳过有问题的表
		}

		// 过滤字段信息
		columns := make([]ColumnInfo, 0)
		for _, columnRow := range columnRows {
			columnName := columnRow["column_name"].String()

			// 过滤掉时间字段
			if timeFields[columnName] {
				continue
			}

			columnInfo := ColumnInfo{
				ColumnName:    columnName,
				DataType:      columnRow["data_type"].String(),
				ColumnType:    columnRow["column_type"].String(),
				IsNullable:    columnRow["is_nullable"].String(),
				ColumnDefault: columnRow["column_default"].String(),
				ColumnComment: columnRow["column_comment"].String(),
				ColumnKey:     columnRow["column_key"].String(),
				Extra:         columnRow["extra"].String(),
			}
			columns = append(columns, columnInfo)
		}

		// 只有当表有非时间字段时才添加到结果中
		if len(columns) > 0 {
			tableInfo := TableWithColumnsInfo{
				TableName:    tableName,
				TableComment: tableComment,
				Columns:      columns,
			}
			result = append(result, tableInfo)
		}
	}

	return result, nil
}

// ColumnInfo 字段信息结构体
type ColumnInfo struct {
	ColumnName    string `json:"columnName"`
	DataType      string `json:"dataType"`
	ColumnType    string `json:"columnType"`
	IsNullable    string `json:"isNullable"`
	ColumnDefault string `json:"columnDefault"`
	ColumnComment string `json:"columnComment"`
	ColumnKey     string `json:"columnKey"`
	Extra         string `json:"extra"`
}

// TableWithColumnsInfo 表及其字段信息结构体
type TableWithColumnsInfo struct {
	TableName    string       `json:"tableName"`
	TableComment string       `json:"tableComment"`
	Columns      []ColumnInfo `json:"columns"`
}
