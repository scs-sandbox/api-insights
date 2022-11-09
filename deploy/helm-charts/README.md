# API Insights

### Prerequisites
* Install and configure the following tools:
    * [kubernetes](https://kubernetes.io/)
    * [kuebctl](https://kubernetes.io/docs/reference/kubectl/)
    * [helm](https://helm.sh/)
* Install an application for managing a local Kubernetes cluster and associated tools. For example, [Rancher Desktop](https://rancherdesktop.io/) allows you to locally set up Kubernetes, `kubectl` and `helm`. Other Kubernetes setup applications include *Minikube* and *Kind*.

### Deploy API Insights using Helm

1. Deploy API Insights
    Option1: If you want to use the prebuild public images:
    ```shell
    helm install api-insights ./api-insights -n api-insights --create-namespace
    ```
    Option2: If you want to use the images built from source code locally:
    ```shell
    # build images
    docker-compose build
    # tag the images built above
    docker tag api-insights-backend ghcr.io/cisco-developer/api-insights-api:local
    docker tag api-insights-ui ghcr.io/cisco-developer/api-insights-ui:local
    # install
    helm install api-insights ./api-insights -n api-insights --create-namespace \
        --set api-insights.docker.backendImageTag=local \
        --set api-insights.docker.frontendImageTag=local \
        --set api-insights.docker.imagePullPolicy=Never
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
5. Delete all if needed.
    ```shell
    kubectl delete ns api-insights
    ```
