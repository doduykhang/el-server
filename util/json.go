package util

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
)

func NewJSONEncoder(w io.Writer) (enc *json.Encoder) {
	enc = json.NewEncoder(w)
	enc.SetEscapeHTML(JSONEscapeHTML)
	enc.SetIndent(JSONPrefix, JSONIndent)
	return
}

func SetContentTypeJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SetContentTypeText(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
}

func JSONResp(v interface{}, status int, w http.ResponseWriter) error {
	SetContentTypeJSON(w)
	w.WriteHeader(status)

	if vv := reflect.ValueOf(v); vv.Kind() == reflect.Slice && vv.IsZero() {
		_, err := w.Write([]byte("[]"))
		return err
	}
	return NewJSONEncoder(w).Encode(v)
}
