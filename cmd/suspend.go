package cmd

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/retry"
)

var suspendCmd = &cobra.Command{
	Use:   "suspend [cronjob name]",
	Short: "Suspend a cronjob",
	Long:  "Suspend a cronjob",
	Run:   suspend,
	Args:  cobra.MinimumNArgs(1),
}

var unSuspendCmd = &cobra.Command{
	Use:   "unsuspend [cronjob name]",
	Short: "Unsuspend a cronjob",
	Long:  "Unsuspend a cronjob",
	Run:   unsuspend,
	Args:  cobra.MinimumNArgs(1),
}

func suspend(_ *cobra.Command, args []string) {
	toggleSuspend(args, true)
}

func unsuspend(_ *cobra.Command, args []string) {
	toggleSuspend(args, false)
}

func toggleSuspend(args []string, status bool) {

	cronjobName := args[0]

	clientset := mustGetClientset()

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		cronjob := getCronjob(namespace, cronjobName)

		cronjob.Spec.Suspend = &status
		_, updateErr := clientset.BatchV1beta1().CronJobs(namespace).Update(context.TODO(), cronjob, metav1.UpdateOptions{})
		return updateErr
	})

	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Printf("Updated cronjob %q.\n", cronjobName)
}
