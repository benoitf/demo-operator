## Demo Operator

This is a demo operator created with `operator-sdk` command. It has been slightly modified to create a service and a route. Thus, it works on OpenShift only.

The operator creates an **nginx** pod with a service and a route to preview Welcome to Nginx page (this piece of code is commented and should be uncommented during development in a Che workspace).

Custom resource spec has `routeName` field that you can use to create a route with a specific name.
There's also `url` field in custom resource status, that you can update once a route is created.

## Pre reqs:

Golang, docker

## How to build binary

In the root of this repo:

```$shell
go build -o build/demo-operator cmd/manager/main.go 
```

## How to build image

In the root of this repo:


```$shell
go build -o build/demo-operator cmd/manager/main.go 
docker build -t $registry/$org/$repo:tag -f build/Dockerfile .
```

Use the resulted image in deploy/operator.yaml. By default, it uses `eivantsov/demo-operator` image.

## How to deploy an operator

Cluster-admin privileges are required to run the following commands:

```$bash
# register a custom resource
oc apply -f deploy/crds/app_v1alpha1_appservice_crd.yaml

# create service account, role and rolebinding
oc apply -f deploy/rbac.yaml

# create operator deployment
oc apply -f deploy/operator.yaml

# create a custom resource which will trigger operator controller
oc apply -f deploy/crds/app_v1alpha1_appservice_cr.yaml
```