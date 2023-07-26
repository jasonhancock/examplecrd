package main

import (
	"context"
	"fmt"
	"log"
	"os"

	clientv1 "github.com/jasonhancock/examplecrd/pkg/generated/clientset/versioned/typed/examplecrd.jasonhancock.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var config *rest.Config
	var err error
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		// we're running in-cluster. Try using the service account
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		log.Fatal(fmt.Errorf("getting kubeconfig: %w", err))
	}

	clientSet, err := clientv1.NewForConfig(config)
	if err != nil {
		log.Fatal(fmt.Errorf("getting vulture client: %w", err))
	}

	resp, err := clientSet.ExampleResources("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range resp.Items {
		fmt.Printf("%s %s %s\n", v.ObjectMeta.Name, v.Spec.Color, v.Spec.Size)
	}
}
