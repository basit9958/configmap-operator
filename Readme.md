# Configmap Operator
Configmap operator helps to manage ConfigMaps cluster-wide.

## Usage
Make sure you have a Kubernetes cluster running before using the operator.
* Clone the repository.
* To install the CRD (`DemoClusterConfigmap`) into the cluster: `make install`
* To run the controller: `make run ENABLE_WEBHOOKS=false`

### Using Docker
* `make deploy bh3139/configmap-operator:0.1`

---
* Create a Namespace:
    ```yaml
    # ns1.yaml
    apiVersion: v1
    kind: Namespace
    metadata:
        name: ns1
        labels:
            managed: "dev1"
    ```
  `kubectl apply -f ns1.yaml`

* Create a ClusterConfigMap object: `kubectl apply -f example/test1.yaml`