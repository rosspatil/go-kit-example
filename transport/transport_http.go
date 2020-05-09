package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/opentracing/opentracing-go"

	"github.com/go-kit/kit/log"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"

	"github.com/gin-gonic/gin"
	kithttp "github.com/go-kit/kit/transport/http"
	endpoint "github.com/rosspatil/go-kit-example/endpoint"
)

// NewHTTP - ...
func NewHTTP(e endpoint.Endpoint) *gin.Engine {
	g := gin.Default()
	g.POST("/", func(c *gin.Context) {
		kithttp.NewServer(e.Register,
			decodeRegisterRequest,
			encodeRegisterResponse).ServeHTTP(c.Writer, c.Request)
	})
	g.GET("/", func(c *gin.Context) {
		kithttp.NewServer(e.GetByID,
			decodeGetByIDRequest,
			encodeGetByIDResponse,
			kithttp.ServerBefore(kitopentracing.HTTPToContext(opentracing.GlobalTracer(), "Get", log.NewNopLogger())),
		).ServeHTTP(c.Writer, c.Request)
	})

	g.DELETE("/", func(c *gin.Context) {
		kithttp.NewServer(e.Delete, decodeDeleteRequest, encodeErrorOnlyResponse).ServeHTTP(c.Writer, c.Request)
	})
	g.PUT("/", func(c *gin.Context) {
		kithttp.NewServer(e.UpdateEmail, decodeUpdateEmailRequest, encodeErrorOnlyResponse).ServeHTTP(c.Writer, c.Request)
	})
	return g
}

func decodeRegisterRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.RegisterRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Employee); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeUpdateEmailRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.UpdateEmailRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	req.ID = r.URL.Query().Get("id")
	return req, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.GetRequest
	req.ID = r.URL.Query().Get("id")
	return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req endpoint.DeleteRequest
	req.ID = r.URL.Query().Get("id")
	return req, nil
}

func encodeGetByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp, ok := response.(endpoint.GetResponse)
	if ok && resp.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(`{"error":"` + resp.Error.Error() + `"}`))
		return err
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeErrorOnlyResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp, ok := response.(endpoint.ErrorOnlyResponse)
	if ok && resp.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(`{"error":"` + resp.Error.Error() + `"}`))
		return err
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeRegisterResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp, ok := response.(endpoint.RegisterResponse)
	if ok && resp.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(`{"error":"` + resp.Error.Error() + `"}`))
		return err
	}
	return json.NewEncoder(w).Encode(response)
}
