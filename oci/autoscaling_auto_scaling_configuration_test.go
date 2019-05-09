// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	oci_autoscaling "github.com/oracle/oci-go-sdk/autoscaling"
	"github.com/oracle/oci-go-sdk/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AutoScalingConfigurationRequiredOnlyResource = AutoScalingConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", Required, Create, autoScalingConfigurationRepresentation)

	AutoScalingConfigurationResourceConfig = AutoScalingConfigurationResourceDependencies +
		generateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Update, autoScalingConfigurationRepresentation)

	autoScalingConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"auto_scaling_configuration_id": Representation{repType: Required, create: `${oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration.id}`},
	}

	autoScalingConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"filter":         RepresentationGroup{Required, autoScalingConfigurationDataSourceFilterRepresentation}}
	autoScalingConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration.id}`}},
	}

	autoScalingConfigurationRepresentation = map[string]interface{}{
		"auto_scaling_resources": RepresentationGroup{Required, autoScalingConfigurationAutoScalingResourcesRepresentation},
		"compartment_id":         Representation{repType: Required, create: `${var.compartment_id}`},
		"policies":               RepresentationGroup{Required, autoScalingConfigurationPoliciesRepresentation},
		"cool_down_in_seconds":   Representation{repType: Optional, create: `300`, update: `400`},
		"defined_tags":           Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":           Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":          Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_enabled":             Representation{repType: Optional, create: `false`, update: `true`},
	}
	autoScalingConfigurationAutoScalingResourcesRepresentation = map[string]interface{}{
		"id":   Representation{repType: Required, create: `${oci_core_instance_pool.test_instance_pool.id}`},
		"type": Representation{repType: Required, create: `instancePool`},
	}
	autoScalingConfigurationPoliciesRepresentation = map[string]interface{}{
		"capacity":     RepresentationGroup{Required, autoScalingConfigurationPoliciesCapacityRepresentation},
		"policy_type":  Representation{repType: Required, create: `threshold`, update: `threshold`},
		"rules":        []RepresentationGroup{{Required, autoScalingConfigurationPoliciesScaleOutRuleRepresentation}, {Required, autoScalingConfigurationPoliciesScaleInRuleRepresentation}},
		"display_name": Representation{repType: Optional, create: `displayName`, update: `displayName2`},
	}
	autoScalingConfigurationPoliciesCapacityRepresentation = map[string]interface{}{
		"initial": Representation{repType: Required, create: `2`, update: `4`},
		"max":     Representation{repType: Required, create: `3`, update: `5`},
		"min":     Representation{repType: Required, create: `2`, update: `3`},
	}
	autoScalingConfigurationPoliciesScaleOutRuleRepresentation = map[string]interface{}{
		"action":       RepresentationGroup{Required, autoScalingConfigurationPoliciesScaleOutRuleActionRepresentation},
		"metric":       RepresentationGroup{Required, autoScalingConfigurationPoliciesScaleOutRuleMetricRepresentation},
		"display_name": Representation{repType: Required, create: `scale out rule`, update: `scale out rule - updated`},
	}
	autoScalingConfigurationPoliciesScaleOutRuleActionRepresentation = map[string]interface{}{
		"type":  Representation{repType: Required, create: `CHANGE_COUNT_BY`, update: `CHANGE_COUNT_BY`},
		"value": Representation{repType: Required, create: `1`, update: `2`},
	}
	autoScalingConfigurationPoliciesScaleOutRuleMetricRepresentation = map[string]interface{}{
		"metric_type": Representation{repType: Required, create: `CPU_UTILIZATION`, update: `CPU_UTILIZATION`},
		"threshold":   RepresentationGroup{Required, autoScalingConfigurationPoliciesScaleOutRuleMetricThresholdRepresentation},
	}
	autoScalingConfigurationPoliciesScaleOutRuleMetricThresholdRepresentation = map[string]interface{}{
		"operator": Representation{repType: Required, create: `GT`, update: `GT`},
		"value":    Representation{repType: Required, create: `1`, update: `3`},
	}
	autoScalingConfigurationPoliciesScaleInRuleRepresentation = map[string]interface{}{
		"action":       RepresentationGroup{Required, autoScalingConfigurationPoliciesScaleInRuleActionRepresentation},
		"metric":       RepresentationGroup{Required, autoScalingConfigurationPoliciesScaleInRuleMetricRepresentation},
		"display_name": Representation{repType: Required, create: `scale in rule`, update: `scale in rule - updated`},
	}
	autoScalingConfigurationPoliciesScaleInRuleActionRepresentation = map[string]interface{}{
		"type":  Representation{repType: Required, create: `CHANGE_COUNT_BY`, update: `CHANGE_COUNT_BY`},
		"value": Representation{repType: Required, create: `-1`, update: `-3`},
	}
	autoScalingConfigurationPoliciesScaleInRuleMetricRepresentation = map[string]interface{}{
		"metric_type": Representation{repType: Required, create: `CPU_UTILIZATION`, update: `CPU_UTILIZATION`},
		"threshold":   RepresentationGroup{Required, autoScalingConfigurationPoliciesScaleInRuleMetricThresholdRepresentation},
	}
	autoScalingConfigurationPoliciesScaleInRuleMetricThresholdRepresentation = map[string]interface{}{
		"operator": Representation{repType: Required, create: `LT`, update: `LT`},
		"value":    Representation{repType: Required, create: `1`, update: `3`},
	}

	AutoScalingConfigurationResourceDependencies = InstancePoolResourceDependenciesWithoutSecondaryVnic +
		generateResourceFromRepresentationMap("oci_core_instance_pool", "test_instance_pool", Required, Create, instancePoolRepresentation)
)

func TestAutoscalingAutoScalingConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAutoscalingAutoScalingConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"
	datasourceName := "data.oci_autoscaling_auto_scaling_configurations.test_auto_scaling_configurations"
	singularDatasourceName := "data.oci_autoscaling_auto_scaling_configuration.test_auto_scaling_configuration"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckAutoscalingAutoScalingConfigurationDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", Required, Create, autoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
					resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "3"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "threshold"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "1",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "GT",
						"metric.0.threshold.0.value":    "1",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "-1",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "LT",
						"metric.0.threshold.0.value":    "1",
					},
						[]string{}),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Create, autoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
					resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "300"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "3"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "threshold"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "1",
						"display_name":                  "scale out rule",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "GT",
						"metric.0.threshold.0.value":    "1",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "-1",
						"display_name":                  "scale in rule",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "LT",
						"metric.0.threshold.0.value":    "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Update, autoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "auto_scaling_resources.0.id"),
					resource.TestCheckResourceAttr(resourceName, "auto_scaling_resources.0.type", "instancePool"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cool_down_in_seconds", "400"),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "policies.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "policies.0.id"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.initial", "4"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.max", "5"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.capacity.0.min", "3"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.policy_type", "threshold"),
					resource.TestCheckResourceAttr(resourceName, "policies.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "2",
						"display_name":                  "scale out rule - updated",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "GT",
						"metric.0.threshold.0.value":    "3",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "-3",
						"display_name":                  "scale in rule - updated",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "LT",
						"metric.0.threshold.0.value":    "3",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "policies.0.time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId == resId2 {
							return fmt.Errorf("Resource updated when it was supposed to be recreated.")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configurations", "test_auto_scaling_configurations", Optional, Update, autoScalingConfigurationDataSourceRepresentation) +
					compartmentIdVariableStr + AutoScalingConfigurationResourceDependencies +
					generateResourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", Optional, Update, autoScalingConfigurationRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

					resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.auto_scaling_resources.0.type", "instancePool"),
					resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.cool_down_in_seconds", "400"),
					resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.display_name", "displayName2"),
					resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "auto_scaling_configurations.0.is_enabled", "true"),
					resource.TestCheckResourceAttrSet(datasourceName, "auto_scaling_configurations.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_autoscaling_auto_scaling_configuration", "test_auto_scaling_configuration", Required, Create, autoScalingConfigurationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutoScalingConfigurationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_configuration_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.#", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_scaling_resources.0.id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "auto_scaling_resources.0.type", "instancePool"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "cool_down_in_seconds", "400"),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_enabled", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.initial", "4"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.max", "5"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.capacity.0.min", "3"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.policy_type", "threshold"),
					resource.TestCheckResourceAttr(singularDatasourceName, "policies.0.rules.#", "2"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "2",
						"display_name":                  "scale out rule - updated",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "GT",
						"metric.0.threshold.0.value":    "3",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "policies.0.rules", map[string]string{
						"action.#":                      "1",
						"action.0.type":                 "CHANGE_COUNT_BY",
						"action.0.value":                "-3",
						"display_name":                  "scale in rule - updated",
						"metric.#":                      "1",
						"metric.0.metric_type":          "CPU_UTILIZATION",
						"metric.0.threshold.#":          "1",
						"metric.0.threshold.0.operator": "LT",
						"metric.0.threshold.0.value":    "3",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "policies.0.time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + AutoScalingConfigurationResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckAutoscalingAutoScalingConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).autoScalingClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_autoscaling_auto_scaling_configuration" {
			noResourceFound = false
			request := oci_autoscaling.GetAutoScalingConfigurationRequest{}

			tmp := rs.Primary.ID
			request.AutoScalingConfigurationId = &tmp

			_, err := client.GetAutoScalingConfiguration(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}
