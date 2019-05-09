// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"errors"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func IdentityAuthTokenResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createIdentityAuthToken,
		Read:     readIdentityAuthToken,
		Update:   updateIdentityAuthToken,
		Delete:   deleteIdentityAuthToken,
		Schema: map[string]*schema.Schema{
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"inactive_state": {
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
			"time_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"token": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return CreateResource(d, sync)
}

func readIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
}

func updateIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return UpdateResource(d, sync)
}

func deleteIdentityAuthToken(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityAuthTokenResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type IdentityAuthTokenResourceCrud struct {
	BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.AuthToken
	DisableNotFoundRetries bool
}

func (s *IdentityAuthTokenResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityAuthTokenResourceCrud) State() oci_identity.AuthTokenLifecycleStateEnum {
	return s.Res.LifecycleState
}

func (s *IdentityAuthTokenResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateCreating),
	}
}

func (s *IdentityAuthTokenResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateActive),
	}
}

func (s *IdentityAuthTokenResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateDeleting),
	}
}

func (s *IdentityAuthTokenResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.AuthTokenLifecycleStateDeleted),
	}
}

func (s *IdentityAuthTokenResourceCrud) Create() error {
	request := oci_identity.CreateAuthTokenRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateAuthToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthToken
	return nil
}

func (s *IdentityAuthTokenResourceCrud) Get() error {
	request := oci_identity.ListAuthTokensRequest{}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.ListAuthTokens(context.Background(), request)
	if err != nil {
		return err
	}

	id := s.D.Id()
	for _, item := range response.Items {
		if *item.Id == id {
			s.Res = &item
			return nil
		}
	}
	return errors.New("AuthToken with expected identifier not found")

}

func (s *IdentityAuthTokenResourceCrud) Update() error {
	request := oci_identity.UpdateAuthTokenRequest{}

	tmp := s.D.Id()
	request.AuthTokenId = &tmp

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateAuthToken(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuthToken
	return nil
}

func (s *IdentityAuthTokenResourceCrud) Delete() error {
	request := oci_identity.DeleteAuthTokenRequest{}

	tmp := s.D.Id()
	request.AuthTokenId = &tmp

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteAuthToken(context.Background(), request)
	return err
}

func (s *IdentityAuthTokenResourceCrud) SetData() error {
	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.String())
	}

	if s.Res.Token != nil && *s.Res.Token != "" {
		s.D.Set("token", *s.Res.Token)
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}
