package controllers

import (
	"context"
	"fmt"

	finopsv1 "github.com/yourname/k8s-cost-optimizer/api/v1"
)

// Recommender turns analysis + metrics into CostRecommendation CRs

type Recommender struct {
	// TODO: add Kubernetes client and logger when wiring controller-runtime
}

func NewRecommender() *Recommender {
	return &Recommender{}
}

// Recommend creates CostRecommendation objects based on inputs.
func (r *Recommender) Recommend(ctx context.Context, rec *finopsv1.CostRecommendation) error {
	// TODO: implement creation/upsert of CostRecommendation CR
	fmt.Printf("Recommend: would create CostRecommendation for %s/%s\n", rec.Spec.Namespace, rec.Spec.ResourceName)
	return nil
}
