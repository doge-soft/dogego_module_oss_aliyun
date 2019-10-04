package dogego_module_oss_aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func NewAliyunOSS() (*oss.Client, *oss.Bucket, error) {
	client, err := oss.New(os.Getenv("OSS_END_POINT"),
		os.Getenv("OSS_ACCESS_KEY_ID"),
		os.Getenv("OSS_ACCESS_KEY_SECRET"))
	
	if err != nil {
		return nil, nil, err
	}

	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))

	if err != nil {
		return nil, nil, err
	}

	return client, bucket, nil
}
