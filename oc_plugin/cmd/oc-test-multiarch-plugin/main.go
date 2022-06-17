package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	// Printing out the OS arguments
	for _, val := range os.Args {
		fmt.Printf("OS Argument: %s", val)
		fmt.Println()
	}
	fmt.Println()

	// Assembling the path to .kube
	fmt.Printf("The Home Directory for Kube %s", homedir.HomeDir())
	fmt.Println()

	// Load the Configuration
	kubeconfig := homedir.HomeDir() + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// context.TODO() is an empty context.
	nodes, _ := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	fmt.Println("The list of the nodes:")
	for _, node := range nodes.Items {
		fmt.Printf("%s", node.Name)
		fmt.Println()
	}
}
