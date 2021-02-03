// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUploadChunkParams creates a new UploadChunkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadChunkParams() *UploadChunkParams {
	return &UploadChunkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadChunkParamsWithTimeout creates a new UploadChunkParams object
// with the ability to set a timeout on a request.
func NewUploadChunkParamsWithTimeout(timeout time.Duration) *UploadChunkParams {
	return &UploadChunkParams{
		timeout: timeout,
	}
}

// NewUploadChunkParamsWithContext creates a new UploadChunkParams object
// with the ability to set a context for a request.
func NewUploadChunkParamsWithContext(ctx context.Context) *UploadChunkParams {
	return &UploadChunkParams{
		Context: ctx,
	}
}

// NewUploadChunkParamsWithHTTPClient creates a new UploadChunkParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadChunkParamsWithHTTPClient(client *http.Client) *UploadChunkParams {
	return &UploadChunkParams{
		HTTPClient: client,
	}
}

/* UploadChunkParams contains all the parameters to send to the API endpoint
   for the upload chunk operation.

   Typically these are written to a http.Request.
*/
type UploadChunkParams struct {

	/* ChunkData.

	   Chunk file data
	*/
	ChunkData runtime.NamedReadCloser

	/* ChunkHash.

	   The file hash of the chunk
	*/
	ChunkHash string

	/* ChunkNumber.

	   The position index of the chunk in the file the chunk is from
	*/
	ChunkNumber int64

	/* UploadRequestID.

	   The identifier for the composite file upload request
	*/
	UploadRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload chunk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadChunkParams) WithDefaults() *UploadChunkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload chunk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadChunkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload chunk params
func (o *UploadChunkParams) WithTimeout(timeout time.Duration) *UploadChunkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload chunk params
func (o *UploadChunkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload chunk params
func (o *UploadChunkParams) WithContext(ctx context.Context) *UploadChunkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload chunk params
func (o *UploadChunkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload chunk params
func (o *UploadChunkParams) WithHTTPClient(client *http.Client) *UploadChunkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload chunk params
func (o *UploadChunkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChunkData adds the chunkData to the upload chunk params
func (o *UploadChunkParams) WithChunkData(chunkData runtime.NamedReadCloser) *UploadChunkParams {
	o.SetChunkData(chunkData)
	return o
}

// SetChunkData adds the chunkData to the upload chunk params
func (o *UploadChunkParams) SetChunkData(chunkData runtime.NamedReadCloser) {
	o.ChunkData = chunkData
}

// WithChunkHash adds the chunkHash to the upload chunk params
func (o *UploadChunkParams) WithChunkHash(chunkHash string) *UploadChunkParams {
	o.SetChunkHash(chunkHash)
	return o
}

// SetChunkHash adds the chunkHash to the upload chunk params
func (o *UploadChunkParams) SetChunkHash(chunkHash string) {
	o.ChunkHash = chunkHash
}

// WithChunkNumber adds the chunkNumber to the upload chunk params
func (o *UploadChunkParams) WithChunkNumber(chunkNumber int64) *UploadChunkParams {
	o.SetChunkNumber(chunkNumber)
	return o
}

// SetChunkNumber adds the chunkNumber to the upload chunk params
func (o *UploadChunkParams) SetChunkNumber(chunkNumber int64) {
	o.ChunkNumber = chunkNumber
}

// WithUploadRequestID adds the uploadRequestID to the upload chunk params
func (o *UploadChunkParams) WithUploadRequestID(uploadRequestID string) *UploadChunkParams {
	o.SetUploadRequestID(uploadRequestID)
	return o
}

// SetUploadRequestID adds the uploadRequestId to the upload chunk params
func (o *UploadChunkParams) SetUploadRequestID(uploadRequestID string) {
	o.UploadRequestID = uploadRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadChunkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param chunkData
	if err := r.SetFileParam("chunkData", o.ChunkData); err != nil {
		return err
	}

	// form param chunkHash
	frChunkHash := o.ChunkHash
	fChunkHash := frChunkHash
	if fChunkHash != "" {
		if err := r.SetFormParam("chunkHash", fChunkHash); err != nil {
			return err
		}
	}

	// form param chunkNumber
	frChunkNumber := o.ChunkNumber
	fChunkNumber := swag.FormatInt64(frChunkNumber)
	if fChunkNumber != "" {
		if err := r.SetFormParam("chunkNumber", fChunkNumber); err != nil {
			return err
		}
	}

	// form param uploadRequestId
	frUploadRequestID := o.UploadRequestID
	fUploadRequestID := frUploadRequestID
	if fUploadRequestID != "" {
		if err := r.SetFormParam("uploadRequestId", fUploadRequestID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
