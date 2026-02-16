package pkg

import (
	"anarchive/entity"
	"os"
	"path/filepath"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetArchivePath(namespace string, bucketID int64, fileID entity.ID) string {
	bh := strconv.FormatInt(bucketID, 16)
	fh := fileID.String()
	b1, b2, b := bh[:2], bh[2:4], bh[4:]
	f1, f2, f := fh[:2], fh[2:4], fh[4:]
	return filepath.Join(namespace, b1, b2, b, f1, f2, f)
}
