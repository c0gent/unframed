package unframed

import (
	"compress/gzip"
	"github.com/gorilla/pat"
	"io"
	"net/http"
	"strings"
)

func NewRouter() *Router {
	r := new(Router)
	r.pat = pat.New()
	return r
}

type Router struct {
	pat *pat.Router
}

func (r *Router) Get(Path string, HandlerFunc http.HandlerFunc) {
	r.pat.Get(Path, Zhhf(HandlerFunc))
}

func (r *Router) Post(Path string, HandlerFunc http.HandlerFunc) {
	r.pat.Post(Path, Zhhf(HandlerFunc))
}

func (r *Router) Dir(DirPath string) {
	UrlPath := "/" + DirPath
	r.pat.PathPrefix(UrlPath).Handler(Mah(http.StripPrefix(UrlPath, http.FileServer(http.Dir(DirPath)))))
}

func (r *Router) Serve(PortNo string) {
	http.Handle("/", r.pat)
	if err := http.ListenAndServe(":"+PortNo, nil); err != nil {
		panic(err)
	}
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
