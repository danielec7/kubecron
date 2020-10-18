kubecron
========

A small utility to do things not currently possible with kubectl. 
At the moment the following commands are implemented:
- run
    - Running CronJobs immediately for testing purposes.
- suspend
    - Suspend a CronJob
- unsuspend
    - Unsuspend a CronJob

**Installation**

Download a release or

```bash
go get github.com/iJanki/kubecron
```

**Usage**

```bash
kubecron run cronjobname
```

It also accepts context and namespace flags as kubectl

```bash
kubecron --context=default-cluster -n default run cronjobname
```

**Usage as kubectl pluginn**

Rename the kubecron executable into kubectl-cron and place it
in your $PATH

```bash
mv kubecron /usr/local/bin/kubectl-cron
```

Run it like this:
```bash
kubectl cron run cronjobname
```

License
-------

This library is licensed under the MIT License - see the `LICENSE` file for details

