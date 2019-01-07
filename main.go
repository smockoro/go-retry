package retry

import (
	"net/http"
	"time"
)

type Errors []error

func Do(c *http.Client, options ...Option) (http.Response, Errors) {
	config := InitConfig()

	// オプションを更新
	for _, update := range options {
		update(config)
	}

	errors := make(Errors, config.attempts)

	for n := 0; n < config.attempts; n++ {
		// 実行

		// リトライ判定

		// リトライ強制解除

		// リトライ時間の決定
		time.Sleep(config.retryAlgoritm())
	}

	return resp, errors
}
