package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/metric/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var (
	lemonsKey = attribute.Key("ex.com/lemons")
)

func initMeter() {
	exporter, err := prometheus.InstallNewPipeline(prometheus.Config{})
	if err != nil {
		log.Panicf("failed to initialize prometheus exporter %v", err)
	}

	server := &http.Server{
		Addr:         ":8443",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig(),
		//TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	http.HandleFunc("/metrics", exporter.ServeHTTP)

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Protocol: %s", r.Proto)))
	})

	go func() {

		if err := server.ListenAndServeTLS("", ""); err != nil {
			log.Fatal(err)
		}

	}()

	fmt.Println("Prometheus server running on :8443")
}

func main() {

	initMeter()

	meter := global.Meter("ex.com/basic")
	observerLock := new(sync.RWMutex)
	observerValueToReport := new(float64)
	observerLabelsToReport := new([]attribute.KeyValue)
	cb := func(_ context.Context, result metric.Float64ObserverResult) {
		(*observerLock).RLock()
		value := *observerValueToReport
		labels := *observerLabelsToReport
		(*observerLock).RUnlock()
		result.Observe(value, labels...)
	}
	_ = metric.Must(meter).NewFloat64ValueObserver("ex.com.one", cb,
		metric.WithDescription("A ValueObserver set to 1.0"),
	)

	valuerecorder := metric.Must(meter).NewFloat64ValueRecorder("ex.com.two")
	counter := metric.Must(meter).NewFloat64Counter("ex.com.three")

	commonLabels := []attribute.KeyValue{lemonsKey.Int(10), attribute.String("A", "1"), attribute.String("B", "2"), attribute.String("C", "3")}
	notSoCommonLabels := []attribute.KeyValue{lemonsKey.Int(13)}

	ctx := context.Background()

	(*observerLock).Lock()
	*observerValueToReport = 1.0
	*observerLabelsToReport = commonLabels
	(*observerLock).Unlock()
	meter.RecordBatch(
		ctx,
		commonLabels,
		valuerecorder.Measurement(2.0),
		counter.Measurement(12.0),
	)

	time.Sleep(5 * time.Second)

	(*observerLock).Lock()
	*observerValueToReport = 1.0
	*observerLabelsToReport = notSoCommonLabels
	(*observerLock).Unlock()
	meter.RecordBatch(
		ctx,
		notSoCommonLabels,
		valuerecorder.Measurement(2.0),
		counter.Measurement(22.0),
	)

	time.Sleep(5 * time.Second)

	(*observerLock).Lock()
	*observerValueToReport = 13.0
	*observerLabelsToReport = commonLabels
	(*observerLock).Unlock()
	meter.RecordBatch(
		ctx,
		commonLabels,
		valuerecorder.Measurement(12.0),
		counter.Measurement(13.0),
	)

	fmt.Println("Example finished updating, please visit :8443")

	select {}

}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("/root/certs/prometheus.crt")
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile("/root/certs/prometheus.key")
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}
