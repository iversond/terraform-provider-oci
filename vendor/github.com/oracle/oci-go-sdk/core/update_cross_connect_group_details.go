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

// UpdateCrossConnectGroupDetails The representation of UpdateCrossConnectGroupDetails
type UpdateCrossConnectGroupDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A reference name or identifier for the physical fiber connection that this cross-connect
	// group uses.
	CustomerReferenceName *string `mandatory:"false" json:"customerReferenceName"`
}

func (m UpdateCrossConnectGroupDetails) String() string {
	return common.PointerString(m)
}
