# file-uploader-lib-go

### Pre-requisites
    [1] Golang
    [2] Minio server to store files

### Process To Use

    [1] import the github url in your project's import ()
    [2] Use saveDataFromUploader() method to save data from xlsx file to database.
        [...] Above method takes multiple parameters: 
            [file_object_name]: pass file object name that is stored in minio server in a a specific bucket.
            [column_list]: pass list of column names of the table where you want to save data.
            [table_name]: pass table_name string in which table you want to save data.
            [db_dsn]: pass database data source name for connect database.
            [minio_host]: pass minio host value for connect minio server.
            [minio_access_key_name]: pass minio access key name for connect minio server.
            [minio_secret_key_name]: pass minio secret key name for connect minio server.
            [minio_bucket_name]: pass minio bucket name for get file object from the specific bucket in connected minio server.


