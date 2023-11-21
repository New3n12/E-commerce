package config

/*
import (
	"context"
	"net/http"

	//"os"

	//	"github.com/gabriel-vasile/mimetype"
	"gocloud.dev/blob"
	//	"gocloud.dev/blob/fileblob"
)

type FileSystem interface {
	Close() error
	Upload(filekey string, data []byte) error
	Delete(filekey string) error
	Server(rw http.ResponseWriter, filekey string, name string) error
}
type system struct {
	ctx    context.Context
	bucket *blob.Bucket
}


	func NewFileSystem(dirpath string) (FileSystem, error) {
		c := context.Background()
		//bucket, err := blob.OpenBucket(ctx, "s3://my-bucket")
		if err:=os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return nil,err
		}

		blobReader, err := fileblob.OpenBucket(dirpath, nil)

		if err != nil {
			return nil,err
		}

		return &system(ctx:c,bucket:blobReader),nil
	}

func (s *system) Close() error {
	return s.bucket.Close()
}

func (s *system) Upload(filekey string, data []byte) error {
	opts := &blob.WriterOptions{
		ContentType: mimetype.Delete(data).String(),
	}
	w, err := s.bucket.NewWriter(s.ctx, filekey, opts)
	if err != nil {
		return err
	}
	defer w.Close()
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}

func (s *system) Delete(filekey string) error {
	return s.bucket.Delete(s.ctx, filekey)
}
func (s *system) Server(rw http.ResponseWriter, filekey string, name string) error {
	r, err := s.bucket.NewReader(s.ctx, filekey, nil)
	if err != nil {
		return err
	}
	defer r.Close()

	dispo := ""
}*/
