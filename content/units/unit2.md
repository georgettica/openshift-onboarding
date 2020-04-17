---
title: "Unit 2"
date: 2020-04-17T15:17:00+02:00
draft: true
---

# Unit 2

## What's happening in this unit?
You will set up your own OpenShift dedicated cluster and deploy and example application to it

## Action items

* Set up OpenSift Dedicated Cluster using OCM
* Deploy the example docker image to your newly set up OpenShift Dedicated Cluster

### Note
We encourage you to always take a look at the code even if it's not explicitly mentioned.
In this case we have changed the code so that it has some extra functionality over just printing a statement once.
If you browse it, you will find that we've added capilities to answer to http requests, making it a lot easier to deploy on OpenShift.
In case you are curious, you can build it like we did in the first lection and then run it.
It will listen on port 8080.

### Task 1
These steps assume you have enough quota to start a cluster.
If this is not the case and you are a Red Hat employee, talk to your manager to get this fixed.
If this is not the case and you are a customer, talk to your Sales contact to buy more quota.

* log in to https://cloud.redhat.com
* navigate to "cluster manager"
* hit "create cluster"
* choose OpenShift Dedicated
* choose the cloud provider you prefer, we will assume that is Google Cloud Platform
* Fill out the form (since this is a training lab, the smallest possible cluster is good enough)
* Make yourself a nice cup of coffee or tea while the cluster is being set up

To help you navigate through this a bit easier, there's a gif below showing the flow.


### Task 2
You will be deploying the code and docker image from this repository just that we have uploaded it to quay.io for you already.


#### Preparations
There are a variety of different ways this can be achieved, we will try to stay as OpenShift native therefore those instructions might not work with vanilla Kubernetes
as things like `oc create` don't exist there.

Before you begin you should have `oc` set up already.
You can download it from [here][https://mirror.openshift.com/pub/openshift-v4/clients/oc/latest/].

In case you prefer to use a CLI to interact with OCM instead of the UI, you can also download the do so. Just get the command line client using `go get -u github.com/openshift-online/ocm-cli/cmd/ocm`.
Afterwards get your offline token from [here][https://cloud.redhat.com/openshift/token] and then login using `ocm login --token $OFFLINE_TOKEN`.

Now you will be able to use commands like `ocm cluster list`, `ocm cluster list --managed` or login like `ocm cluster login <cluster_id>`

Log in to your newly created cluster with your preferred method.

Verify everything worked with `oc whoami` which should tell you which user you are on the cluster.

#### Task fulfilling

Now that you're logged into the cluster, we can go ahead and deploy our application.
First though we have to create a new project.

`oc new-project my-great-project`

Theoretically your context will be set to that new project but let's just be extra sure

`oc project my-great-project`

Now let's deploy our application.
`oc new-app my-great-app:quay.io/rira12621/openshift-onboarding:unit2`

