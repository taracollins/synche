// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gitlab.computing.dcu.ie/collint9/2021-ca400-collint9-coynemt2/src/server/models"
)

// NewUploadOKCode is the HTTP code returned for type NewUploadOK
const NewUploadOKCode int = 200

/*NewUploadOK OK

swagger:response newUploadOK
*/
type NewUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.NewFileUploadRequestAccepted `json:"body,omitempty"`
}

// NewNewUploadOK creates NewUploadOK with default headers values
func NewNewUploadOK() *NewUploadOK {

	return &NewUploadOK{}
}

// WithPayload adds the payload to the new upload o k response
func (o *NewUploadOK) WithPayload(payload *models.NewFileUploadRequestAccepted) *NewUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new upload o k response
func (o *NewUploadOK) SetPayload(payload *models.NewFileUploadRequestAccepted) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewUploadBadRequestCode is the HTTP code returned for type NewUploadBadRequest
const NewUploadBadRequestCode int = 400

/*NewUploadBadRequest bad request

swagger:response newUploadBadRequest
*/
type NewUploadBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNewUploadBadRequest creates NewUploadBadRequest with default headers values
func NewNewUploadBadRequest() *NewUploadBadRequest {

	return &NewUploadBadRequest{}
}

// WithPayload adds the payload to the new upload bad request response
func (o *NewUploadBadRequest) WithPayload(payload *models.Error) *NewUploadBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new upload bad request response
func (o *NewUploadBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewUploadBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NewUploadNotImplementedCode is the HTTP code returned for type NewUploadNotImplemented
const NewUploadNotImplementedCode int = 501

/*NewUploadNotImplemented not implemented

swagger:response newUploadNotImplemented
*/
type NewUploadNotImplemented struct {

	/*
	  In: Body
	*/
	Payload models.NotImplemented `json:"body,omitempty"`
}

// NewNewUploadNotImplemented creates NewUploadNotImplemented with default headers values
func NewNewUploadNotImplemented() *NewUploadNotImplemented {

	return &NewUploadNotImplemented{}
}

// WithPayload adds the payload to the new upload not implemented response
func (o *NewUploadNotImplemented) WithPayload(payload models.NotImplemented) *NewUploadNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the new upload not implemented response
func (o *NewUploadNotImplemented) SetPayload(payload models.NotImplemented) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NewUploadNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
