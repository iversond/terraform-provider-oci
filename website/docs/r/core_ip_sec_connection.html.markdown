---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipsec"
sidebar_current: "docs-oci-resource-core-ipsec"
description: |-
  Provides the Ip Sec Connection resource in Oracle Cloud Infrastructure Core service
---

# oci_core_ipsec
This resource provides the Ip Sec Connection resource in Oracle Cloud Infrastructure Core service.

Creates a new IPSec connection between the specified DRG and CPE. For more information, see
[IPSec VPNs](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPsec.htm).

In the request, you must include at least one static route to the CPE object (you're allowed a maximum
of 10). For example: 10.0.8.0/16.

For the purposes of access control, you must provide the OCID of the compartment where you want the
IPSec connection to reside. Notice that the IPSec connection doesn't have to be in the same compartment
as the DRG, CPE, or other Networking Service components. If you're not sure which compartment to
use, put the IPSec connection in the same compartment as the DRG. For more information about
compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the IPSec connection, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

After creating the IPSec connection, you need to configure your on-premises router
with tunnel-specific information returned by
[GetIPSecConnectionDeviceConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnectionDeviceConfig/GetIPSecConnectionDeviceConfig).
For each tunnel, that operation gives you the IP address of Oracle's VPN headend and the shared secret
(that is, the pre-shared key). For more information, see
[Configuring Your On-Premises Router for an IPSec VPN](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/configuringCPE.htm).

To get the status of the tunnels (whether they're up or down), use
[GetIPSecConnectionDeviceStatus](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnectionDeviceStatus/GetIPSecConnectionDeviceStatus).


## Example Usage

```hcl
resource "oci_core_ipsec" "test_ip_sec_connection" {
	#Required
	compartment_id = "${var.compartment_id}"
	cpe_id = "${oci_core_cpe.test_cpe.id}"
	drg_id = "${oci_core_drg.test_drg.id}"
	static_routes = "${var.ip_sec_connection_static_routes}"

	#Optional
	cpe_local_identifier = "${var.ip_sec_connection_cpe_local_identifier}"
	cpe_local_identifier_type = "${var.ip_sec_connection_cpe_local_identifier_type}"
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.ip_sec_connection_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the IPSec connection.
* `cpe_id` - (Required) The OCID of the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/) object.
* `cpe_local_identifier` - (Optional) (Updatable) Your identifier for your CPE device. Can be either an IP address or a hostname (specifically, the fully qualified domain name (FQDN)). The type of identifier you provide here must correspond to the value for `cpeLocalIdentifierType`.

	If you don't provide a value, the `ipAddress` attribute for the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/) object specified by `cpeId` is used as the `cpeLocalIdentifier`.

	Example IP address: `10.0.3.3`

	Example hostname: `cpe.example.com` 
* `cpe_local_identifier_type` - (Optional) (Updatable) The type of identifier for your CPE device. The value you provide here must correspond to the value for `cpeLocalIdentifier`. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `drg_id` - (Required) The OCID of the DRG.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `static_routes` - (Required) (Updatable) Static routes to the CPE. At least one route must be included. A static route's CIDR must not be a multicast address or class E address.  Example: `10.0.1.0/24` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the IPSec connection.
* `cpe_id` - The OCID of the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/) object.
* `cpe_local_identifier` - Your identifier for your CPE device. Can be either an IP address or a hostname (specifically, the fully qualified domain name (FQDN)). The type of identifier here must correspond to the value for `cpeLocalIdentifierType`.

	If you don't provide a value when creating the IPSec connection, the `ipAddress` attribute for the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/) object specified by `cpeId` is used as the `cpeLocalIdentifier`.

	Example IP address: `10.0.3.3`

	Example hostname: `cpe.example.com` 
* `cpe_local_identifier_type` - The type of identifier for your CPE device. The value here must correspond to the value for `cpeLocalIdentifier`. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `drg_id` - The OCID of the DRG.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The IPSec connection's Oracle ID (OCID).
* `state` - The IPSec connection's current state.
* `static_routes` - Static routes to the CPE. At least one route must be included. The CIDR must not be a multicast address or class E address.  Example: `10.0.1.0/24` 
* `time_created` - The date and time the IPSec connection was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

IpSecConnections can be imported using the `id`, e.g.

```
$ terraform import oci_core_ipsec.test_ip_sec_connection "id"
```

