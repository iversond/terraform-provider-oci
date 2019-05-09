// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_ons "github.com/oracle/oci-go-sdk/ons"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SubscriptionRequiredOnlyResource = SubscriptionResourceDependencies +
		generateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", Required, Create, subscriptionRepresentation)

	SubscriptionResourceConfig = SubscriptionResourceDependencies +
		generateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", Optional, Create, subscriptionRepresentation)

	subscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": Representation{repType: Required, create: `${oci_ons_subscription.test_subscription.id}`},
	}

	subscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"topic_id":       Representation{repType: Optional, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"filter":         RepresentationGroup{Required, subscriptionDataSourceFilterRepresentation}}
	subscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_ons_subscription.test_subscription.id}`}},
	}

	subscriptionRepresentation = map[string]interface{}{
		"compartment_id":  Representation{repType: Required, create: `${var.compartment_id}`},
		"endpoint":        Representation{repType: Required, create: `john.smith@example.com`},
		"protocol":        Representation{repType: Required, create: `EMAIL`},
		"topic_id":        Representation{repType: Required, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"defined_tags":    Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":   Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"delivery_policy": Representation{repType: Optional, update: `{\"backoffRetryPolicy\":{\"initialDelayInFailureRetry\":60000,\"maxRetryDuration\":7000000,\"policyType\":\"EXPONENTIAL\"}, \"maxReceiveRatePerSecond\" : 0}`},
	}

	SubscriptionResourceDependencies = NotificationTopicResourceDependencies +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, getTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, charsetWithoutDigits, "tsubscription"))
)

func TestOnsSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnsSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_ons_subscription.test_subscription"
	datasourceName := "data.oci_ons_subscriptions.test_subscriptions"
	singularDatasourceName := "data.oci_ons_subscription.test_subscription"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOnsSubscriptionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", Required, Create, subscriptionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "endpoint", "john.smith@example.com"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "EMAIL"),
					resource.TestCheckResourceAttrSet(resourceName, "topic_id"),
					resource.TestCheckResourceAttrSet(resourceName, "delivery_policy"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", Optional, Create, subscriptionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "endpoint", "john.smith@example.com"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "EMAIL"),
					resource.TestCheckResourceAttr(resourceName, "state", "PENDING"),
					resource.TestCheckResourceAttrSet(resourceName, "topic_id"),
					resource.TestCheckResourceAttrSet(resourceName, "delivery_policy"),
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", Optional, Update, subscriptionRepresentation),
				ExpectError: regexp.MustCompile("Subscription(.*) is not active."),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_ons_subscriptions", "test_subscriptions", Optional, Create, subscriptionDataSourceRepresentation) +
					compartmentIdVariableStr + SubscriptionResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", Optional, Create, subscriptionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "topic_id"),

					resource.TestCheckResourceAttr(datasourceName, "subscriptions.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.delivery_policy.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.endpoint", "john.smith@example.com"),
					resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.etag"),
					resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.protocol", "EMAIL"),
					resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.state"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_ons_subscription", "test_subscription", Required, Create, subscriptionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + SubscriptionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "endpoint", "john.smith@example.com"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "EMAIL"),
					resource.TestCheckResourceAttr(singularDatasourceName, "state", "PENDING"),
				),
			},
		},
	})
}
func testAccCheckOnsSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).notificationDataPlaneClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ons_subscription" {
			noResourceFound = false
			request := oci_ons.GetSubscriptionRequest{}

			tmp := rs.Primary.ID
			request.SubscriptionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ons")

			response, err := client.GetSubscription(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ons.SubscriptionLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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

func init() {
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	resource.AddTestSweepers("OnsSubscription", &resource.Sweeper{
		Name:         "OnsSubscription",
		Dependencies: DependencyGraph["subscription"],
		F:            sweepOnsSubscriptionResource,
	})
}

func sweepOnsSubscriptionResource(compartment string) error {
	notificationDataPlaneClient := GetTestClients(&schema.ResourceData{}).notificationDataPlaneClient
	subscriptionIds, err := getSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, subscriptionId := range subscriptionIds {
		if ok := SweeperDefaultResourceId[subscriptionId]; !ok {
			deleteSubscriptionRequest := oci_ons.DeleteSubscriptionRequest{}

			deleteSubscriptionRequest.SubscriptionId = &subscriptionId

			deleteSubscriptionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ons")
			_, error := notificationDataPlaneClient.DeleteSubscription(context.Background(), deleteSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting Subscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", subscriptionId, error)
				continue
			}
			waitTillCondition(testAccProvider, &subscriptionId, subscriptionSweepWaitCondition, time.Duration(3*time.Minute),
				subscriptionSweepResponseFetchOperation, "ons", true)
		}
	}
	return nil
}

func getSubscriptionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "SubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	notificationDataPlaneClient := GetTestClients(&schema.ResourceData{}).notificationDataPlaneClient

	listSubscriptionsRequest := oci_ons.ListSubscriptionsRequest{}
	listSubscriptionsRequest.CompartmentId = &compartmentId
	listSubscriptionsResponse, err := notificationDataPlaneClient.ListSubscriptions(context.Background(), listSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Subscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subscription := range listSubscriptionsResponse.Items {
		id := *subscription.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "SubscriptionId", id)
	}
	return resourceIds, nil
}

func subscriptionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subscriptionResponse, ok := response.Response.(oci_ons.GetSubscriptionResponse); ok {
		return subscriptionResponse.LifecycleState != oci_ons.SubscriptionLifecycleStateDeleted
	}
	return false
}

func subscriptionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.notificationDataPlaneClient.GetSubscription(context.Background(), oci_ons.GetSubscriptionRequest{
		SubscriptionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
