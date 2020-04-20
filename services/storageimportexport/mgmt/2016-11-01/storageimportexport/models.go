package storageimportexport

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/test-repo-arcturus/azure-sdk-for-go/services/storageimportexport/mgmt/2016-11-01/storageimportexport"

// DriveState enumerates the values for drive state.
type DriveState string

const (
	// Completed ...
	Completed DriveState = "Completed"
	// CompletedMoreInfo ...
	CompletedMoreInfo DriveState = "CompletedMoreInfo"
	// NeverReceived ...
	NeverReceived DriveState = "NeverReceived"
	// Received ...
	Received DriveState = "Received"
	// ShippedBack ...
	ShippedBack DriveState = "ShippedBack"
	// Specified ...
	Specified DriveState = "Specified"
	// Transferring ...
	Transferring DriveState = "Transferring"
)

// PossibleDriveStateValues returns an array of possible values for the DriveState const type.
func PossibleDriveStateValues() []DriveState {
	return []DriveState{Completed, CompletedMoreInfo, NeverReceived, Received, ShippedBack, Specified, Transferring}
}

// DriveBitLockerKey bitLocker recovery key or password to the specified drive
type DriveBitLockerKey struct {
	// BitLockerKey - BitLocker recovery key or password
	BitLockerKey *string `json:"bitLockerKey,omitempty"`
	// DriveID - Drive ID
	DriveID *string `json:"driveId,omitempty"`
}

// DriveStatus provides information about the drive's status
type DriveStatus struct {
	// DriveID - The drive's hardware serial number, without spaces.
	DriveID *string `json:"driveId,omitempty"`
	// BitLockerKey - The BitLocker key used to encrypt the drive.
	BitLockerKey *string `json:"bitLockerKey,omitempty"`
	// ManifestFile - The relative path of the manifest file on the drive.
	ManifestFile *string `json:"manifestFile,omitempty"`
	// ManifestHash - The Base16-encoded MD5 hash of the manifest file on the drive.
	ManifestHash *string `json:"manifestHash,omitempty"`
	// DriveHeaderHash - The drive header hash value.
	DriveHeaderHash *string `json:"driveHeaderHash,omitempty"`
	// State - The drive's current state. Possible values include: 'Specified', 'Received', 'NeverReceived', 'Transferring', 'Completed', 'CompletedMoreInfo', 'ShippedBack'
	State DriveState `json:"state,omitempty"`
	// CopyStatus - Detailed status about the data transfer process. This field is not returned in the response until the drive is in the Transferring state.
	CopyStatus *string `json:"copyStatus,omitempty"`
	// PercentComplete - Percentage completed for the drive.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// VerboseLogURI - A URI that points to the blob containing the verbose log for the data transfer operation.
	VerboseLogURI *string `json:"verboseLogUri,omitempty"`
	// ErrorLogURI - A URI that points to the blob containing the error log for the data transfer operation.
	ErrorLogURI *string `json:"errorLogUri,omitempty"`
	// ManifestURI - A URI that points to the blob containing the drive manifest file.
	ManifestURI *string `json:"manifestUri,omitempty"`
	// BytesSucceeded - Bytes successfully transferred for the drive.
	BytesSucceeded *int64 `json:"bytesSucceeded,omitempty"`
}

// ErrorResponse response when errors occurred
type ErrorResponse struct {
	// ErrorResponseError - Describes the error information.
	*ErrorResponseError `json:"error,omitempty"`
}

// MarshalJSON is the custom marshaler for ErrorResponse.
func (er ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if er.ErrorResponseError != nil {
		objectMap["error"] = er.ErrorResponseError
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ErrorResponse struct.
func (er *ErrorResponse) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "error":
			if v != nil {
				var errorResponseError ErrorResponseError
				err = json.Unmarshal(*v, &errorResponseError)
				if err != nil {
					return err
				}
				er.ErrorResponseError = &errorResponseError
			}
		}
	}

	return nil
}

// ErrorResponseError describes the error information.
type ErrorResponseError struct {
	// Code - Provides information about the error code.
	Code *string `json:"code,omitempty"`
	// Message - Provides information about the error message.
	Message *string `json:"message,omitempty"`
	// Target - Provides information about the error target.
	Target *string `json:"target,omitempty"`
	// Details - Describes the error details if present.
	Details *[]ErrorResponseErrorDetailsItem `json:"details,omitempty"`
	// Innererror - Inner error object if present.
	Innererror interface{} `json:"innererror,omitempty"`
}

// ErrorResponseErrorDetailsItem ...
type ErrorResponseErrorDetailsItem struct {
	// Code - Provides information about the error code.
	Code *string `json:"code,omitempty"`
	// Target - Provides information about the error target.
	Target *string `json:"target,omitempty"`
	// Message - Provides information about the error message.
	Message *string `json:"message,omitempty"`
}

// Export a property containing information about the blobs to be exported for an export job. This property
// is required for export jobs, but must not be specified for import jobs.
type Export struct {
	// ExportBlobList - A list of the blobs to be exported.
	*ExportBlobList `json:"blobList,omitempty"`
	// BlobListblobPath - The relative URI to the block blob that contains the list of blob paths or blob path prefixes as defined above, beginning with the container name. If the blob is in root container, the URI must begin with $root.
	BlobListblobPath *string `json:"blobListblobPath,omitempty"`
}

// MarshalJSON is the custom marshaler for Export.
func (e Export) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if e.ExportBlobList != nil {
		objectMap["blobList"] = e.ExportBlobList
	}
	if e.BlobListblobPath != nil {
		objectMap["blobListblobPath"] = e.BlobListblobPath
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Export struct.
func (e *Export) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "blobList":
			if v != nil {
				var exportBlobList ExportBlobList
				err = json.Unmarshal(*v, &exportBlobList)
				if err != nil {
					return err
				}
				e.ExportBlobList = &exportBlobList
			}
		case "blobListblobPath":
			if v != nil {
				var blobListblobPath string
				err = json.Unmarshal(*v, &blobListblobPath)
				if err != nil {
					return err
				}
				e.BlobListblobPath = &blobListblobPath
			}
		}
	}

	return nil
}

// ExportBlobList a list of the blobs to be exported.
type ExportBlobList struct {
	// BlobPath - A collection of blob-path strings.
	BlobPath *[]string `json:"blobPath,omitempty"`
	// BlobPathPrefix - A collection of blob-prefix strings.
	BlobPathPrefix *[]string `json:"blobPathPrefix,omitempty"`
}

// GetBitLockerKeysResponse getBitLockerKeys response
type GetBitLockerKeysResponse struct {
	autorest.Response `json:"-"`
	// Value - drive status
	Value *[]DriveBitLockerKey `json:"value,omitempty"`
}

// JobDetails specifies the job properties
type JobDetails struct {
	// StorageAccountID - The resource identifier of the storage account where data will be imported to or exported from.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
	// JobType - The type of job
	JobType *string `json:"jobType,omitempty"`
	// ReturnAddress - Specifies the return address information for the job.
	ReturnAddress *ReturnAddress `json:"returnAddress,omitempty"`
	// ReturnShipping - Specifies the return carrier and customer's account with the carrier.
	ReturnShipping *ReturnShipping `json:"returnShipping,omitempty"`
	// ShippingInformation - Contains information about the Microsoft datacenter to which the drives should be shipped.
	ShippingInformation *ShippingInformation `json:"shippingInformation,omitempty"`
	// DeliveryPackage - Contains information about the package being shipped by the customer to the Microsoft data center.
	DeliveryPackage *PackageInfomation `json:"deliveryPackage,omitempty"`
	// ReturnPackage - Contains information about the package being shipped from the Microsoft data center to the customer to return the drives. The format is the same as the deliveryPackage property above. This property is not included if the drives have not yet been returned.
	ReturnPackage *PackageInfomation `json:"returnPackage,omitempty"`
	// DiagnosticsPath - The virtual blob directory to which the copy logs and backups of drive manifest files (if enabled) will be stored.
	DiagnosticsPath *string `json:"diagnosticsPath,omitempty"`
	// LogLevel - Default value is Error. Indicates whether error logging or verbose logging will be enabled.
	LogLevel *string `json:"logLevel,omitempty"`
	// BackupDriveManifest - Default value is false. Indicates whether the manifest files on the drives should be copied to block blobs.
	BackupDriveManifest *bool `json:"backupDriveManifest,omitempty"`
	// State - Current state of the job.
	State *string `json:"state,omitempty"`
	// CancelRequested - Indicates whether a request has been submitted to cancel the job.
	CancelRequested *bool `json:"cancelRequested,omitempty"`
	// PercentComplete - Overall percentage completed for the job.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// IncompleteBlobListURI - A blob path that points to a block blob containing a list of blob names that were not exported due to insufficient drive space. If all blobs were exported successfully, then this element is not included in the response.
	IncompleteBlobListURI *string `json:"incompleteBlobListUri,omitempty"`
	// DriveList - List of up to ten drives that comprise the job. The drive list is a required element for an import job; it is not specified for export jobs.
	DriveList *[]DriveStatus `json:"driveList,omitempty"`
	// Export - A property containing information about the blobs to be exported for an export job. This property is included for export jobs only.
	Export *Export `json:"export,omitempty"`
	// ProvisioningState - Specifies the provisioning state of the job.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// JobResponse contains the job information.
type JobResponse struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Specifies the resource identifier of the job.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Specifies the name of the job.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Specifies the type of the job resource.
	Type *string `json:"type,omitempty"`
	// Location - Specifies the Azure location where the job is created.
	Location *string `json:"location,omitempty"`
	// Tags - Specifies the tags that are assigned to the job.
	Tags interface{} `json:"tags,omitempty"`
	// Properties - Specifies the job properties
	Properties *JobDetails `json:"properties,omitempty"`
}

// ListJobsResponse list jobs response
type ListJobsResponse struct {
	autorest.Response `json:"-"`
	// NextLink - link to next batch of jobs
	NextLink *string `json:"nextLink,omitempty"`
	// Value - Job list
	Value *[]JobResponse `json:"value,omitempty"`
}

// ListJobsResponseIterator provides access to a complete listing of JobResponse values.
type ListJobsResponseIterator struct {
	i    int
	page ListJobsResponsePage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListJobsResponseIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListJobsResponseIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ListJobsResponseIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListJobsResponseIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListJobsResponseIterator) Response() ListJobsResponse {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListJobsResponseIterator) Value() JobResponse {
	if !iter.page.NotDone() {
		return JobResponse{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListJobsResponseIterator type.
func NewListJobsResponseIterator(page ListJobsResponsePage) ListJobsResponseIterator {
	return ListJobsResponseIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ljr ListJobsResponse) IsEmpty() bool {
	return ljr.Value == nil || len(*ljr.Value) == 0
}

// listJobsResponsePreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ljr ListJobsResponse) listJobsResponsePreparer(ctx context.Context) (*http.Request, error) {
	if ljr.NextLink == nil || len(to.String(ljr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ljr.NextLink)))
}

// ListJobsResponsePage contains a page of JobResponse values.
type ListJobsResponsePage struct {
	fn  func(context.Context, ListJobsResponse) (ListJobsResponse, error)
	ljr ListJobsResponse
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListJobsResponsePage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListJobsResponsePage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ljr)
	if err != nil {
		return err
	}
	page.ljr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListJobsResponsePage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListJobsResponsePage) NotDone() bool {
	return !page.ljr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListJobsResponsePage) Response() ListJobsResponse {
	return page.ljr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListJobsResponsePage) Values() []JobResponse {
	if page.ljr.IsEmpty() {
		return nil
	}
	return *page.ljr.Value
}

// Creates a new instance of the ListJobsResponsePage type.
func NewListJobsResponsePage(getNextPage func(context.Context, ListJobsResponse) (ListJobsResponse, error)) ListJobsResponsePage {
	return ListJobsResponsePage{fn: getNextPage}
}

// ListOperationsResponse list operations response
type ListOperationsResponse struct {
	autorest.Response `json:"-"`
	// Value - operations
	Value *[]Operation `json:"value,omitempty"`
}

// Location provides information about an Azure data center location.
type Location struct {
	autorest.Response `json:"-"`
	// ID - Specifies the resource identifier of the location.
	ID *string `json:"id,omitempty"`
	// Name - Specifies the name of the location. Use List Locations to get all supported locations.
	Name *string `json:"name,omitempty"`
	// Type - Specifies the type of the location.
	Type *string `json:"type,omitempty"`
	// LocationProperties - location properties
	*LocationProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Location.
func (l Location) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if l.ID != nil {
		objectMap["id"] = l.ID
	}
	if l.Name != nil {
		objectMap["name"] = l.Name
	}
	if l.Type != nil {
		objectMap["type"] = l.Type
	}
	if l.LocationProperties != nil {
		objectMap["properties"] = l.LocationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Location struct.
func (l *Location) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				l.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				l.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				l.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var locationProperties LocationProperties
				err = json.Unmarshal(*v, &locationProperties)
				if err != nil {
					return err
				}
				l.LocationProperties = &locationProperties
			}
		}
	}

	return nil
}

// LocationProperties location properties
type LocationProperties struct {
	// RecipientName - The recipient name to use when shipping the drives to the Azure data center.
	RecipientName *string `json:"recipientName,omitempty"`
	// StreetAddress1 - The first line of the street address to use when shipping the drives to the Azure data center.
	StreetAddress1 *string `json:"streetAddress1,omitempty"`
	// StreetAddress2 - The second line of the street address to use when shipping the drives to the Azure data center.
	StreetAddress2 *string `json:"streetAddress2,omitempty"`
	// City - The city name to use when shipping the drives to the Azure data center.
	City *string `json:"city,omitempty"`
	// StateOrProvince - The state or province to use when shipping the drives to the Azure data center.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// PostalCode - The postal code to use when shipping the drives to the Azure data center.
	PostalCode *string `json:"postalCode,omitempty"`
	// CountryOrRegion - The country or region to use when shipping the drives to the Azure data center.
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// Phone - The phone number for the Azure data center.
	Phone *string `json:"phone,omitempty"`
	// SupportedCarriers - A list of carriers that are supported at this location.
	SupportedCarriers *[]string `json:"supportedCarriers,omitempty"`
	// AlternateLocations - A list of location IDs that should be used to ship shipping drives to for jobs created against the current location. If the current location is active, it will be part of the list. If it is temporarily closed due to maintenance, this list may contain other locations.
	AlternateLocations *[]string `json:"alternateLocations,omitempty"`
}

// LocationsResponse locations response
type LocationsResponse struct {
	autorest.Response `json:"-"`
	// Value - locations
	Value *[]Location `json:"value,omitempty"`
}

// Operation describes a supported operation by the Storage Import/Export job API.
type Operation struct {
	// Name - Name of the operation.
	Name *string `json:"name,omitempty"`
	// OperationDisplay - operation display properties
	*OperationDisplay `json:"display,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Name != nil {
		objectMap["name"] = o.Name
	}
	if o.OperationDisplay != nil {
		objectMap["display"] = o.OperationDisplay
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Operation struct.
func (o *Operation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				o.Name = &name
			}
		case "display":
			if v != nil {
				var operationDisplay OperationDisplay
				err = json.Unmarshal(*v, &operationDisplay)
				if err != nil {
					return err
				}
				o.OperationDisplay = &operationDisplay
			}
		}
	}

	return nil
}

// OperationDisplay operation display properties
type OperationDisplay struct {
	// Provider - The resource provider name to which the operation belongs.
	Provider *string `json:"provider,omitempty"`
	// Resource - The name of the resource to which the operation belongs.
	Resource *string `json:"resource,omitempty"`
	// Operation - The display name of the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Short description of the operation.
	Description *string `json:"description,omitempty"`
}

// PackageInfomation contains information about the package being shipped by the customer to the Microsoft
// data center.
type PackageInfomation struct {
	// CarrierName - The name of the carrier that is used to ship the import or export drives.
	CarrierName *string `json:"carrierName,omitempty"`
	// TrackingNumber - The tracking number of the package.
	TrackingNumber *string `json:"trackingNumber,omitempty"`
	// DriveCount - The number of drives included in the package.
	DriveCount *int32 `json:"driveCount,omitempty"`
	// ShipDate - The date when the package is shipped.
	ShipDate *string `json:"shipDate,omitempty"`
}

// PutJobParameters put Job parameters
type PutJobParameters struct {
	// Location - Specifies the supported Azure location where the job should be created
	Location *string `json:"location,omitempty"`
	// Tags - Specifies the tags that will be assigned to the job.
	Tags interface{} `json:"tags,omitempty"`
	// Properties - Specifies the job properties
	Properties *JobDetails `json:"properties,omitempty"`
}

// ReturnAddress specifies the return address information for the job.
type ReturnAddress struct {
	// RecipientName - The name of the recipient who will receive the hard drives when they are returned.
	RecipientName *string `json:"recipientName,omitempty"`
	// StreetAddress1 - The first line of the street address to use when returning the drives.
	StreetAddress1 *string `json:"streetAddress1,omitempty"`
	// StreetAddress2 - The second line of the street address to use when returning the drives.
	StreetAddress2 *string `json:"streetAddress2,omitempty"`
	// City - The city name to use when returning the drives.
	City *string `json:"city,omitempty"`
	// StateOrProvince - The state or province to use when returning the drives.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// PostalCode - The postal code to use when returning the drives.
	PostalCode *string `json:"postalCode,omitempty"`
	// CountryOrRegion - The country or region to use when returning the drives.
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// Phone - Phone number of the recipient of the returned drives.
	Phone *string `json:"phone,omitempty"`
	// Email - Email address of the recipient of the returned drives.
	Email *string `json:"email,omitempty"`
}

// ReturnShipping specifies the return carrier and customer's account with the carrier.
type ReturnShipping struct {
	// CarrierName - The carrier's name.
	CarrierName *string `json:"carrierName,omitempty"`
	// CarrierAccountNumber - The customer's account number with the carrier.
	CarrierAccountNumber *string `json:"carrierAccountNumber,omitempty"`
}

// ShippingInformation contains information about the Microsoft datacenter to which the drives should be
// shipped.
type ShippingInformation struct {
	// RecipientName - The name of the recipient who will receive the hard drives when they are returned.
	RecipientName *string `json:"recipientName,omitempty"`
	// StreetAddress1 - The first line of the street address to use when returning the drives.
	StreetAddress1 *string `json:"streetAddress1,omitempty"`
	// StreetAddress2 - The second line of the street address to use when returning the drives.
	StreetAddress2 *string `json:"streetAddress2,omitempty"`
	// City - The city name to use when returning the drives.
	City *string `json:"city,omitempty"`
	// StateOrProvince - The state or province to use when returning the drives.
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
	// PostalCode - The postal code to use when returning the drives.
	PostalCode *string `json:"postalCode,omitempty"`
	// CountryOrRegion - The country or region to use when returning the drives.
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// Phone - Phone number of the recipient of the returned drives.
	Phone *string `json:"phone,omitempty"`
}

// UpdateJobParameters update Job parameters
type UpdateJobParameters struct {
	// Tags - Specifies the tags that will be assigned to the job
	Tags interface{} `json:"tags,omitempty"`
	// UpdateJobParametersProperties - Specifies the properties of a UpdateJob.
	*UpdateJobParametersProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for UpdateJobParameters.
func (ujp UpdateJobParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ujp.Tags != nil {
		objectMap["tags"] = ujp.Tags
	}
	if ujp.UpdateJobParametersProperties != nil {
		objectMap["properties"] = ujp.UpdateJobParametersProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for UpdateJobParameters struct.
func (ujp *UpdateJobParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "tags":
			if v != nil {
				var tags interface{}
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				ujp.Tags = tags
			}
		case "properties":
			if v != nil {
				var updateJobParametersProperties UpdateJobParametersProperties
				err = json.Unmarshal(*v, &updateJobParametersProperties)
				if err != nil {
					return err
				}
				ujp.UpdateJobParametersProperties = &updateJobParametersProperties
			}
		}
	}

	return nil
}

// UpdateJobParametersProperties specifies the properties of a UpdateJob.
type UpdateJobParametersProperties struct {
	// CancelRequested - If specified, the value must be true. The service will attempt to cancel the job.
	CancelRequested *bool `json:"cancelRequested,omitempty"`
	// State - If specified, the value must be Shipping, which tells the Import/Export service that the package for the job has been shipped. The ReturnAddress and DeliveryPackage properties must have been set either in this request or in a previous request, otherwise the request will fail.
	State *string `json:"state,omitempty"`
	// ReturnAddress - Specifies the return address information for the job.
	ReturnAddress *ReturnAddress `json:"returnAddress,omitempty"`
	// ReturnShipping - Specifies the return carrier and customer's account with the carrier.
	ReturnShipping *ReturnShipping `json:"returnShipping,omitempty"`
	// DeliveryPackage - Contains information about the package being shipped by the customer to the Microsoft data center.
	DeliveryPackage *PackageInfomation `json:"deliveryPackage,omitempty"`
	// LogLevel - Indicates whether error logging or verbose logging is enabled.
	LogLevel *string `json:"logLevel,omitempty"`
	// BackupDriveManifest - Indicates whether the manifest files on the drives should be copied to block blobs.
	BackupDriveManifest *bool `json:"backupDriveManifest,omitempty"`
	// DriveList - List of drives that comprise the job.
	DriveList *[]DriveStatus `json:"driveList,omitempty"`
}
