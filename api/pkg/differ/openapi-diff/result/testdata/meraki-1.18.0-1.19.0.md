#### What's New

##### `POST` /organizations/{organizationId}/inventory/claim

> Claim a list of devices, licenses, and/or orders into an organization

##### `GET` /organizations/{organizationId}/inventory/devices

> Return the device inventory for an organization

##### `GET` /organizations/{organizationId}/inventory/devices/{serial}

> Return a single device from the inventory of an organization

##### `POST` /organizations/{organizationId}/inventory/release

> Release a list of claimed devices from an organization.

##### `GET` /organizations/{organizationId}/sensor/readings/history

> Return all reported readings from sensors in a given timespan, sorted by timestamp

##### `GET` /organizations/{organizationId}/sensor/readings/latest

> Return the latest available reading for each metric from each sensor, sorted by sensor serial

#### What's Deleted

##### `GET` /organizations/{organizationId}/inventoryDevices

> Return the device inventory for an organization

##### `GET` /organizations/{organizationId}/inventoryDevices/{serial}

> Return a single device from the inventory of an organization

#### What's Modified

##### `PUT` /devices/{serial}/switch/ports/{portId}

> Update a switch port

###### Request:

Modified content type: `application/json`

* Added property `adaptivePolicyGroupId` (string)
    > The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.

* Added property `peerSgtCapable` (boolean)
    > If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port.
    > Cannot be applied to a port on a switch bound to profile.


##### `POST` /networks/{networkId}/camera/qualityRetentionProfiles

> Creates new quality retention profile for this network.

###### Request:

Modified content type: `application/json`

* Modified property `videoSettings` (object)
    > Video quality and resolution settings for all the camera models.

    * Added property `MV52` (object)
        > Quality and resolution for MV52 camera models.


##### `PUT` /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}

> Update an existing quality retention profile for this network.

###### Request:

Modified content type: `application/json`

* Modified property `videoSettings` (object)
    > Video quality and resolution settings for all the camera models.

    * Added property `MV52` (object)
        > Quality and resolution for MV52 camera models.


##### `POST` /networks/{networkId}/wireless/rfProfiles

> Creates new RF profile for this network

###### Request:

Modified content type: `application/json`

* Added property `perSsidSettings` (object)
    > Per-SSID radio settings by number.


##### `PUT` /networks/{networkId}/wireless/rfProfiles/{rfProfileId}

> Updates specified RF profile for this network

###### Request:

Modified content type: `application/json`

* Added property `perSsidSettings` (object)
    > Per-SSID radio settings by number.


##### `PUT` /networks/{networkId}/wireless/ssids/{number}

> Update the attributes of an MR SSID

###### Request:

Modified content type: `application/json`

* Added property `secondaryConcentratorNetworkId` (string)
    > The secondary concentrator to use when the ipAssignmentMode is 'VPN'. If configured, the APs will switch to using this concentrator if the primary concentrator is unreachable. This param is optional. ('disabled' represents no secondary concentrator.)

* Added property `disassociateClientsOnVpnFailover` (boolean)
    > Disassociate clients when 'VPN' concentrator failover occurs in order to trigger clients to re-associate and generate new DHCP requests. This param is only valid if ipAssignmentMode is 'VPN'.

* Added property `speedBurst` (object)
    > The SpeedBurst setting for this SSID'

* Modified property `authMode` (string)
    > The association control method for the SSID ('open', 'open-enhanced', 'psk', 'open-with-radius', '8021x-meraki', '8021x-radius', '8021x-google', '8021x-localradius', 'ipsk-with-radius' or 'ipsk-without-radius')

* Modified property `ipAssignmentMode` (string)
    > The client IP assignment mode ('NAT mode', 'Bridge mode', 'Layer 3 roaming', 'Ethernet over GRE', 'Layer 3 roaming with a concentrator' or 'VPN')

* Modified property `gre` (object)
    > Ethernet over GRE settings

    * Modified property `concentrator` (object)
        > The EoGRE concentrator's settings


##### `GET` /organizations

> List the organizations that the user has privileges on

###### Response:

Modified response: **200 OK**
> Successful operation

* Modified content type: `application/json`
    * Modified items (array):

        * Added property `cloud` (object)
            > Data for this organization


##### `GET` /organizations/{organizationId}

> Return an organization

###### Response:

Modified response: **200 OK**
> Successful operation

* Modified content type: `application/json`
    * Added property `cloud` (object)
        > Data for this organization


##### `POST` /organizations/{organizationId}/alerts/profiles

> Create an organization-wide alert configuration

###### Request:

Modified content type: `application/json`

* Modified property `type` (string)
    > The alert type


##### `PUT` /organizations/{organizationId}/alerts/profiles/{alertConfigId}

> Update an organization-wide alert config

###### Request:

Modified content type: `application/json`

* Modified property `type` (string)
    > The alert type


##### `GET` /organizations/{organizationId}/clients/bandwidthUsageHistory

> Return data usage (in megabits per second) over time for all clients in the given organization within a given time range.

###### Response:

Modified response: **200 OK**
> Successful operation

* Modified content type: `application/json`
    * Modified items (array):

        * Added property `ts` (string)
            > Timestamp for the bandwidth usage snapshot.

        * Added property `total` (integer)
            > Total bandwidth usage, in mbps.

        * Added property `upstream` (integer)
            > Uploaded data, in mbps.

        * Added property `downstream` (integer)
            > Downloaded data, in mbps.


##### `GET` /organizations/{organizationId}/loginSecurity

> Returns the login security settings for an organization.

###### Response:

Modified response: **200 OK**
> Successful operation

* Modified content type: `application/json`
    * Added property `enforceDifferentPasswords` (boolean)
        > Whether the org is forced to choose a password that is distinct from previous passwords

    * Added property `enforceStrongPasswords` (boolean)
        > Whether the org must choose a strong password

    * Added property `accountLockoutAttempts` (integer)
        > Number of failed login attempts before the admin is locked out of the org

    * Added property `enforceLoginIpRanges` (boolean)
        > Whether the org restricts access to Dashboard by IP addresses

    * Added property `apiAuthentication` (object)
        > Collection of properties related to API authentication

    * Added property `enforceTwoFactorAuth` (boolean)
        > Whether the org uses two factor authentication

    * Added property `enforcePasswordExpiration` (boolean)
        > Whether password expiration should be enforced

    * Added property `passwordExpirationDays` (integer)
        > How long the password is valid, in days

    * Added property `numDifferentPasswords` (integer)
        > How many recent password that cannot be reused

    * Added property `enforceIdleTimeout` (boolean)
        > Whether the org admin should be logged out after idling for the given idle timeout duration

    * Added property `idleTimeoutMinutes` (integer)
        > Length of time, in minutes, an org admin is idle before being logged out


##### `PUT` /organizations/{organizationId}/loginSecurity

> Update the login security settings for an organization

###### Request:

Modified content type: `application/json`

* Added property `apiAuthentication` (object)
    > Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.


