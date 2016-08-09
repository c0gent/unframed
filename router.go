package unframed

import (
	"compress/gzip"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strings"
	//"github.com/c0gent/unframed/log"
)

func NewRouter() *Router {
	r := new(Router)
	r.Router = mux.NewRouter()
	return r
}

type Router struct {
	*mux.Router
}

func (r *Router) Get(Path string, HandlerFunc http.HandlerFunc) {
	r.Router.HandleFunc(Path, Zhhf(HandlerFunc)).Methods("GET")
}
func (r *Router) Post(Path string, HandlerFunc http.HandlerFunc) {
	r.Router.HandleFunc(Path, Zhhf(HandlerFunc)).Methods("POST")
}
func (r *Router) Dir(DirPath string) {
	UrlPath := "/" + DirPath
	r.Router.PathPrefix(UrlPath).Handler(Mah(http.StripPrefix(UrlPath, http.FileServer(http.Dir(DirPath)))))
}
func (r *Router) Serve(PortNo string) {
	http.Handle("/", r.Router)
	if err := http.ListenAndServe(":"+PortNo, nil); err != nil {
		panic(err)
	}
}

func (r *Router) Subrouter(pathPrefix string) *Router {
	return &Router{r.Router.PathPrefix(pathPrefix).Subrouter()}
}

func (n *NetHandle) QueryUrl(s string, r *http.Request) (i int) {
	requestVars := mux.Vars(r)
	//temp_vars_map[s]
	i = Atoi(requestVars[s])
	//log.Message(s)
	//log.Message(temp_vars_map[s])
	return
}

func (n *NetHandle) UrlVar(s string, r *http.Request) (v string) {
	v = mux.Vars(r)[s]
	return
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	//sniffDone bool
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {

	/*if !w.sniffDone {
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", http.DetectContentType(b))
		}
		w.sniffDone = true
	}*/

	return w.Writer.Write(b)
}

//ZipFileHandler
func Zfh(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//w.Header().Set("Access-Control-Allow-Origin", "*")

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		/*
			if r.Method == "HEAD" {
				h.ServeHTTP(w, r)
				return
			}
		*/

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		h.ServeHTTP(&gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})
}

//ZipHtmlHandlerFunc
func Zhhf(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "text/html")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		fn(&gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	}
}

//maxAgeHandler
func Mah(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=31536000")
		h.ServeHTTP(w, r)
	})
}
