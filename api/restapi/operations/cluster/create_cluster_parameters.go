// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/linkall-labs/vanus-operator/api/models"
)

// NewCreateClusterParams creates a new CreateClusterParams object
//
// There are no default values defined in the spec.
func NewCreateClusterParams() CreateClusterParams {

	return CreateClusterParams{}
}

// CreateClusterParams contains all the bound params for the create cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters createCluster
type CreateClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*需要创建的Cluster信息
	  Required: true
	  In: body
	*/
	Cluster *models.ClusterInfo
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateClusterParams() beforehand.
func (o *CreateClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ClusterInfo
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("cluster", "body", ""))
			} else {
				res = append(res, errors.NewParseError("cluster", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Cluster = &body
			}
		}
	} else {
		res = append(res, errors.Required("cluster", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}