# hello-operator

Author: Brian Tomlinson <darthlukan@gmail.com> / <btomlins@redhat.com>


## Description

Just a simple operator so that I can practice with the [Operator Framework](https://github.com/operator-framework).


## Usage

```
$ oc new-project hello-operator --description="A place for Greetings"

$ oc create -f deploy/crds/greeting_v1alpha1_greeting_crd.yaml

$ oc create -f deploy/service_account.yaml
$ oc create -f deploy/role.yaml
$ oc create -f deploy/role_binding.yaml
$ oc create -f deploy/operator.yaml

$ oc create -f deploy/crds/greeting_v1alpha1_greeting_cr.yaml
$ oc logs example-greeting-pod
>> "Hello World!"
```
