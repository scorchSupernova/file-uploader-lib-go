package file_uploader_lib_go

import (
	"testing"
)

func TestUpload(t *testing.T) {
	// TODO: Need to pass verified table column list
	var columnList = []string{}

	// TODO: Need to fill below parameters to create db dsn
	dbDsn := "postgres://{username}:{password}@{host}:{port}/{db_name}?sslmode={sql_mode}"
	// TODO: Need to fill below parameters to connect minio
	minioHost := "{localhost}:{port}"
	minioAccessKey := "{minio_access_key}"
	minioSecretKey := "{minio_secret_key}"
	bucketName := "{minio_bucket_name}"
	//TODO: Need to pass all preffered parameters to upload data from excel file to database table
	done, err := saveDataFromUploader("{file_object_name}", columnList, "{table_name}", dbDsn, minioHost, minioAccessKey, minioSecretKey, bucketName)
	if err != nil {
		t.Errorf("Failed ! got %v", err)
	}
	t.Logf("Success! got %v", done)
}
