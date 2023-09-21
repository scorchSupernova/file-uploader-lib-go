package db_config

import (
	"log"
	"strconv"
	"strings"
)

func PrepareQuery(tableName string, columnList []string) string {
	var idExtractedColumnList = make([]string, 0)
	var valuesList = make([]string, 0)
	for idx, column := range columnList {
		if column != "id" {
			idExtractedColumnList = append(idExtractedColumnList, column)
			newIdx := "$" + strconv.Itoa(idx)
			valuesList = append(valuesList, newIdx)
		}
	}

	queryStr := "INSERT INTO " + tableName + " (" + strings.Join(idExtractedColumnList[:], ", ") + ") VALUES (" + strings.Join(valuesList[:], ", ") + ")"
	log.Println("Query Preparation Done...")
	return queryStr
}
