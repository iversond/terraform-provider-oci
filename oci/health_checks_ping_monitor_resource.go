// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	oci_health_checks "github.com/oracle/oci-go-sdk/healthchecks"
)

func HealthChecksPingMonitorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createHealthChecksPingMonitor,
		Read:     readHealthChecksPingMonitor,
		Update:   updateHealthChecksPingMonitor,
		Delete:   deleteHealthChecksPingMonitor,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"interval_in_seconds": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vantage_point_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"results_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return CreateResource(d, sync)
}

func readHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return ReadResource(sync)
}

func updateHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return UpdateResource(d, sync)
}

func deleteHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type HealthChecksPingMonitorResourceCrud struct {
	BaseCrud
	Client                 *oci_health_checks.HealthChecksClient
	Res                    *oci_health_checks.PingMonitor
	DisableNotFoundRetries bool
}

func (s *HealthChecksPingMonitorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *HealthChecksPingMonitorResourceCrud) Create() error {
	request := oci_health_checks.CreatePingMonitorRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if intervalInSeconds, ok := s.D.GetOkExists("interval_in_seconds"); ok {
		tmp := intervalInSeconds.(int)
		request.IntervalInSeconds = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_health_checks.CreatePingMonitorDetailsProtocolEnum(protocol.(string))
	}

	request.Targets = []string{}
	if targets, ok := s.D.GetOkExists("targets"); ok {
		interfaces := targets.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Targets = tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	request.VantagePointNames = []string{}
	if vantagePointNames, ok := s.D.GetOkExists("vantage_point_names"); ok {
		interfaces := vantagePointNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.VantagePointNames = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.CreatePingMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PingMonitor
	return nil
}

func (s *HealthChecksPingMonitorResourceCrud) Get() error {
	request := oci_health_checks.GetPingMonitorRequest{}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.GetPingMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PingMonitor
	return nil
}

func (s *HealthChecksPingMonitorResourceCrud) Update() error {
	request := oci_health_checks.UpdatePingMonitorRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if intervalInSeconds, ok := s.D.GetOkExists("interval_in_seconds"); ok {
		tmp := intervalInSeconds.(int)
		request.IntervalInSeconds = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_health_checks.UpdatePingMonitorDetailsProtocolEnum(protocol.(string))
	}

	request.Targets = []string{}
	if targets, ok := s.D.GetOkExists("targets"); ok {
		interfaces := targets.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.Targets = tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	request.VantagePointNames = []string{}
	if vantagePointNames, ok := s.D.GetOkExists("vantage_point_names"); ok {
		interfaces := vantagePointNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		request.VantagePointNames = tmp
	}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.UpdatePingMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PingMonitor
	return nil
}

func (s *HealthChecksPingMonitorResourceCrud) Delete() error {
	request := oci_health_checks.DeletePingMonitorRequest{}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	_, err := s.Client.DeletePingMonitor(context.Background(), request)
	return err
}

func (s *HealthChecksPingMonitorResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IntervalInSeconds != nil {
		s.D.Set("interval_in_seconds", *s.Res.IntervalInSeconds)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	s.D.Set("protocol", s.Res.Protocol)

	if s.Res.ResultsUrl != nil {
		s.D.Set("results_url", *s.Res.ResultsUrl)
	}

	s.D.Set("targets", s.Res.Targets)

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	s.D.Set("vantage_point_names", s.Res.VantagePointNames)

	return nil
}
