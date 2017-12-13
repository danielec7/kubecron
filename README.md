kubecron
========

A small utility to do things not currently possible with kubectl. For example running cronjobs immediately for testing purposes.

**Installation**

No releases yet. Use:

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

