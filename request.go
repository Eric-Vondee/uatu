package uatu

import "net/http"

type GenericRequest struct{}

func (g GenericRequest) Bind(_ *http.Request) error { return nil }
