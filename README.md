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

License
-------

This library is licensed under the MIT License - see the `LICENSE` file for details

