package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Printf("kubeconfig file: %s", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error build config: %v", config)
	}
	log.Printf("k8s config, host: %v", config.Host)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("error create clientset: %v", err)
	}

	api := clientset.CoreV1()

	log.Println("namespaces list:")
	options := metav1.ListOptions{}
	nsList, err := api.Namespaces().List(context.TODO(), options)
	if err != nil {
		log.Fatalf("error get list of namespaces: %v", err)
	}

	for _, ns := range nsList.Items {
		log.Printf("ns: %v", ns.Name)
	}

}
