// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// GetLabelsOfNodeOKCode is the HTTP code returned for type GetLabelsOfNodeOK
const GetLabelsOfNodeOKCode int = 200

/*GetLabelsOfNodeOK Detailed namespace and namespace information.

swagger:response getLabelsOfNodeOK
*/
type GetLabelsOfNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.LabelsResponse `json:"body,omitempty"`
}

// NewGetLabelsOfNodeOK creates GetLabelsOfNodeOK with default headers values
func NewGetLabelsOfNodeOK() *GetLabelsOfNodeOK {

	return &GetLabelsOfNodeOK{}
}

// WithPayload adds the payload to the get labels of node o k response
func (o *GetLabelsOfNodeOK) WithPayload(payload *models.LabelsResponse) *GetLabelsOfNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get labels of node o k response
func (o *GetLabelsOfNodeOK) SetPayload(payload *models.LabelsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLabelsOfNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLabelsOfNodeUnauthorizedCode is the HTTP code returned for type GetLabelsOfNodeUnauthorized
const GetLabelsOfNodeUnauthorizedCode int = 401

/*GetLabelsOfNodeUnauthorized Unauthorized

swagger:response getLabelsOfNodeUnauthorized
*/
type GetLabelsOfNodeUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLabelsOfNodeUnauthorized creates GetLabelsOfNodeUnauthorized with default headers values
func NewGetLabelsOfNodeUnauthorized() *GetLabelsOfNodeUnauthorized {

	return &GetLabelsOfNodeUnauthorized{}
}

// WithPayload adds the payload to the get labels of node unauthorized response
func (o *GetLabelsOfNodeUnauthorized) WithPayload(payload *models.Error) *GetLabelsOfNodeUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get labels of node unauthorized response
func (o *GetLabelsOfNodeUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLabelsOfNodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLabelsOfNodeNotFoundCode is the HTTP code returned for type GetLabelsOfNodeNotFound
const GetLabelsOfNodeNotFoundCode int = 404

/*GetLabelsOfNodeNotFound url to add namespace not found.

swagger:response getLabelsOfNodeNotFound
*/
type GetLabelsOfNodeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLabelsOfNodeNotFound creates GetLabelsOfNodeNotFound with default headers values
func NewGetLabelsOfNodeNotFound() *GetLabelsOfNodeNotFound {

	return &GetLabelsOfNodeNotFound{}
}

// WithPayload adds the payload to the get labels of node not found response
func (o *GetLabelsOfNodeNotFound) WithPayload(payload *models.Error) *GetLabelsOfNodeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get labels of node not found response
func (o *GetLabelsOfNodeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLabelsOfNodeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}