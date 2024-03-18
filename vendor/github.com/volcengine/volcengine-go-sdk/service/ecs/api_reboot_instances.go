// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRebootInstancesCommon = "RebootInstances"

// RebootInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RebootInstancesCommon operation. The "output" return
// value will be populated with the RebootInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RebootInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after RebootInstancesCommon Send returns without error.
//
// See RebootInstancesCommon for more information on using the RebootInstancesCommon
// API call, and error handling.
//
//	// Example sending a request using the RebootInstancesCommonRequest method.
//	req, resp := client.RebootInstancesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) RebootInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRebootInstancesCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// RebootInstancesCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation RebootInstancesCommon for usage and error information.
func (c *ECS) RebootInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RebootInstancesCommonRequest(input)
	return out, req.Send()
}

// RebootInstancesCommonWithContext is the same as RebootInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RebootInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RebootInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RebootInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRebootInstances = "RebootInstances"

// RebootInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the RebootInstances operation. The "output" return
// value will be populated with the RebootInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RebootInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after RebootInstancesCommon Send returns without error.
//
// See RebootInstances for more information on using the RebootInstances
// API call, and error handling.
//
//	// Example sending a request using the RebootInstancesRequest method.
//	req, resp := client.RebootInstancesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) RebootInstancesRequest(input *RebootInstancesInput) (req *request.Request, output *RebootInstancesOutput) {
	op := &request.Operation{
		Name:       opRebootInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebootInstancesInput{}
	}

	output = &RebootInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RebootInstances API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation RebootInstances for usage and error information.
func (c *ECS) RebootInstances(input *RebootInstancesInput) (*RebootInstancesOutput, error) {
	req, out := c.RebootInstancesRequest(input)
	return out, req.Send()
}

// RebootInstancesWithContext is the same as RebootInstances with the addition of
// the ability to pass a context and additional request options.
//
// See RebootInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) RebootInstancesWithContext(ctx volcengine.Context, input *RebootInstancesInput, opts ...request.Option) (*RebootInstancesOutput, error) {
	req, out := c.RebootInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ErrorForRebootInstancesOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForRebootInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForRebootInstancesOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForRebootInstancesOutput) SetCode(v string) *ErrorForRebootInstancesOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForRebootInstancesOutput) SetMessage(v string) *ErrorForRebootInstancesOutput {
	s.Message = &v
	return s
}

type OperationDetailForRebootInstancesOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForRebootInstancesOutput `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForRebootInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForRebootInstancesOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForRebootInstancesOutput) SetError(v *ErrorForRebootInstancesOutput) *OperationDetailForRebootInstancesOutput {
	s.Error = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *OperationDetailForRebootInstancesOutput) SetInstanceId(v string) *OperationDetailForRebootInstancesOutput {
	s.InstanceId = &v
	return s
}

type RebootInstancesInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	ForceStop *bool `type:"boolean"`

	InstanceIds []*string `type:"list"`
}

// String returns the string representation
func (s RebootInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RebootInstancesInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *RebootInstancesInput) SetClientToken(v string) *RebootInstancesInput {
	s.ClientToken = &v
	return s
}

// SetForceStop sets the ForceStop field's value.
func (s *RebootInstancesInput) SetForceStop(v bool) *RebootInstancesInput {
	s.ForceStop = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *RebootInstancesInput) SetInstanceIds(v []*string) *RebootInstancesInput {
	s.InstanceIds = v
	return s
}

type RebootInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForRebootInstancesOutput `type:"list"`
}

// String returns the string representation
func (s RebootInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RebootInstancesOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *RebootInstancesOutput) SetOperationDetails(v []*OperationDetailForRebootInstancesOutput) *RebootInstancesOutput {
	s.OperationDetails = v
	return s
}
