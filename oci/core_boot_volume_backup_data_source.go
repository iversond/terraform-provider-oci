// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func CoreBootVolumeBackupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["boot_volume_backup_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CoreBootVolumeBackupResource(), fieldMap, readSingularCoreBootVolumeBackup)
}

func readSingularCoreBootVolumeBackup(d *schema.ResourceData, m interface{}) error {
	sync := &CoreBootVolumeBackupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return ReadResource(sync)
}

type CoreBootVolumeBackupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.GetBootVolumeBackupResponse
}

func (s *CoreBootVolumeBackupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreBootVolumeBackupDataSourceCrud) Get() error {
	request := oci_core.GetBootVolumeBackupRequest{}

	if bootVolumeBackupId, ok := s.D.GetOkExists("boot_volume_backup_id"); ok {
		tmp := bootVolumeBackupId.(string)
		request.BootVolumeBackupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetBootVolumeBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreBootVolumeBackupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BootVolumeId != nil {
		s.D.Set("boot_volume_id", *s.Res.BootVolumeId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExpirationTime != nil {
		s.D.Set("expiration_time", s.Res.ExpirationTime.String())
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeRequestReceived != nil {
		s.D.Set("time_request_received", s.Res.TimeRequestReceived.String())
	}

	s.D.Set("type", s.Res.Type)

	if s.Res.UniqueSizeInGBs != nil {
		s.D.Set("unique_size_in_gbs", strconv.FormatInt(*s.Res.UniqueSizeInGBs, 10))
	}

	return nil
}
