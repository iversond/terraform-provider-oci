// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_containerengine "github.com/oracle/oci-go-sdk/containerengine"
)

func ContainerengineNodePoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["node_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ContainerengineNodePoolResource(), fieldMap, readSingularContainerengineNodePool)
}

func readSingularContainerengineNodePool(d *schema.ResourceData, m interface{}) error {
	sync := &ContainerengineNodePoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).containerEngineClient

	return ReadResource(sync)
}

type ContainerengineNodePoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_containerengine.ContainerEngineClient
	Res    *oci_containerengine.GetNodePoolResponse
}

func (s *ContainerengineNodePoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ContainerengineNodePoolDataSourceCrud) Get() error {
	request := oci_containerengine.GetNodePoolRequest{}

	if nodePoolId, ok := s.D.GetOkExists("node_pool_id"); ok {
		tmp := nodePoolId.(string)
		request.NodePoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "containerengine")

	response, err := s.Client.GetNodePool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ContainerengineNodePoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ClusterId != nil {
		s.D.Set("cluster_id", *s.Res.ClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	initialNodeLabels := []interface{}{}
	for _, item := range s.Res.InitialNodeLabels {
		initialNodeLabels = append(initialNodeLabels, KeyValueToMap(item))
	}
	s.D.Set("initial_node_labels", initialNodeLabels)

	if s.Res.KubernetesVersion != nil {
		s.D.Set("kubernetes_version", *s.Res.KubernetesVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NodeImageId != nil {
		s.D.Set("node_image_id", *s.Res.NodeImageId)
	}

	if s.Res.NodeImageName != nil {
		s.D.Set("node_image_name", *s.Res.NodeImageName)
	}

	s.D.Set("node_metadata", s.Res.NodeMetadata)

	if s.Res.NodeShape != nil {
		s.D.Set("node_shape", *s.Res.NodeShape)
	}

	nodes := []interface{}{}
	for _, item := range s.Res.Nodes {
		nodes = append(nodes, NodeToMap(item))
	}
	s.D.Set("nodes", nodes)

	if s.Res.QuantityPerSubnet != nil {
		s.D.Set("quantity_per_subnet", *s.Res.QuantityPerSubnet)
	}

	if s.Res.SshPublicKey != nil {
		s.D.Set("ssh_public_key", *s.Res.SshPublicKey)
	}

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	return nil
}
