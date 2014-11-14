package spider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05 MST"

func Log(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		params, _ := json.Marshal(r.URL.Query())
		fmt.Printf("[%s] %v %s %v\n", t2.Format(TimeFormat), r.Method, r.URL.Path, t2.Sub(t1))
		fmt.Printf("[%s] %s %s\n", t2.Format(TimeFormat), "Parameters:", string(params))
	}

	return http.HandlerFunc(fn)
}
