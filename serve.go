package serve

import (
 "net/http"
 "compress/gzip"
)

type gzipResponseWriterStruct struct {
	http.ResponseWriter
	gz *gzip.Writer
}
func (w gzipResponseWriterStruct) Write(b []byte) (int, error) {
	return w.gz.Write(b)
}
func Gzip(w http.ResponseWriter, r *http.Request, fn func(http.ResponseWriter, *http.Request)) {
	w.Header().Set(`Content-Encoding`, `gzip`)
	gz, _ := gzip.NewWriterLevel(w, 1)
	fn(gzipResponseWriterStruct{w, gz}, r)
	gz.Close()
}
