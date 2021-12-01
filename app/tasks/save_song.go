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

	audioUpload := bucket.Object(lib.UploadsPrefix + body.AudioFileName)
	wavFile := bucket.Object(musicPrefix + body.SongName + ".wav")
	if _, err := wavFile.CopierFrom(audioUpload).Run(ctx); err != nil {
		lib.InternalError(err, w)
		return
	}
	if err := audioUpload.Delete(ctx); err != nil {
		lib.InternalError(err, w)
		return
	}

	reader, err := wavFile.NewReader(ctx)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	defer func(reader *storage.Reader) {
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}(reader)

	mp3File := bucket.Object(musicPrefix + body.SongName + ".mp3")
	writer := mp3File.NewWriter(ctx)
	writer.ContentType = "audio/mpeg"

	enc := lame.NewEncoder(writer)
	defer enc.Close()

	if _, err = bufio.NewReader(reader).WriteTo(enc); err != nil {
		lib.InternalError(err, w)
		return
	}

	if err := writer.Close(); err != nil {
		lib.InternalError(err, w)
		return
	}

	mp3Acl := mp3File.ACL()
	if err := mp3Acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		lib.InternalError(err, w)
		return
	}

	imageUpload := bucket.Object(lib.UploadsPrefix + body.ImageFileName)
	imageFile := bucket.Object(musicPrefix + body.SongName + ".jpg")
	if _, err := imageFile.CopierFrom(imageUpload).Run(ctx); err != nil {
		lib.InternalError(err, w)
		return
	}
	imageAcl := imageFile.ACL()
	if err := imageAcl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		lib.InternalError(err, w)
		return
	}
	if err = imageUpload.Delete(ctx); err != nil {
		lib.InternalError(err, w)
		return
	}
	if err = db.SaveSong(body.SongName, body.Description); err != nil {
		lib.InternalError(err, w)
		return
	}
}
