# OpenShift Onboarding

[![Docker Repository on Quay](https://quay.io/repository/rira12621/openshift-onboarding-hugo/status "Docker Repository on Quay")](https://quay.io/repository/rira12621/openshift-onboarding-hugo)
![Build and Deploy to Cloud Run](https://github.com/RiRa12621/openshift-onboarding/workflows/Build%20and%20Deploy%20to%20Cloud%20Run/badge.svg)


## What this will do
Through a couple of units this will take you through setting up an Openshift Dedicated Cluster, deploying an example application and instrumenting it.
In the future I will add some extra units to break your cluster to learn more about an alerting strategy for your application.

## What's included
Each unit comes with an example go application, the corresponding dockerfile to build it and instructions and what to do with it.

## TODO
- [] add chaos monkey style unit
- [] convert to lab environment
- [] find place to host learning environment


## Running the project
We are make use of [Hugo](https://gohugo.io/) to simply run this.
The Dockerfile in the repository root is used to build a Hugo docker image, that is build and hosted on [quay.io](https://quay.io/rira12621/openshift-onboarding-hugo)

To run the website exectue `docker run --rm -p 1313:1313 -ti quay.io/rira12621/openshift-onboarding-hugo:latest`

## Adding more Units
This is a living project, if you feel something is missing, please add it.
Hugo is making this really easy.
You can follow the the [Hugo documentation](https://gohugo.io/getting-started/) but the tl;dr is the folling command:

```
hugo new unit/<your unit>
```
Please stick to the template for units included in the root of this repository

Also make sure to remove `draft: true` when you're ready to have it published.

We have github actions set up for build and deployment once, your changes are merged to master.


## Notes
While this is intended for OpenShift dedicated Admins/Users, most of the Units can be done on OpenShift Container Platform as well, if you don't have access to a dedicated cluster.
We are happy to extend this guide to cover OCP as well, PRs welcome.
