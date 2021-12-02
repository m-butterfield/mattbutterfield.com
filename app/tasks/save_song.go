package tasks

import (
	"bufio"
	"cloud.google.com/go/storage"
	"context"
	"encoding/json"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/viert/go-lame"
	"log"
	"net/http"
)

const (
	musicPrefix = "music/"
)

func SaveSong(w http.ResponseWriter, r *http.Request) {
	body := &lib.SaveSongRequest{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	defer func(client *storage.Client) {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}(client)
	bucket := client.Bucket(lib.FilesBucket)

	if err = copyAndConvertAudio(ctx, bucket, body.AudioFileName, body.SongName); err != nil {
		lib.InternalError(err, w)
		return
	}

	if err = copyImageUpload(ctx, bucket, body.ImageFileName, body.SongName); err != nil {
		lib.InternalError(err, w)
		return
	}

	if err = db.SaveSong(body.SongName, body.Description); err != nil {
		lib.InternalError(err, w)
		return
	}
}

func copyAndConvertAudio(ctx context.Context, bucket *storage.BucketHandle, fileName, songName string) error {
	audioUpload := bucket.Object(lib.UploadsPrefix + fileName)
	wavFile := bucket.Object(musicPrefix + songName + ".wav")
	if _, err := wavFile.CopierFrom(audioUpload).Run(ctx); err != nil {
		return err
	}

	reader, err := wavFile.NewReader(ctx)
	if err != nil {
		return err
	}
	defer func(reader *storage.Reader) {
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}(reader)

	mp3File := bucket.Object(musicPrefix + songName + ".mp3")
	writer := mp3File.NewWriter(ctx)
	writer.ContentType = "audio/mpeg"

	enc := lame.NewEncoder(writer)
	defer enc.Close()

	if _, err = bufio.NewReader(reader).WriteTo(enc); err != nil {
		return err
	}

	if err = writer.Close(); err != nil {
		return err
	}
	if err = deleteUploadAndMakePublic(ctx, audioUpload, mp3File.ACL()); err != nil {
		return err
	}
	return nil
}

func copyImageUpload(ctx context.Context, bucket *storage.BucketHandle, fileName, songName string) error {
	imageUpload := bucket.Object(lib.UploadsPrefix + fileName)
	imageFile := bucket.Object(musicPrefix + songName + ".jpg")
	if _, err := imageFile.CopierFrom(imageUpload).Run(ctx); err != nil {
		return err
	}
	if err := deleteUploadAndMakePublic(ctx, imageUpload, imageFile.ACL()); err != nil {
		return err
	}
	return nil
}

func deleteUploadAndMakePublic(ctx context.Context, upload *storage.ObjectHandle, acl *storage.ACLHandle) error {
	if err := upload.Delete(ctx); err != nil {
		return err
	}
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}
	return nil
}
