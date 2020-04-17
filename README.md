# OpenShift Onboarding

[![Docker Repository on Quay](https://quay.io/repository/rira12621/openshift-onboarding-hugo/status "Docker Repository on Quay")](https://quay.io/repository/rira12621/openshift-onboarding-hugo)

## What this will do
Through a couple of units this will take you through setting up an Openshift Dedicated Cluster, deploying an example application and instrumenting it.
In the future I will add some extra units to break your cluster to learn more about an alerting strategy for your application.

## What's included
Each unit comes with an example go application, the corresponding dockerfile to build it and instructions and what to do with it.

## TODO
* add chaos monkey style unit
* convert to lab environment
* find place to host learning environment


## Running the project
We are make use of [Hugo][https://gohugo.io/] to simply run this.
The Dockerfile in the repository root is used to build a Hugo docker image, that is build and hosted on [quay.io][https://quay.io/rira12621/hugo-docker]

To run the website exectue `docker run --rm -ti --volume $PWD:/src quay.io/rira12621/hugo-docker:latest hugo`
