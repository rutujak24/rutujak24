#!/usr/bin/env bash
set -euo pipefail

echo "This script deploys demo workloads that are intentionally overprovisioned for the demo."

# Example deployment (replace with real manifests)
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: demo-overprovisioned
  namespace: demo
spec:
  replicas: 3
  selector:
    matchLabels:
      app: demo
  template:
    metadata:
      labels:
        app: demo
    spec:
      containers:
      - name: web
        image: nginx:stable
        resources:
          requests:
            memory: "1Gi"
            cpu: "500m"
          limits:
            memory: "2Gi"
            cpu: "1"
EOF

echo "Demo deployment applied."
