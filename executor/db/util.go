package db

import (
	"database/sql"

	s "github.com/aak1247/AssertTiDB/generator/sql"
)

func Scan(rows sql.Rows, columns []s.Column) []map[string]s.ColumnData {
	res := make([]map[string]s.ColumnData, 0)
	defer rows.Close()
	for rows.Next() {
		//定义变量接收查询数据
		data := make([]interface{}, len(columns))
		for i := range data {
			curData := columns[i].GenerateColumnRefData().Data
			data[i] = curData
		}
		err := rows.Scan(data...)
		if err != nil {
			panic(err.Error())
		}
		curRes := make(map[string]s.ColumnData, len(columns))
		for i, datum := range data {
			curData := s.ColumnData{Name: columns[i].Name, Data: datum}
			curData.DeRef()
			curRes[columns[i].Name] = curData
		}
		res = append(res, curRes)
	}
	return res
}
