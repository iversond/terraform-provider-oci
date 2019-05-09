// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	PreauthenticatedRequestRequiredOnlyResource = PreauthenticatedRequestResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, preauthenticatedRequestRepresentation)

	PreauthenticatedRequestResourceConfig = PreauthenticatedRequestResourceDependencies +
		generateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation)

	preauthenticatedRequestSingularDataSourceRepresentation = map[string]interface{}{
		"bucket":    Representation{repType: Required, create: testBucketName},
		"namespace": Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"par_id":    Representation{repType: Required, create: `${oci_objectstorage_preauthrequest.test_preauthenticated_request.par_id}`},
	}

	preauthenticatedRequestDataSourceRepresentation = map[string]interface{}{
		"bucket":             Representation{repType: Required, create: testBucketName},
		"namespace":          Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"object_name_prefix": Representation{repType: Optional, create: `my-test-object`},
		"filter":             RepresentationGroup{Required, preauthenticatedRequestDataSourceFilterRepresentation}}
	preauthenticatedRequestDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_objectstorage_preauthrequest.test_preauthenticated_request.par_id}`}},
	}

	preauthenticatedRequestRepresentation = map[string]interface{}{
		"access_type":  Representation{repType: Required, create: `AnyObjectWrite`, update: `ObjectRead`},
		"bucket":       Representation{repType: Required, create: testBucketName},
		"name":         Representation{repType: Required, create: `-tf-par`},
		"namespace":    Representation{repType: Required, create: `${oci_objectstorage_bucket.test_bucket.namespace}`},
		"time_expires": Representation{repType: Required, create: `2020-01-01T00:00:00Z`},
		"object":       Representation{repType: Optional, create: `my-test-object-1`},
	}

	PreauthenticatedRequestResourceDependencies = ObjectRequiredOnlyResource
)

func TestObjectStoragePreauthenticatedRequestResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestObjectStoragePreauthenticatedRequestResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_objectstorage_preauthrequest.test_preauthenticated_request"
	datasourceName := "data.oci_objectstorage_preauthrequests.test_preauthenticated_requests"
	singularDatasourceName := "data.oci_objectstorage_preauthrequest.test_preauthenticated_request"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckObjectStoragePreauthenticatedRequestDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, preauthenticatedRequestRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "AnyObjectWrite"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "time_expires", "2020-01-01T00:00:00Z"),
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(resourceName, "access_uri"),
					resource.TestCheckResourceAttr(resourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttrSet(resourceName, "namespace"),
					resource.TestCheckResourceAttr(resourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "time_expires", "2020-01-01T00:00:00Z"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequests", "test_preauthenticated_requests", Optional, Update, preauthenticatedRequestDataSourceRepresentation) +
					compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttrSet(datasourceName, "namespace"),
					resource.TestCheckResourceAttr(datasourceName, "object_name_prefix", "my-test-object"),

					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.id"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.name", "-tf-par"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.object", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(datasourceName, "preauthenticated_requests.0.time_created"),
					resource.TestCheckResourceAttr(datasourceName, "preauthenticated_requests.0.time_expires", "2020-01-01 00:00:00 +0000 UTC"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Required, Create, preauthenticatedRequestSingularDataSourceRepresentation) +
					compartmentIdVariableStr + PreauthenticatedRequestResourceDependencies +
					generateResourceFromRepresentationMap("oci_objectstorage_preauthrequest", "test_preauthenticated_request", Optional, Update, preauthenticatedRequestRepresentation),

				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "bucket", testBucketName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "par_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "access_type", "ObjectRead"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "-tf-par"),
					resource.TestCheckResourceAttr(singularDatasourceName, "object", "my-test-object-1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttr(singularDatasourceName, "time_expires", "2020-01-01 00:00:00 +0000 UTC"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + PreauthenticatedRequestResourceConfig,
			},
			//verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"access_uri",
					"time_expires",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckObjectStoragePreauthenticatedRequestDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).objectStorageClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_objectstorage_preauthrequest" {
			noResourceFound = false
			request := oci_object_storage.GetPreauthenticatedRequestRequest{}

			if value, ok := rs.Primary.Attributes["bucket"]; ok {
				request.BucketName = &value
			}

			if value, ok := rs.Primary.Attributes["namespace"]; ok {
				request.NamespaceName = &value
			}

			bucket, namespace, parId, er := parsePreauthenticatedRequestCompositeId(rs.Primary.ID)
			if er == nil {
				request.BucketName = &bucket
				request.NamespaceName = &namespace
				request.ParId = &parId
			} else {
				log.Printf("[WARN] Get() unable to parse current ID: %s", rs.Primary.ID)
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "object_storage")

			_, err := client.GetPreauthenticatedRequest(context.Background(), request)

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
