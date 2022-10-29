package file

import (
	"testing"

	"github.com/beshrek/pan/conf"
)

func TestUpload(t *testing.T) {
	fileUploader := NewUploader(conf.TestData.AccessToken, conf.TestData.Path, conf.TestData.LocalFilePath)
	res, err := fileUploader.Upload()
	if err != nil {
		t.Fail()
	} else {
		t.Logf("TestUpload Success res: %+v", res)
	}
}
