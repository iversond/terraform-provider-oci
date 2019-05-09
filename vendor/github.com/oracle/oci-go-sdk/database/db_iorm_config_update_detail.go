// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DbIormConfigUpdateDetail IORM Config setting request for this database
type DbIormConfigUpdateDetail struct {

	// Database Name. For updating default DbPlan, pass in dbName as `default`
	DbName *string `mandatory:"false" json:"dbName"`

	// Relative priority of a database
	Share *int `mandatory:"false" json:"share"`
}

func (m DbIormConfigUpdateDetail) String() string {
	return common.PointerString(m)
}
