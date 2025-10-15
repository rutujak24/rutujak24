# k8s-cost-optimizer

Kubernetes Cost Optimizer — a kubebuilder-style operator scaffold that analyzes cluster resources, queries Prometheus, and creates CostRecommendation CRs with actionable FinOps guidance.

This repository is a scaffold to accelerate building a custom Kubernetes operator that detects over-provisioning and unused resources and suggests savings.

Key features (planned):
- Detect pods with requests >> usage
- Detect over-provisioned node pools
- Detect unused PVCs
- Calculate monthly savings using cloud pricing APIs (AWS/GCP/Azure)
- Create `CostRecommendation` CRDs
- Optional auto-apply safe recommendations
- Grafana dashboard for cluster cost and savings

Quick start (development):

1. Install Go 1.20+ and kubebuilder if you plan to run generator tooling.
2. From the project root:

```bash
# build the controller
make build

# run locally against a cluster (kubeconfig required)
make run
```

3. To generate CRD manifests later (requires controller-gen / kubebuilder):

```bash
make generate
make install-crd
```

What's included
- `cmd/controller/main.go` — manager bootstrap
- `api/v1` — CRD type for `CostRecommendation`
- `controllers/` — analyzer, prometheus client and recommender scaffolds
- `config/` — CRD, RBAC and sample manifests
- `dashboards/` — Grafana dashboard JSON placeholder
- `docs/` — installation guide and architecture diagram (placeholder)
- `hack/` — demo deployment and waste generation scripts

Notes
- This repository is a scaffold. Many functions are TODO and require implementation (Prometheus queries, cost calculations, CRD generation).
- Replace `github.com/yourname/k8s-cost-optimizer` in `go.mod` with your desired module path.
