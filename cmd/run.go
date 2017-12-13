package cmd

import (
	"fmt"
	"time"
	"strconv"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
)

var runCmd = &cobra.Command{
	Use:   "run [cronjob name]",
	Short: "Run a cronjob immediately",
	Long:  "Run a cronjob immediately",
	Run:   run,
}

func run(_ *cobra.Command, args []string) {

	cronjob := args[0]

	config, err := buildConfigFromFlags(kubecontext, kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	cj, err := clientset.BatchV1beta1().CronJobs(namespace).Get(cronjob, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Cronjob not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting Cronjob %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {

		jobsClient := clientset.BatchV1().Jobs(namespace)

		suffix := "-" + strconv.Itoa(int(time.Now().Unix()))

		job := &batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Name:      cronjob + suffix,
				Namespace: namespace,
			},
			Spec: cj.Spec.JobTemplate.Spec,
		}

		result, err := jobsClient.Create(job)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created job %q.\n", result.GetObjectMeta().GetName())
	}

}

func buildConfigFromFlags(context, kubeconfigPath string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
