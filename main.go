// kubecron is a small utility to help managing kubernetes CronJobs.
// - Run a CronJob for test purposes.
// - Suspend/unsuspend a CronJob
package main

import "github.com/iJanki/kubecron/cmd"

func main() {
	cmd.Execute()
}
