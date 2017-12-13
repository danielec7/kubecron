package cmd

import (
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/util/homedir"
	"github.com/spf13/cobra"
	"path/filepath"
)

var namespace, kubecontext, kubeconfig string

var rootCmd = &cobra.Command{
	Use: "kubecron [cronjob]",
	Short: "Utilities for kubernetes cronjobs",
	Long: "Utilities for kubernetes cronjobs",
	Args: cobra.MinimumNArgs(1),
}

func Execute() {
	rootCmd.Execute()
}


func init() {
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", apiv1.NamespaceDefault, "Namespace")
	rootCmd.PersistentFlags().StringVar(&kubecontext, "context", "", "Context")
	if home := homedir.HomeDir(); home != "" {
		rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}

	rootCmd.AddCommand(runCmd)
}
