package cloudinary

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
)

type CloudinaryService struct {
	Client *cloudinary.Cloudinary
}

func New(cloud, key, secret string) (*CloudinaryService, error) {
	cld, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		return nil, err
	}

	return &CloudinaryService{
		Client: cld,
	}, nil
}

func (c *CloudinaryService) Ping() error {
	_, err := c.Client.Admin.Ping(context.Background())
	return err
}
