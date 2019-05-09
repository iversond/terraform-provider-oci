---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_metrics"
sidebar_current: "docs-oci-datasource-monitoring-metrics"
description: |-
  Provides the list of Metrics in Oracle Cloud Infrastructure Monitoring service
---

# Data Source: oci_monitoring_metrics
This data source provides the list of Metrics in Oracle Cloud Infrastructure Monitoring service.

Returns metric definitions that match the criteria specified in the request. Compartment OCID required.
For information about metrics, see [Metrics Overview](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#MetricsOverview).


## Example Usage

```hcl
data "oci_monitoring_metrics" "test_metrics" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	compartment_id_in_subtree = "${var.metric_compartment_id_in_subtree}"
	dimension_filters = "${var.metric_dimension_filters}"
	group_by = "${var.metric_group_by}"
	name = "${var.metric_name}"
	namespace = "${var.metric_namespace}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources monitored by the metric that you are searching for. Use tenancyId to search in the root compartment. 
* `compartment_id_in_subtree` - (Optional) When true, returns resources from all compartments and subcompartments. The parameter can only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, returns resources from only the compartment specified in compartmentId. Default is false. 
* `dimension_filters` - (Optional) Qualifiers that you want to use when searching for metric definitions. Available dimensions vary by metric namespace. Each dimension takes the form of a key-value pair.  Example: { "resourceId": "<var>&lt;instance_OCID&gt;</var>" } 
* `group_by` - (Optional) Group metrics by these fields in the response. For example, to list all metric namespaces available in a compartment, groupBy the "namespace" field.

	Example - group by namespace and resource: `[ "namespace", "resourceId" ]` 
* `name` - (Optional) The metric name to use when searching for metric definitions.  Example: `CpuUtilization` 
* `namespace` - (Optional) The source service or application to use when searching for metric definitions.  Example: `oci_computeagent` 


## Attributes Reference

The following attributes are exported:

* `metrics` - The list of metrics.

### Metric Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources monitored by the metric. 
* `dimensions` - Qualifiers provided in a metric definition. Available dimensions vary by metric namespace. Each dimension takes the form of a key-value pair.  Example: `"resourceId": "ocid1.instance.region1.phx.exampleuniqueID"` 
* `name` - The name of the metric.  Example: `CpuUtilization` 
* `namespace` - The source service or application emitting the metric.  Example: `oci_computeagent` 

