# k8s-mutating-webhook
This repository has the implementation of the admission controller's mutating webhook from scratch.

This example mutating webhook adds the label `example-webhook: "it-worked"` on every new pod that has a label `example-webhook-enabled: "true"`.

# Steps to run

1. Clone the repository and navigate into the folder

    ```
    git clone https://github.com/AxiomSamarth/k8s-mutating-webhook
    cd k8s-mutating-webhook
    ```

2. Apply all the manifests on to your K8s cluster
    ```
    kubectl create -f ./manifests/
    ```

3. The new pod created out of the `manifests/demo-pod.yaml` would have an additional label in the `metadata` as `example-webhook: "it-worked"` If this has appeared then everything has worked well.

## Issues/bugs

Feel free to raise issues and PRs if you see any scope to improve this example repository for wider consumption. Happy helming!