package media_library

import (
	"database/sql/driver"
	"image"
	"io"
	"mime/multipart"
	"strings"

	"github.com/jinzhu/gorm"

	"os"
)

type Size struct {
	Width  int
	Height int
}

type MediaLibrary interface {
	Scan(value interface{}) error
	Value() (driver.Value, error)

	GetURLTemplate(*Option) string
	GetURL(option *Option, scope *gorm.Scope, field *gorm.Field) string

	GetFileHeader() *multipart.FileHeader
	GetFileName() string

	GetSizes() map[string]Size
	SetCropOption(*CropOption)
	GetCropOption() *image.Rectangle

	Store(url string, option *Option, reader io.Reader) error
	Retrieve(url string) (*os.File, error)

	IsImage() bool

	URL(style ...string) string
	String() string
}

type Option map[string]string

func (option Option) Get(key string) string {
	return option[key]
}

func parseTagOption(str string) *Option {
	tags := strings.Split(str, ";")
	setting := Option{}
	for _, value := range tags {
		v := strings.Split(value, ":")
		k := strings.TrimSpace(strings.ToUpper(v[0]))
		if len(v) == 2 {
			setting[k] = v[1]
		} else {
			setting[k] = k
		}
	}
	return &setting
}
