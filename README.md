spider
======

Collection of middlewares for use in go web servers

## Usage
###### Logging
```go
package main

import (
  "github.com/astephenb/spider"
  "net/http"
)

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func main() {
  http.Handle("/", spider.Log(handler{}))
	http.ListenAndServe(":8080", nil)
}
```
###### Output
```
[2014-11-14 14:11:22 CST] GET / 4.792us
[2014-11-14 14:11:22 CST] Parameters: {"test":["param1","param2"]}
```

