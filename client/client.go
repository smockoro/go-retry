package retry

import (
	"net/http"
	"time"
)

type Errors []error

type RetryClient interface {
	Do(req *http.Request, options ...Option) (*http.Response, Errors)
}

type retryClient struct {
	c *http.Client
}

func NewRetryClient(c *http.Client) RetryClient {
	return &retryClient{
		c: c,
	}
}

//func Retry(hf http.HandleFunc) http.HandleFunc {
//return func(w http.ResponseWriter, r *http.Request) {
//
//}
//}

func (rc *retryClient) Do(req *http.Request, options ...Option) (*http.Response, Errors) {
	config := InitConfig()

	// オプションを更新
	for _, update := range options {
		update(config)
	}

	var errors Errors

	for n := 0; n < config.Attempts; n++ {
		// 実行
		resp, err := rc.c.Do(req)

		// リトライ判定
		if err == nil {
			return resp, errors
		}
		errors[n] = err

		// リトライ時間の決定
		time.Sleep(config.RetryAlgoritm())
	}

	return nil, errors
}
