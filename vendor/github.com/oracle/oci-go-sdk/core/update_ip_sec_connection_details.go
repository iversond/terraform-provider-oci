// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateIpSecConnectionDetails The representation of UpdateIpSecConnectionDetails
type UpdateIpSecConnectionDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Your identifier for your CPE device. Can be either an IP address or a hostname (specifically, the
	// fully qualified domain name (FQDN)). The type of identifier you provide here must correspond
	// to the value for `cpeLocalIdentifierType`.
	// Example IP address: `10.0.3.3`
	// Example hostname: `cpe.example.com`
	CpeLocalIdentifier *string `mandatory:"false" json:"cpeLocalIdentifier"`

	// The type of identifier for your CPE device. The value you provide here must correspond to the value
	// for `cpeLocalIdentifier`.
	CpeLocalIdentifierType UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum `mandatory:"false" json:"cpeLocalIdentifierType,omitempty"`

	// Static routes to the CPE. If you provide this attribute, it replaces the entire current set of
	// static routes. A static route's CIDR must not be a multicast address or class E address.
	// Example: `10.0.1.0/24`
	StaticRoutes []string `mandatory:"false" json:"staticRoutes"`
}

func (m UpdateIpSecConnectionDetails) String() string {
	return common.PointerString(m)
}

// UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum Enum with underlying type: string
type UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum string

// Set of constants representing the allowable values for UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum
const (
	UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeIpAddress UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum = "IP_ADDRESS"
	UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeHostname  UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum = "HOSTNAME"
)

var mappingUpdateIpSecConnectionDetailsCpeLocalIdentifierType = map[string]UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum{
	"IP_ADDRESS": UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeIpAddress,
	"HOSTNAME":   UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeHostname,
}

// GetUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumValues Enumerates the set of values for UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum
func GetUpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnumValues() []UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum {
	values := make([]UpdateIpSecConnectionDetailsCpeLocalIdentifierTypeEnum, 0)
	for _, v := range mappingUpdateIpSecConnectionDetailsCpeLocalIdentifierType {
		values = append(values, v)
	}
	return values
}
