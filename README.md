# KuberScale

KuberScale is an early-stage Kubernetes autoscaler project. The current executable is a connectivity check: it creates a Kubernetes client, lists cluster nodes, and prints their names. It does **not** make scaling decisions or change cluster capacity yet.

## Requirements

- Go 1.26.6 or later
- Access to a running Kubernetes cluster
- A kubeconfig file with credentials for that cluster
- Permission to list cluster nodes

## Configure access

The application loads credentials from the file named by `KUBECONFIG`. If that variable is unset, it looks for `~/.kube/config`.

For example, to use a kubeconfig stored outside the default location:

```sh
export KUBECONFIG="$HOME/path/to/kubeconfig"
```

Make sure the kubeconfig refers to a reachable cluster and the active identity can list nodes. The app uses kubeconfig-based access; it does not currently load in-cluster service-account credentials.

## Run

From the repository root:

```sh
go run ./cmd/autoscaler
```

On success, the program prints `Connected to Kubernetes` followed by the node names. If kubeconfig is missing, invalid, or cannot reach the cluster, it exits with an error. If the identity lacks node-list permission, the Kubernetes API request fails.

## Test

```sh
go test ./...
```

## Project layout

- `cmd/autoscaler/` — executable entry point; requests and prints the cluster's nodes.
- `internal/kubernetes/` — creates a client from `KUBECONFIG` or the default kubeconfig path.
