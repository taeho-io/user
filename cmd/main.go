package main

import (
	"os"
	"sync"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/taeho-io/user/server"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func main() {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		panic(err)
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		panic(err)
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer func() {
		err := closer.Close()
		if err != nil {
			panic(err)
		}
	}()

	logrus.SetFormatter(&logrus.JSONFormatter{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		addr := ":80"
		log := logrus.WithField("addr", addr)

		cfg := server.NewConfig(server.NewSettings())

		log.Info("String User gRPC server")
		if err := server.Serve(addr, cfg); err != nil {
			log.Error(err)
			return
		}
	}()

	go func() {
		defer wg.Done()

		addr := ":81"
		log := logrus.WithField("addr", addr)

		cfg := server.NewConfig(server.NewSettings())

		log.Info("String User gRPC server")
		if err := server.Serve(addr, cfg); err != nil {
			log.Error(err)
			return
		}
	}()

	wg.Wait()
	os.Exit(1)
}
