package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var ErrMissingParam = errors.New("missing parameter")

func GetUInt64Param(r *http.Request, name string) (uint64, error) {
	param := chi.URLParam(r, name)
	if param == "" {
		return 0, ErrMissingParam
	}

	return ParseUInt64(param)
}

func ParseUInt64(value string) (uint64, error) {
	return strconv.ParseUint(value, 10, 64)
}

func DecodeJSON(body io.Reader, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}
