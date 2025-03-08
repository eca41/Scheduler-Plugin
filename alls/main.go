package main

import (
	"os"

	"github.com/eca41/plugins"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	command := app.NewSchedulerCommand(
		app.WithPlugin(plugins.Name, plugins.New),
	)

	if err := command.Execute(); err != nil {
		klog.Fatalf("Error running scheduler: %v", err)
		os.Exit(1)
	}
}
