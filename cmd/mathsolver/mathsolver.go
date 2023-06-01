package main

import (
	"net/http"

	"github.com/phgermanov/ft-mathsolver/internal/handler"
	"github.com/phgermanov/ft-mathsolver/internal/middleware"
	"github.com/phgermanov/ft-mathsolver/internal/model"
)

func main() {
	mux := http.NewServeMux()
	// Create your dependencies
	errorRecorder := &model.MemoryErrorRecorder{
		ExpressionErrorMap: make(map[string]*model.Error),
	}

	// Pass the dependencies to the handlers
	handlers := &handler.Handler{
		ErrorRecorder: errorRecorder,
	}

	mux.Handle("/evaluate", middleware.LoggingMiddleware(http.HandlerFunc(handlers.EvaluateHandler)))
	mux.Handle("/validate", middleware.LoggingMiddleware(http.HandlerFunc(handlers.ValidateHandler)))
	mux.Handle("/errors", middleware.LoggingMiddleware(http.HandlerFunc(handlers.ErrorsHandler)))

	http.ListenAndServe(":8080", mux)
}
