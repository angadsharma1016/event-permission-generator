// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint "angadsharma1016/omega/src/participants/pkg/endpoint"
	http1 "angadsharma1016/omega/src/participants/pkg/http"
	service "angadsharma1016/omega/src/participants/pkg/service"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	opentracing "github.com/go-kit/kit/tracing/opentracing"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
)

func createService(endpoints endpoint.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger, tracer opentracinggo.Tracer) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"CreateFeedback": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateFeedback", logger))},
		"CreateP":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "CreateP", logger))},
		"DeleteP":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "DeleteP", logger))},
		"PostAttendance": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "PostAttendance", logger))},
		"ReadAttendance": {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ReadAttendance", logger))},
		"ReadFeedback":   {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ReadFeedback", logger))},
		"ReadP":          {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "ReadP", logger))},
		"UpdateP":        {http.ServerErrorEncoder(http1.ErrorEncoder), http.ServerErrorLogger(logger), http.ServerBefore(opentracing.HTTPToContext(tracer, "UpdateP", logger))},
	}
	return options
}
func addDefaultEndpointMiddleware(logger log.Logger, duration *prometheus.Summary, mw map[string][]endpoint1.Middleware) {
	mw["CreateP"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateP")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateP"))}
	mw["ReadP"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ReadP")), endpoint.InstrumentingMiddleware(duration.With("method", "ReadP"))}
	mw["UpdateP"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "UpdateP")), endpoint.InstrumentingMiddleware(duration.With("method", "UpdateP"))}
	mw["DeleteP"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "DeleteP")), endpoint.InstrumentingMiddleware(duration.With("method", "DeleteP"))}
	mw["CreateFeedback"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "CreateFeedback")), endpoint.InstrumentingMiddleware(duration.With("method", "CreateFeedback"))}
	mw["ReadFeedback"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ReadFeedback")), endpoint.InstrumentingMiddleware(duration.With("method", "ReadFeedback"))}
	mw["PostAttendance"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "PostAttendance")), endpoint.InstrumentingMiddleware(duration.With("method", "PostAttendance"))}
	mw["ReadAttendance"] = []endpoint1.Middleware{endpoint.LoggingMiddleware(log.With(logger, "method", "ReadAttendance")), endpoint.InstrumentingMiddleware(duration.With("method", "ReadAttendance"))}
}
func addDefaultServiceMiddleware(logger log.Logger, mw []service.Middleware) []service.Middleware {
	return append(mw, service.LoggingMiddleware(logger))
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"CreateP", "ReadP", "UpdateP", "DeleteP", "CreateFeedback", "ReadFeedback", "PostAttendance", "ReadAttendance"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}