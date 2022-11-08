# API Insights

### Prerequisites
* Install and configure the following tools:
    * [kubernetes](https://kubernetes.io/)
    * [kuebctl](https://kubernetes.io/docs/reference/kubectl/)
    * [helm](https://helm.sh/)
* Install an application for managing a local Kubernetes cluster and associated tools. For example, [Rancher Desktop](https://rancherdesktop.io/) allows you to locally set up Kubernetes, `kubectl` and `helm`. Other Kubernetes setup applications include *Minikube* and *Kind*.

### Deploy API Insights using Helm

1. Deploy API Insights
    ```shell
    helm install api-insights ./api-insights -n api-insights --create-namespace
    ```

2. Wait until API Insights is ready and running
    ```shell
    kubectl -n api-insights wait --for=condition=ready pods --all --timeout=300s
    ```
   
3. Expose the API Insights UI
    ```shell
    kubectl -n api-insights port-forward svc/api-insights 8080:8080
    ```

4. Open http://localhost:8080 to Access UI