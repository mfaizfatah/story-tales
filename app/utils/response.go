package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
)

// response using for sending response to frontend
type response struct {
	Status       string      `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

// Response is sending data response to frontend
func Response(ctx context.Context, w http.ResponseWriter, status bool, statuscode int, data interface{}) {
	var (
		res  response
		resp string
	)

	if !status {
		res.Status = "error"
		res.ErrorMessage = data.(string)
		res.Data = ""

		resp = data.(string)
	} else {
		d, _ := json.Marshal(data)
		/* temporary disable if not necesarry
		// func for encrypt response
		encrypt, err := encryption.Encrypt(d)
		if err != nil {
			panic(err)
		}
		*/
		res.Status = "success"
		res.ErrorMessage = ""
		res.Data = data

		resp = string(d)
	}

	datares, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	go logger.EndRecord(ctx, resp, statuscode)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	fmt.Fprintf(w, string(datares))
}

// HTMLResponse for response html
func HTMLResponse(w http.ResponseWriter, statuscode int, AddForm string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statuscode)
	fmt.Fprint(w, AddForm)
	return
}
