package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/api/batch/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var namespace, kubecontext, kubeconfig string

var rootCmd = &cobra.Command{
	Use:   "kubecron [cronjob]",
	Short: "Utilities for kubernetes cronjobs",
	Long:  "Utilities for kubernetes cronjobs",
	Args:  cobra.MinimumNArgs(1),
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
	rootCmd.AddCommand(suspendCmd)
	rootCmd.AddCommand(unSuspendCmd)
}

func mustGetClientset() *kubernetes.Clientset {
	config, err := buildConfigFromFlags(kubecontext, kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}

func getCronjob(namespace, cronjob string) *v1beta1.CronJob {

	clientset := mustGetClientset()

	cj, err := clientset.BatchV1beta1().CronJobs(namespace).Get(cronjob, metav1.GetOptions{})

	if errors.IsNotFound(err) {
		fmt.Printf("Cronjob not found\n")
		os.Exit(1)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting Cronjob %v\n", statusError.ErrStatus.Message)
		os.Exit(1)
	} else if err != nil {
		panic(err.Error())
	}

	return cj
}
