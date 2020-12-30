/**
 * Created by zc on 2020/6/6.
 */
package ctr

import (
	"encoding/json"
	"net/http"
)

// Success writes ok message to the response.
func Success(w http.ResponseWriter) {
	Str(w, "success")
}

// Bytes writes the Bytes message to the response.
func Bytes(w http.ResponseWriter, bytes []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// Str writes the string message to the response.
func Str(w http.ResponseWriter, str string) {
	Bytes(w, []byte(str))
}

// OK writes the json-encoded data to the response.
func OK(w http.ResponseWriter, v interface{}) {
	JSON(w, v, http.StatusOK)
}

// NoContent writes no content to the response.
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Redirect replies to the request with a redirect to url.
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, url, code)
}

// Found replies to the request with a redirect to url.
func Found(w http.ResponseWriter, r *http.Request, url string) {
	Redirect(w, r, url, http.StatusFound)
}

// ErrorCode writes the json-encoded error message to the response.
func ErrorCode(w http.ResponseWriter, err error, status int) {
	Log.Errorln(err)
	JSON(w, err.Error(), status)
}

// InternalError writes the json-encoded error message to the response
// with a 500 internal server error.
func InternalError(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusInternalServerError)
}

// NotImplemented writes the json-encoded error message to the
// response with a 501 not found status code.
func NotImplemented(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusNotImplemented)
}

// NotFound writes the json-encoded error message to the response
// with a 404 not found status code.
func NotFound(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusNotFound)
}

// Unauthorized writes the json-encoded error message to the response
// with a 401 unauthorized status code.
func Unauthorized(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusUnauthorized)
}

// Forbidden writes the json-encoded error message to the response
// with a 403 forbidden status code.
func Forbidden(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusForbidden)
}

// BadRequest writes the json-encoded error message to the response
// with a 400 bad request status code.
func BadRequest(w http.ResponseWriter, err error) {
	ErrorCode(w, err, http.StatusBadRequest)
}

// JSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
