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
	log.Printf("kubeconfig: %s", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error build config: %v", config)
	}
	log.Println("k8s config created")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("error create clientset: %v", err)
	}

	api := clientset.CoreV1()

	options := metav1.ListOptions{}
	nsList, err := api.Namespaces().List(context.TODO(), options)
	log.Printf("%v", nsList)

}
