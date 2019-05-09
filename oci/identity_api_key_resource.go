// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"errors"
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityApiKeyResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createIdentityApiKey,
		Read:     readIdentityApiKey,
		Delete:   deleteIdentityApiKey,
		Schema: map[string]*schema.Schema{
			// Required
			"key_value": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					r := regexp.MustCompile("\\s")
					strippedOld := r.ReplaceAllString(old, "")
					strippedNew := r.ReplaceAllString(new, "")
					return (strippedOld == strippedNew)
				},
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_status": {
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

func createIdentityApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return CreateResource(d, sync)
}

func readIdentityApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

func deleteIdentityApiKey(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityApiKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type IdentityApiKeyResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.ApiKey
	DisableNotFoundRetries bool
}

func (s *IdentityApiKeyResourceCrud) ID() string {
	return *s.Res.KeyId
}

func (s *IdentityApiKeyResourceCrud) State() oci_identity.ApiKeyLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *IdentityApiKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateCreating),
	}
}

func (s *IdentityApiKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateActive),
	}
}

func (s *IdentityApiKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateDeleting),
	}
}

func (s *IdentityApiKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.ApiKeyLifecycleStateDeleted),
	}
}

func (s *IdentityApiKeyResourceCrud) Create() error {
	request := oci_identity.UploadApiKeyRequest{}

	if keyValue, ok := s.D.GetOkExists("key_value"); ok {
		tmp := keyValue.(string)
		request.Key = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UploadApiKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ApiKey
	return nil
}

func (s *IdentityApiKeyResourceCrud) Get() error {
	request := oci_identity.ListApiKeysRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListApiKeys(context.Background(), request)
	if err != nil {
		return err
	}

	fingerprint := s.D.Get("fingerprint").(string)
	for _, item := range response.Items {
		if *item.Fingerprint == fingerprint {
			s.Res = &item
			return nil
		}
	}
	return errors.New("ApiKey with expected identifier not found")

}

func (s *IdentityApiKeyResourceCrud) Delete() error {
	request := oci_identity.DeleteApiKeyRequest{}

	if fingerprint, ok := s.D.GetOkExists("fingerprint"); ok {
		tmp := fingerprint.(string)
		request.Fingerprint = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteApiKey(context.Background(), request)
	return err
}

func (s *IdentityApiKeyResourceCrud) SetData() error {
	if s.Res.Fingerprint != nil {
		s.D.Set("fingerprint", *s.Res.Fingerprint)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_status", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.KeyValue != nil {
		s.D.Set("key_value", *s.Res.KeyValue)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}
