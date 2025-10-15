# Installation

This document will explain how to install the k8s-cost-optimizer controller in a Kubernetes cluster. It's a placeholder — update with real instructions once CRDs and manifests are generated.

1. Build the controller binary:

```bash
make build
```

2. Generate CRDs (requires controller-gen / kubebuilder):

```bash
make generate
make install-crd
```

3. Deploy manifests from `config/` (RBAC, deployment, CRDs).

4. Configure Prometheus and Grafana to scrape metrics and import the dashboard in `dashboards/`.
