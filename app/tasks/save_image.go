package tasks

import (
	"cloud.google.com/go/storage"
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h2non/bimg"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	_ "image/jpeg"
	"io"
	"log"
	"math"
)

const (
	maxWidth  = 1000
	maxHeight = 1200
)

func saveImage(c *gin.Context) {
	body := &lib.SaveImageRequest{}
	err := c.Bind(body)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	defer func(client *storage.Client) {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}(client)

	upload := client.Bucket(lib.FilesBucket).Object(lib.UploadsPrefix + body.ImageFileName)

	size, imgData, err := processImage(ctx, upload)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	hash, err := getHash(imgData)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	fileName := hash + ".jpg"

	result := client.Bucket(lib.ImagesBucket).Object(hash + ".jpg")
	w := result.NewWriter(ctx)
	w.ContentType = "image/jpeg"
	if _, err = w.Write(imgData); err != nil {
		lib.InternalError(err, c)
		return
	}
	if err = w.Close(); err != nil {
		lib.InternalError(err, c)
		return
	}
	if err := upload.Delete(ctx); err != nil {
		lib.InternalError(err, c)
		return
	}

	var imageTypes []data.ImageType
	if body.ImageType != "" {
		imageTypes = append(imageTypes, data.ImageType{
			Type: body.ImageType,
		})
	}

	if err = ds.SaveImage(&data.Image{
		ID:         fileName,
		Caption:    body.Caption,
		Location:   body.Location,
		Width:      size.Width,
		Height:     size.Height,
		ImageTypes: imageTypes,
		CreatedAt:  body.CreatedDate.Time,
	}); err != nil {
		lib.InternalError(err, c)
		return
	}
}

func processImage(ctx context.Context, obj *storage.ObjectHandle) (*bimg.ImageSize, []byte, error) {
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer func(reader *storage.Reader) {
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}(reader)
	buffer, err := io.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}

	img := bimg.NewImage(buffer)
	if _, err = img.AutoRotate(); err != nil {
		return nil, nil, err
	}

	size, err := img.Size()
	if err != nil {
		return nil, nil, err
	}
	width := size.Width
	height := size.Height

	if width > maxWidth {
		ratio := float64(height) / float64(width)
		width = maxWidth
		height = int(math.Round(float64(width) * ratio))
	}
	if height > maxHeight {
		ratio := float64(width) / float64(height)
		height = maxHeight
		width = int(math.Round(float64(height) * ratio))
	}

	imgData, err := img.Process(bimg.Options{
		Width:   width,
		Height:  height,
		Quality: 92,
	})
	if err != nil {
		return nil, nil, err
	}
	return &bimg.ImageSize{
		Width:  width,
		Height: height,
	}, imgData, nil
}

func getHash(data []byte) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write(data); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
