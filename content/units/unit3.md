---
title: "Unit 3"
description: "Starting to instrument"
date: 2020-04-22T16:40:09+02:00
weight: 30
---

## What's happening in this unit?
You will start to instrument your application.


## Action Items
* expose metrics basic metircs
* access prometheus, grafana, alertmanger
* access your applications metrics

### Task 1

Check out the code for this unit.
You will find that the Dockerfile itself did not actually change: We are still just building and exposing.

The interesting bits are in the actual go application.
In addition to the handling of basic request, we now serve some metrics on `/metrics`.
You will find some more explanation in the code comments.

But a tl;dr is that there is a [prometheus library](https://prometheus.io/docs/guides/go-application/) that we are making use of.
You can find the corresponding documentation [here](https://prometheus.io/docs/guides/go-application/).
Those are basic metrics but good enough for us as our app is not doing a whole lot right now.

The task here is just to check out the code and understand what's happening.
Feel free to change things and play around with it locally.

## Task 2
This is just for you to understand where the various components are and how to find them
Your OpenShift Cluster comes with the [Cluster Monitoring Operato](https://github.com/openshift/cluster-monitoring-operator) and that includes Prometheus, Alertmanager and Grafana/.
Additionally there's some neat magic coming from [Prometheus Operator](https://github.com/coreos/prometheus-operator/) as well as [Kubernetest Mixins](https://github.com/kubernetes-monitoring/kubernetes-mixin)

You can discover those either from the UI or by using the CLI. We'll be going the CLI way.
All you need is the following command:

```
oc get route -n openshift-monitoring
```

This will show you the routes including the URLs to pull either of the Web UIs.
Visit them and familiarize yourself with the UI a bit in case you've never used them.

## Task 3

While every OpenShift cluster comes with a prometheus operator stack, that is only used to specifically observer and monitor the cluster components.
The `cluster-monitoring-operator` has some built in functionality that enables that.
Workarounds for that exist but are temporary and should not be part of this guide.
At the time of writing "User Workload Monitoring" has not yet landed on OpenShift so we will have to take a bit of a detour.

There are two ways to possibly achieve this, we will only touch the official way here.

Since "User Workload Monitoring", while not available officially, is available as a tech preview feature.
As we are admins on our clusters, we can happily enable those.
We want to also give a little warning here: you might not want to do this on your production clusters.
There's some more documentation on possible risks [here](https://docs.openshift.com/container-platform/4.1/nodes/clusters/nodes-cluster-enabling-features.html).

Let's move ahead, ignoring the risk.

Lucky us, the monitoring team has written some really nice [documentation](https://docs.openshift.com/container-platform/4.3/monitoring/monitoring-your-own-services.html) that we're going base this task off.
