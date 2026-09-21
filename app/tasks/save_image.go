package tasks

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"math"

	"cloud.google.com/go/storage"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
)

const (
	maxWidth  = 2000
	maxHeight = 2400
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

	originalData, err := readImage(ctx, upload)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	size, previewData, err := processImage(originalData)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	images := client.Bucket(lib.ImagesBucket)
	fileName, err := saveImageFile(ctx, images, originalData)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	previewID, err := saveImageFile(ctx, images, previewData)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	var tags []data.Tag
	for _, tagName := range body.Tags {
		if tagName != "" {
			tags = append(tags, data.Tag{Name: tagName})
		}
	}

	if err = ds.SaveImage(&data.Image{
		ID:        fileName,
		PreviewID: previewID,
		Caption:   body.Caption,
		Location:  body.Location,
		Width:     size.Width,
		Height:    size.Height,
		Tags:      tags,
		CreatedAt: body.CreatedDate.Time,
		Camera:    body.Camera,
		Lens:      body.Lens,
		Film:      body.Film,
	}); err != nil {
		lib.InternalError(err, c)
		return
	}
	if err := upload.Delete(ctx); err != nil {
		// The image is saved; a cleanup failure must not retry the database insert.
		log.Println(err)
	}
}

func saveImageFile(ctx context.Context, bucket *storage.BucketHandle, imgData []byte) (string, error) {
	hash, err := getHash(imgData)
	if err != nil {
		return "", err
	}
	fileName := hash + ".jpg"
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	w := bucket.Object(fileName).NewWriter(ctx)
	w.ContentType = "image/jpeg"
	if _, err := w.Write(imgData); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}
	return fileName, nil
}

func readImage(ctx context.Context, obj *storage.ObjectHandle) ([]byte, error) {
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer func(reader *storage.Reader) {
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}(reader)
	return io.ReadAll(reader)
}

func processImage(buffer []byte) (*vips.ImageMetadata, []byte, error) {
	// Keep the uploaded bytes intact, including metadata, for original storage.
	img, err := vips.NewImageFromBuffer(buffer)
	if err != nil {
		return nil, nil, err
	}
	defer img.Close()

	// govips transforms decoded pixels; JPEG is encoded only at final export.
	if err := img.AutoRotate(); err != nil {
		return nil, nil, err
	}
	if err := img.TransformICCProfile(vips.SRGBIEC6196621ICCProfilePath); err != nil {
		return nil, nil, err
	}
	sourceWidth, sourceHeight := img.Width(), img.Height()
	width, height := sourceWidth, sourceHeight

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

	width, height = max(1, width), max(1, height)
	if width != sourceWidth || height != sourceHeight {
		// Match the rounded target dimensions on both axes.
		if err := img.ResizeWithVScale(float64(width)/float64(sourceWidth),
			float64(height)/float64(sourceHeight), vips.KernelLanczos3); err != nil {
			return nil, nil, err
		}
	}
	// Screen sharpening: sigma=0.5, X1=2, M2=2; defaults M1=0, Y2=10, Y3=20.
	if err := img.Sharpen(0.5, 2, 2); err != nil {
		return nil, nil, err
	}
	params := vips.NewJpegExportParams()
	params.Quality = 95
	params.OptimizeCoding = true
	imgData, metadata, err := img.ExportJpeg(params)
	if err != nil {
		return nil, nil, err
	}
	return metadata, imgData, nil
}

func getHash(data []byte) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write(data); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
