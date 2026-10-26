package main

import (
	"context"
	"fmt"
	"os"

	kubeclient "github.com/DevalGarg129/k8s-autoscaler/internal/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	clientset, err := kubeclient.NewClient()

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	//Query K8s API Server
	node, err := clientset.CoreV1().Nodes().List(
		context.Background(),
		metav1.ListOptions{},
	)

	if err != nil {
		fmt.Println("Failed to get nodes: ", err)
		os.Exit(1)
	}

	fmt.Println("Connected to Kubernetes")
	fmt.Println("Cluster Nodes: ")

	for _, node := range node.Items {
		fmt.Println("-", node.Name)
	}
}
