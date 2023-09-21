package file_uploader_lib_go

import (
	"bytes"
	clientConfig "file_uploader_lib/client_config"
	dbConfig "file_uploader_lib/db_config"
	"github.com/tealeg/xlsx"
	"log"
	"time"
)

func saveDataFromUploader(fileObjName string, columnList []string, tableName string, dbDsn string, minioHost string, minioAccessKey string, minioSecretKey string, bucketName string) (bool, error) {
	conn, err := dbConfig.Connect(dbDsn)
	queryStr := dbConfig.PrepareQuery(tableName, columnList)
	if err != nil {
		log.Fatalln(err)
		return false, nil
	}
	err, minioClient := clientConfig.MinioConnect(minioHost, minioAccessKey, minioSecretKey, false)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	file, err := clientConfig.GetFileObject(bucketName, fileObjName, minioClient)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	xlsxReader := bytes.NewReader(file)
	xlFile, err := xlsx.OpenReaderAt(xlsxReader, int64(len(file)))
	if err != nil {
		log.Fatalln(err)
		return false, err
	}
	batch := [][]interface{}{}
	for _, sheet := range xlFile.Sheets {
		for idx, row := range sheet.Rows {
			values := []interface{}{}
			for _, col := range row.Cells {
				if idx != 0 {
					values = append(values, col.String())
				}
			}
			if len(values) > 0 {
				// Adding Created_at and updated_at fields values
				values = append(values, time.Now()) // updated_at
				values = append(values, time.Now()) //created_at
				batch = append(batch, values)
			}

		}
	}
	for _, values := range batch {
		_, err := conn.Exec(queryStr, values...)
		if err != nil {
			log.Fatalln(err)
			return false, err
		}
	}
	log.Println("Data inserted successfully")

	return true, nil

}
