// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFlagByNameBatchParams creates a new GetFlagByNameBatchParams object
// no default values defined in spec.
func NewGetFlagByNameBatchParams() GetFlagByNameBatchParams {

	return GetFlagByNameBatchParams{}
}

// GetFlagByNameBatchParams contains all the bound params for the get flag by name batch operation
// typically these are obtained from a http.Request
//
// swagger:parameters getFlagByNameBatch
type GetFlagByNameBatchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  Min Items: 1
	  In: query
	*/
	FlagNames []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFlagByNameBatchParams() beforehand.
func (o *GetFlagByNameBatchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFlagNames, qhkFlagNames, _ := qs.GetOK("flagNames")
	if err := o.bindFlagNames(qFlagNames, qhkFlagNames, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFlagByNameBatchParams) bindFlagNames(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("flagNames", "query")
	}

	var qvFlagNames string
	if len(rawData) > 0 {
		qvFlagNames = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	flagNamesIC := swag.SplitByFormat(qvFlagNames, "")

	if len(flagNamesIC) == 0 {
		return errors.Required("flagNames", "query")
	}

	var flagNamesIR []string
	for i, flagNamesIV := range flagNamesIC {
		flagNamesI := flagNamesIV

		if err := validate.MinLength(fmt.Sprintf("%s.%v", "flagNames", i), "query", flagNamesI, 1); err != nil {
			return err
		}

		flagNamesIR = append(flagNamesIR, flagNamesI)
	}

	o.FlagNames = flagNamesIR
	if err := o.validateFlagNames(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetFlagByNameBatchParams) validateFlagNames(formats strfmt.Registry) error {

	flagNamesSize := int64(len(o.FlagNames))

	// minItems: 1
	if err := validate.MinItems("flagNames", "query", flagNamesSize, 1); err != nil {
		return err
	}

	return nil
}
