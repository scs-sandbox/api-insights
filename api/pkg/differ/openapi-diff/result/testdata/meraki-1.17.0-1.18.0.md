#### What's New

##### `GET` /networks/{networkId}/webhooks/payloadTemplates

> List the webhook payload templates for a network

##### `POST` /networks/{networkId}/webhooks/payloadTemplates

> Create a webhook payload template for a network

##### `GET` /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}

> Get the webhook payload template for a network

##### `PUT` /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}

> Update a webhook payload template for a network

##### `DELETE` /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId}

> Destroy a webhook payload template for a network

#### What's Modified

##### `GET` /networks/{networkId}/appliance/uplinks/usageHistory

> Get the sent and received bytes for each uplink of a network.

###### Parameters:

Modified: `t1` in `query`
> The end of the timespan for the data. t1 can be a maximum of 31 days after t0.

Modified: `timespan` in `query`
> The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes.


##### `POST` /networks/{networkId}/webhooks/httpServers

> Add an HTTP server to a network

###### Request:

Modified content type: `application/json`

* Modified property `payloadTemplate` (object)
    > The payload template to use when posting data to the HTTP server.

    * Added property `name` (string)
        > The name of the payload template.


##### `POST` /networks/{networkId}/webhooks/webhookTests

> Send a test webhook for a network

###### Request:

Modified content type: `application/json`

* Added property `payloadTemplateName` (string)
    > The name of the payload template.


##### `PUT` /networks/{networkId}/wireless/ssids/{number}/vpn

> Update the VPN settings for the SSID

###### Request:

Modified content type: `application/json`

* Added property `failover` (object)
    > Secondary VPN concentrator settings. This is only used when two VPN concentrators are configured on the SSID.


##### `GET` /organizations

> List the organizations that the user has privileges on

###### Response:

Modified response: **200 OK**
> Successful operation

* Modified content type: `application/json`
    * Modified items (array):

        * Added property `licensing` (object)
            > Licensing related settings

        * Modified property `id` (string)
            > Organization ID

        * Modified property `name` (string)
            > Organization name

        * Modified property `url` (string)
            > Organization URL

        * Modified property `api` (object)
            > API related settings


##### `GET` /organizations/{organizationId}

> Return an organization

###### Response:

Modified response: **200 OK**
> Successful operation

* Modified content type: `application/json`
    * Added property `licensing` (object)
        > Licensing related settings

    * Modified property `api` (object)
        > API related settings

        * Modified property `enabled` (boolean)
            > Enable API access

    * Modified property `id` (string)
        > Organization ID

    * Modified property `name` (string)
        > Organization name

    * Modified property `url` (string)
        > Organization URL


##### `GET` /organizations/{organizationId}/devices/statuses

> List the status of every Meraki device in the organization

###### Parameters:

Deleted: `components` in `query`
> components


##### `GET` /organizations/{organizationId}/summary/top/clients/byUsage

> Return metrics for organization's top 10 clients by data usage (in mb) over given time range.

###### Response:

Modified response: **200 OK**
> Successful operation

* Deleted content type: `application/json`

##### `GET` /organizations/{organizationId}/webhooks/alertTypes

> Return a list of alert types to be used with managing webhook alerts

###### Parameters:

Added: `productType` in `query`
> Filter sample alerts to a specific product type


