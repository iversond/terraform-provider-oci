// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_streaming "github.com/oracle/oci-go-sdk/streaming"
)

func StreamingStreamResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createStreamingStream,
		Read:     readStreamingStream,
		Update:   updateStreamingStream,
		Delete:   deleteStreamingStream,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"partitions": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"retention_in_hours": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"messages_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return CreateResource(d, sync)
}

func readStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return ReadResource(sync)
}

func updateStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient

	return UpdateResource(d, sync)
}

func deleteStreamingStream(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).streamAdminClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type StreamingStreamResourceCrud struct {
	BaseCrud
	Client                 *oci_streaming.StreamAdminClient
	Res                    *oci_streaming.Stream
	DisableNotFoundRetries bool
}

func (s *StreamingStreamResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StreamingStreamResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateCreating),
	}
}

func (s *StreamingStreamResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateActive),
	}
}

func (s *StreamingStreamResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateDeleting),
	}
}

func (s *StreamingStreamResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_streaming.StreamLifecycleStateDeleted),
	}
}

func (s *StreamingStreamResourceCrud) Create() error {
	request := oci_streaming.CreateStreamRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if partitions, ok := s.D.GetOkExists("partitions"); ok {
		tmp := partitions.(int)
		request.Partitions = &tmp
	}

	if retentionInHours, ok := s.D.GetOkExists("retention_in_hours"); ok {
		tmp := retentionInHours.(int)
		request.RetentionInHours = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.CreateStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stream
	return nil
}

func (s *StreamingStreamResourceCrud) Get() error {
	request := oci_streaming.GetStreamRequest{}

	tmp := s.D.Id()
	request.StreamId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.GetStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stream
	return nil
}

func (s *StreamingStreamResourceCrud) Update() error {
	request := oci_streaming.UpdateStreamRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.StreamId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	response, err := s.Client.UpdateStream(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Stream
	return nil
}

func (s *StreamingStreamResourceCrud) Delete() error {
	request := oci_streaming.DeleteStreamRequest{}

	tmp := s.D.Id()
	request.StreamId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "streaming")

	_, err := s.Client.DeleteStream(context.Background(), request)
	return err
}

func (s *StreamingStreamResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.MessagesEndpoint != nil {
		s.D.Set("messages_endpoint", *s.Res.MessagesEndpoint)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Partitions != nil {
		s.D.Set("partitions", *s.Res.Partitions)
	}

	if s.Res.RetentionInHours != nil {
		s.D.Set("retention_in_hours", *s.Res.RetentionInHours)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
