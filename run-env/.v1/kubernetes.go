package main

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func loadYAML(namespace, instance string, always_import_yaml bool) (map[string]string, error) {
	if namespace == "" || instance == "" {
		return nil, errors.New("namespace and instance must be provided")
	}

	if !always_import_yaml {
		cfgIni, err := loadFileIni("~/temp/" + namespace + "." + instance + ".ini")
		if err == nil {
			return cfgIni.Section("Custom").KeysHash(), nil
		}
	}

	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		kubeconfigPath = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}

	rest.InClusterConfig()

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, errors.New("failed to build kubeconfig: " + err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.New("failed to create Kubernetes client: " + err.Error())
	}

	client, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), instance, metav1.GetOptions{})
	if err != nil {
		return nil, errors.New("failed to get instance: " + err.Error())
	}

	if len(client.Spec.Template.Spec.Containers) == 0 {
		return nil, errors.New("no containers found in the deployment")
	}

	var envs = map[string]string{}
	for _, container := range client.Spec.Template.Spec.Containers {
		for _, env := range container.Env {
			envs[env.Name] = env.Value
		}

		if len(container.Env) != 0 {
			break
		}
	}

	if !always_import_yaml {
		err = saveFileIni("~/temp/"+namespace+"."+instance+".ini", envs)
		if err != nil {
			return nil, errors.New("failed to save YAML config: " + err.Error())
		}
	}

	return envs, nil
}
