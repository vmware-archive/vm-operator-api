# vm-operator-api

## Overview

The vm-operator-api project provides the object model and generated client libraries
for the VM Operator project, which is a part of vSphere's [Kubernetes](https://kubernetes.io)
support.

VM Operator allows users to manage the lifecycle of vSphere Virtual Machines using
a declarative Kubernetes consumption model and is an integral part of Project Pacific
in vSphere 7 with Kubernetes.

The state of VirtualMachine CRDs in a Kubernetes cluster is monitored by built-in controllers
which reconcile the CRD specification into real Virtual Machines along with the necessary
storage, networking and compute dependencies. The state of the Virtual Machines is reflected
back in the VirtualMachine CRD status which can be queried for essential state like IP Address(es).

## Use cases

The use cases of vm-operator-api are currently limited to 3rd party integrations with vSphere with Kubernetes.

In vSphere with Kubernetes it is not currently possible to create new VirtualMachines using this API, but we
hope to expand on this functionality over time.

## Contributing

We welcome new contributors who are interested in collaborating on our Kubernetes support for vSphere

More information [here](CONTRIBUTING.md)

## Getting Started

Check out how to get started with vm-operator-api [here](GETTING-STARTED.md)

## Roadmap

Use of the vm-operator-api is currently only supported in vSphere with Kubernetes.
However the intention is to make it more widely available as a Kubernetes-native
way of interacting with vSphere VMs.

## Maintainers

Current list of project maintainers [here](MAINTAINERS.md)

## License

Antrea is licensed under the [Apache License, version 2.0](LICENSE)
