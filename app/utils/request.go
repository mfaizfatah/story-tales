package utils

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*Curl using for request to other service / other API / API External
* @parameter
* ctx, is context from http.Request
* url, url endpoint for api external
* method, method using for method request endpoint
* timer is timeout calling endpoint if didn't get response
* header, header needed when reequest endpoint
* payload, payload is data needed to send when request endpoint
*
* @represent
* context for write any response and request to 3rd party
* []byte response body but return with byte for unmarshal
* int response statuscode when request to endpoint
* error is error interface from golang
 */
func Curl(ctx context.Context, url, method string, timer time.Duration, header map[string]string, payload *bytes.Buffer) (context.Context, []byte, int, error) {
	var req *http.Request
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	start := time.Now().In(loc)

	if method == http.MethodGet {
		request, err := http.NewRequest(method, url, nil)
		if err != nil {
			return ctx, nil, http.StatusInternalServerError, err
		}
		req = request
	} else {
		request, err := http.NewRequest(method, url, payload)
		if err != nil {
			return ctx, nil, http.StatusInternalServerError, err
		}
		req = request
	}

	client := http.Client{
		Timeout: timer,
	}

	for key, val := range header {
		req.Header.Set(key, val)
	}

	if err != nil {
		ctx = logger.RecordThridPartyFailed(ctx, req, start, http.StatusInternalServerError, err.Error())
		return ctx, nil, http.StatusInternalServerError, err
	}

	res, err := client.Do(req)
	if err != nil {
		ctx = logger.RecordThridPartyFailed(ctx, req, start, res.StatusCode, err.Error())
		return ctx, nil, http.StatusInternalServerError, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		ctx = logger.RecordThridPartyFailed(ctx, req, start, http.StatusInternalServerError, err.Error())
		return ctx, nil, http.StatusInternalServerError, err
	}

	ctx = logger.RecordThridParty(ctx, req, start, res.StatusCode, body)
	return ctx, body, res.StatusCode, nil
}

//NewRouterPlus router with health check
func NewRouterPlus() *chi.Mux {

	router := chi.NewRouter()

	router.Get("/metrics", promhttp.Handler().ServeHTTP)

	return router
}
