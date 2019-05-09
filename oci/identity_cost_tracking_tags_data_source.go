// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityCostTrackingTagsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityCostTrackingTags,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_cost_tracking": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_retired": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_namespace_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag_namespace_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readIdentityCostTrackingTags(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityCostTrackingTagsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

type IdentityCostTrackingTagsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListCostTrackingTagsResponse
}

func (s *IdentityCostTrackingTagsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityCostTrackingTagsDataSourceCrud) Get() error {
	request := oci_identity.ListCostTrackingTagsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListCostTrackingTags(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	// TODO- remove this custom code handling paging once service fixes Opc-Next-Page in spec
	if s.Res != nil && s.Res.RawResponse != nil {
		rawResponse := s.Res.RawResponse
		nextPage := rawResponse.Header.Get(OpcNextPageHeader)
		request.Page = &nextPage

		for request.Page != nil && *request.Page != "" {
			listResponse, err := s.Client.ListCostTrackingTags(context.Background(), request)
			if err != nil {
				return err
			}

			s.Res.Items = append(s.Res.Items, listResponse.Items...)
			if listResponse.RawResponse != nil {
				nextPage = listResponse.RawResponse.Header.Get(OpcNextPageHeader)
				request.Page = &nextPage
			} else {
				request.Page = nil
			}

		}
	}

	return nil
}

func (s *IdentityCostTrackingTagsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		costTrackingTag := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			costTrackingTag["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			costTrackingTag["description"] = *r.Description
		}

		costTrackingTag["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			costTrackingTag["id"] = *r.Id
		}

		if r.IsCostTracking != nil {
			costTrackingTag["is_cost_tracking"] = *r.IsCostTracking
		}

		if r.IsRetired != nil {
			costTrackingTag["is_retired"] = *r.IsRetired
		}

		if r.Name != nil {
			costTrackingTag["name"] = *r.Name
		}

		if r.TagNamespaceId != nil {
			costTrackingTag["tag_namespace_id"] = *r.TagNamespaceId
		}

		if r.TagNamespaceName != nil {
			costTrackingTag["tag_namespace_name"] = *r.TagNamespaceName
		}

		if r.TimeCreated != nil {
			costTrackingTag["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, costTrackingTag)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityCostTrackingTagsDataSource().Schema["tags"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("tags", resources); err != nil {
		return err
	}

	return nil
}
