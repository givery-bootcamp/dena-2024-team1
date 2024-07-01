package filestorage

import "mime/multipart"

type FileStorage interface {
	UploadFile(file *multipart.File, fileName string) error
}
