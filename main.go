package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	kubeletEnvPath = "/etc/kubernetes/kubelet.env"
	envKey         = "KUBELET_IMAGE_TAG"
)

func main() {
	kubeconfig := ""
	imagever := ""
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "path to kubeconfig")
	flag.StringVar(&imagever, "imagever", imagever, "static image version to be written")
	flag.Parse()
	if kubeconfig == "" {
		log.Printf("kubeconfig empty, writing from env variable : %s", imagever)
		err := setValueInFile(imagever)
		if err != nil {
			log.Fatalf("unable to set image version : %s", err.Error())
		}
	} else {
		log.Printf("setting from server")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("unable to build client : %s", err.Error())
		}
		serverversion, err := client.ServerVersion()
		if err != nil {
			log.Fatalf("unable to get version from server : %s", err.Error())
		}
		version := strings.Replace(serverversion.GitVersion, "+", "_", -1)
		err = setValueInFile(version)
		if err != nil {
			log.Fatalf("unable to set image version : %s", err.Error())
		}
	}
}

func setValueInFile(version string) error {
	log.Printf("setting version : %s", version)
	log.Printf("creating file : %s", kubeletEnvPath)
	file, err := os.Create(kubeletEnvPath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "%s=%s\n", envKey, version)
	if err != nil {
		return err
	}
	return nil
}
