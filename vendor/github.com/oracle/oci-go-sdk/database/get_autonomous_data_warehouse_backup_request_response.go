// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetAutonomousDataWarehouseBackupRequest wrapper for the GetAutonomousDataWarehouseBackup operation
type GetAutonomousDataWarehouseBackupRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Warehouse backup.
	AutonomousDataWarehouseBackupId *string `mandatory:"true" contributesTo:"path" name:"autonomousDataWarehouseBackupId"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAutonomousDataWarehouseBackupRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAutonomousDataWarehouseBackupRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAutonomousDataWarehouseBackupRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAutonomousDataWarehouseBackupResponse wrapper for the GetAutonomousDataWarehouseBackup operation
type GetAutonomousDataWarehouseBackupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutonomousDataWarehouseBackup instance
	AutonomousDataWarehouseBackup `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAutonomousDataWarehouseBackupResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAutonomousDataWarehouseBackupResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
