# Get started

### Installation

1. Git Clone API Insights
   ```shell
   $ git clone https://wwwin-github.cisco.com/DevNet/api-insights.git
   $ cd api-insights
   ```

2. Deploy API Insights UI & Service in the `api-insights` namespace with Helm
   ```shell
   $ helm install api-insights . -n api-insights --create-namespace \
   --set api-insights.apiclarityUrl=http://apiclarity-apiclarity.apiclarity.svc.cluster.local
   ```

3. Wait until API Insights is Ready & Running
   ```shell
   $ kubectl wait --for=condition=ready pods --all --timeout=300s -n api-insights
   ```

4. Expose API Insights UI & Service

   ```shell
   $ kubectl port-forward -n api-insights svc/api-insights-api-insights 8080:8080
   ```
   
5. Check API Insights Service & Register your API as an API Insights Service
   ```shell
   # Check API Insights Service
   
   # API Insights CLI
   $ api-insights-cli analyzer ls -H http://localhost:8080
   
   # Curl
   $ curl -v http://localhost:8080/v1/apiregistry/services
   ```
   
   ```shell
   # Register your API as an API Insights Service (e.g. cases):
   
   # API Insights CLI
   $ api-insights-cli service create --data '{
     "organization_id": "CX",
     "product_tag": "CX Cloud",
     "name_id": "cases",
     "title": "Case Proxy API",
     "description": "Cases APIs for CX Cloud",
     "contact": {
     "name": "RMA Case Management Team",
       "url": "https://cisco.service-now.com/nav_to.do?uri=%2Fcmdb_ci_business_app.do%3Fsys_id%3D5a932084dba6505033df5ce2ca961910%26sysparm_record_list%3Dalias_operational_statusIN6,1%5Eba_name%3E%3DCase%5EORDERBYba_name%26sysparm_record_row%3D1%26sysparm_record_rows%3D6307%26sysparm_record_target%3Dsn_apm_cisco_business_application_and_modules%26sysparm_view%3DBusiness_Application%26sysparm_view_forced%3Dtrue",
       "email": "rma-case-mgmt-team@cisco.com"
     }
   }'
   
   # Curl
   $ curl --location --request POST 'http://localhost:8080/v1/apiregistry/services' \
   --header 'Content-Type: application/json' \
   --header 'Accept: application/json' \
   --data-raw '{
      "organization_id": "CX",
      "product_tag": "CX Cloud",
      "name_id": "cases",
      "title": "Case Proxy API",
      "description": "Cases APIs for CX Cloud",
      "contact": {
         "name": "RMA Case Management Team",
         "url": "https://cisco.service-now.com/nav_to.do?uri=%2Fcmdb_ci_business_app.do%3Fsys_id%3D5a932084dba6505033df5ce2ca961910%26sysparm_record_list%3Dalias_operational_statusIN6,1%5Eba_name%3E%3DCase%5EORDERBYba_name%26sysparm_record_row%3D1%26sysparm_record_rows%3D6307%26sysparm_record_target%3Dsn_apm_cisco_business_application_and_modules%26sysparm_view%3DBusiness_Application%26sysparm_view_forced%3Dtrue",
         "email": "rma-case-mgmt-team@cisco.com"
      }
   }'
   ```
   
7. Check API Insights UI in the browser: http://localhost:8080/


### Uninstallation
1. Uninstall API Insights UI & Service
   ```shell
   $ kubectl delete ns api-insights
   ```