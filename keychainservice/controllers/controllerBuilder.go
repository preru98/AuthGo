package controllers

import "net/http"

type HandlerBuilder struct {
}

func NewHandlerBuilder() HandlerBuilder {
	return HandlerBuilder{}
}

func (h HandlerBuilder) AddAuthentication() HandlerBuilder {

	return h
}

func (h HandlerBuilder) AddValidation() HandlerBuilder {
	return h
}

func (h HandlerBuilder) Get() HandlerBuilder {
	return h
}

// BUILD
func (h HandlerBuilder) Build() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {}
}
