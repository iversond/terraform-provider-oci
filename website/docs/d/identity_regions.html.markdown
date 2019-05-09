---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_identity_regions"
sidebar_current: "docs-oci-datasource-identity-regions"
description: |-
  Provides the list of Regions in Oracle Cloud Infrastructure Identity service
---

# Data Source: oci_identity_regions
This data source provides the list of Regions in Oracle Cloud Infrastructure Identity service.

Lists all the regions offered by Oracle Cloud Infrastructure.

## Example Usage

```hcl
data "oci_identity_regions" "test_regions" {
}
```

## Attributes Reference

The following attributes are exported:

* `regions` - The list of regions.

### Region Reference

The following attributes are exported:

* `key` - The key of the region.

	Allowed values are:
	* `PHX`
	* `IAD`
	* `FRA`
	* `LHR` 
* `name` - The name of the region.

	Allowed values are:
	* `us-phoenix-1`
	* `us-ashburn-1`
	* `eu-frankfurt-1`
	* `uk-london-1` 

