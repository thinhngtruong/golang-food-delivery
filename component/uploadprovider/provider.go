package uploadprovider

import (
	"context"
	"learn-api/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
