package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/yourname/k8s-cost-optimizer/controllers"
)

func main() {
	var prometheusAddr string
	flag.StringVar(&prometheusAddr, "prometheus-addr", "http://localhost:9090", "Prometheus base URL")
	flag.Parse()

	fmt.Println("k8s-cost-optimizer scaffold starting")

	promClient, err := controllers.NewPrometheusClient(prometheusAddr)
	if err != nil {
		fmt.Printf("warning: could not create prometheus client: %v\n", err)
	} else {
		// quick test query (placeholder)
		_ = promClient
	}

	// start a simple ticker to demonstrate background work (replace with controller-runtime later)
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Analyzer tick — TODO: implement analysis loop and reconcile CostRecommendation CRs")
	}
}
