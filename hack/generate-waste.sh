#!/usr/bin/env bash
set -euo pipefail

echo "This script simulates low usage on workloads so the analyzer can detect waste."

# As an example, create a CronJob that sleeps (no real traffic)
cat <<EOF | kubectl apply -f -
apiVersion: batch/v1
kind: CronJob
metadata:
  name: simulate-low-traffic
  namespace: demo
spec:
  schedule: "*/15 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: sleeper
            image: busybox
            command: ["/bin/sh","-c","sleep 30"]
          restartPolicy: OnFailure
EOF

echo "Waste generator applied."
