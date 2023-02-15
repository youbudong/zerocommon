package upload

import (
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenarateQiniuUploadToken(accessKey, secretKey, bucket, prefix, ossURL string) string {
	auth := qbox.NewMac(accessKey, secretKey)

	putPolicy := storage.PutPolicy{
		Scope:        bucket,
		SaveKey:      generateKey(prefix) + "${ext}",
		ForceSaveKey: true,
		ReturnBody: `{
			"key":"$(key)",
			"hash":"$(etag)",
			"fsize":$(fsize),
			"bucket":"$(bucket)",
			"fname":"$(fname)",
			"mimeType":"$(mimeType)",
			"width":$(imageInfo.width),
			"height":$(imageInfo.height),
			"format":"$(imageInfo.format)",
			"url":"` + ossURL + `$(key)",
			"ext":"$(ext)"}`,
	}
	return putPolicy.UploadToken(auth)
}

func generateKey(prefix string) string {
	return prefix + time.Now().Format("200601/02/") + primitive.NewObjectID().Hex()
}
