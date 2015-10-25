package energy

import (
    "log"
    "net/http"
    "time"
    "encoding/json"

    "github.com/gorilla/mux"
)

type JsonErr struct {
    Code int    `json:"code"`
    Text string `json:"text"`
}

type Handler func(*http.Request) (int, interface{})

type Route struct {
    Method      string
    Pattern     string
    HandlerFunc Handler
}

type Routes []Route

func (f Handler) HandleRequest(r *http.Request) (int, interface{}) {
    return f(r)
}

func GenericHandler(inner Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var data interface{}
        var code int

        code, data = inner.HandleRequest(r)
        content, err := json.Marshal(data)

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        if err != nil {
            w.WriteHeader(500)
            content = []byte(`{"message": "Something unexpected happened"}`)
        } else {
            w.WriteHeader(code)
        }
        w.Write(content)
    })
}

func Logger(inner http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "(%s)\t%s\t%s",
            time.Since(start),
            r.Method,
            r.RequestURI,
        )
    })
}

func NewApplication(routes Routes) *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler

        handler = GenericHandler(route.HandlerFunc)
        handler = Logger(handler)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Handler(handler)
    }

    return router
}
