package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		if err == mongo.ErrNoDocuments {
			body.Msg = "Not Found"
		} else {
			body.Msg = err.Error()
		}
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
