// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSpotAdviceCommon = "DescribeSpotAdvice"

// DescribeSpotAdviceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSpotAdviceCommon operation. The "output" return
// value will be populated with the DescribeSpotAdviceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSpotAdviceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSpotAdviceCommon Send returns without error.
//
// See DescribeSpotAdviceCommon for more information on using the DescribeSpotAdviceCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSpotAdviceCommonRequest method.
//    req, resp := client.DescribeSpotAdviceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeSpotAdviceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSpotAdviceCommon,
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

// DescribeSpotAdviceCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeSpotAdviceCommon for usage and error information.
func (c *ECS) DescribeSpotAdviceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSpotAdviceCommonRequest(input)
	return out, req.Send()
}

// DescribeSpotAdviceCommonWithContext is the same as DescribeSpotAdviceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSpotAdviceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeSpotAdviceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSpotAdviceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSpotAdvice = "DescribeSpotAdvice"

// DescribeSpotAdviceRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSpotAdvice operation. The "output" return
// value will be populated with the DescribeSpotAdviceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSpotAdviceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSpotAdviceCommon Send returns without error.
//
// See DescribeSpotAdvice for more information on using the DescribeSpotAdvice
// API call, and error handling.
//
//    // Example sending a request using the DescribeSpotAdviceRequest method.
//    req, resp := client.DescribeSpotAdviceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeSpotAdviceRequest(input *DescribeSpotAdviceInput) (req *request.Request, output *DescribeSpotAdviceOutput) {
	op := &request.Operation{
		Name:       opDescribeSpotAdvice,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSpotAdviceInput{}
	}

	output = &DescribeSpotAdviceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSpotAdvice API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeSpotAdvice for usage and error information.
func (c *ECS) DescribeSpotAdvice(input *DescribeSpotAdviceInput) (*DescribeSpotAdviceOutput, error) {
	req, out := c.DescribeSpotAdviceRequest(input)
	return out, req.Send()
}

// DescribeSpotAdviceWithContext is the same as DescribeSpotAdvice with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSpotAdvice for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeSpotAdviceWithContext(ctx volcengine.Context, input *DescribeSpotAdviceInput, opts ...request.Option) (*DescribeSpotAdviceOutput, error) {
	req, out := c.DescribeSpotAdviceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AvailableSpotResourceForDescribeSpotAdviceOutput struct {
	_ struct{} `type:"structure"`

	AverageSpotDiscount *int32 `type:"int32"`

	InstanceType *string `type:"string"`

	InterruptionRate *int32 `type:"int32"`

	InterruptionRateRange *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s AvailableSpotResourceForDescribeSpotAdviceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AvailableSpotResourceForDescribeSpotAdviceOutput) GoString() string {
	return s.String()
}

// SetAverageSpotDiscount sets the AverageSpotDiscount field's value.
func (s *AvailableSpotResourceForDescribeSpotAdviceOutput) SetAverageSpotDiscount(v int32) *AvailableSpotResourceForDescribeSpotAdviceOutput {
	s.AverageSpotDiscount = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *AvailableSpotResourceForDescribeSpotAdviceOutput) SetInstanceType(v string) *AvailableSpotResourceForDescribeSpotAdviceOutput {
	s.InstanceType = &v
	return s
}

// SetInterruptionRate sets the InterruptionRate field's value.
func (s *AvailableSpotResourceForDescribeSpotAdviceOutput) SetInterruptionRate(v int32) *AvailableSpotResourceForDescribeSpotAdviceOutput {
	s.InterruptionRate = &v
	return s
}

// SetInterruptionRateRange sets the InterruptionRateRange field's value.
func (s *AvailableSpotResourceForDescribeSpotAdviceOutput) SetInterruptionRateRange(v string) *AvailableSpotResourceForDescribeSpotAdviceOutput {
	s.InterruptionRateRange = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *AvailableSpotResourceForDescribeSpotAdviceOutput) SetZoneId(v string) *AvailableSpotResourceForDescribeSpotAdviceOutput {
	s.ZoneId = &v
	return s
}

type DescribeSpotAdviceInput struct {
	_ struct{} `type:"structure"`

	Cpus *int32 `type:"int32"`

	Gpu *GpuForDescribeSpotAdviceInput `type:"structure"`

	InstanceTypeFamily *string `type:"string"`

	InstanceTypeIds []*string `type:"list"`

	MaxResults *int32 `type:"int32"`

	MemorySize *int32 `type:"int32"`

	MinCpus *int32 `type:"int32"`

	MinMemorySize *int32 `type:"int32"`

	NextToken *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSpotAdviceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSpotAdviceInput) GoString() string {
	return s.String()
}

// SetCpus sets the Cpus field's value.
func (s *DescribeSpotAdviceInput) SetCpus(v int32) *DescribeSpotAdviceInput {
	s.Cpus = &v
	return s
}

// SetGpu sets the Gpu field's value.
func (s *DescribeSpotAdviceInput) SetGpu(v *GpuForDescribeSpotAdviceInput) *DescribeSpotAdviceInput {
	s.Gpu = v
	return s
}

// SetInstanceTypeFamily sets the InstanceTypeFamily field's value.
func (s *DescribeSpotAdviceInput) SetInstanceTypeFamily(v string) *DescribeSpotAdviceInput {
	s.InstanceTypeFamily = &v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *DescribeSpotAdviceInput) SetInstanceTypeIds(v []*string) *DescribeSpotAdviceInput {
	s.InstanceTypeIds = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeSpotAdviceInput) SetMaxResults(v int32) *DescribeSpotAdviceInput {
	s.MaxResults = &v
	return s
}

// SetMemorySize sets the MemorySize field's value.
func (s *DescribeSpotAdviceInput) SetMemorySize(v int32) *DescribeSpotAdviceInput {
	s.MemorySize = &v
	return s
}

// SetMinCpus sets the MinCpus field's value.
func (s *DescribeSpotAdviceInput) SetMinCpus(v int32) *DescribeSpotAdviceInput {
	s.MinCpus = &v
	return s
}

// SetMinMemorySize sets the MinMemorySize field's value.
func (s *DescribeSpotAdviceInput) SetMinMemorySize(v int32) *DescribeSpotAdviceInput {
	s.MinMemorySize = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeSpotAdviceInput) SetNextToken(v string) *DescribeSpotAdviceInput {
	s.NextToken = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeSpotAdviceInput) SetZoneId(v string) *DescribeSpotAdviceInput {
	s.ZoneId = &v
	return s
}

type DescribeSpotAdviceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AvailableSpotResources []*AvailableSpotResourceForDescribeSpotAdviceOutput `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeSpotAdviceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSpotAdviceOutput) GoString() string {
	return s.String()
}

// SetAvailableSpotResources sets the AvailableSpotResources field's value.
func (s *DescribeSpotAdviceOutput) SetAvailableSpotResources(v []*AvailableSpotResourceForDescribeSpotAdviceOutput) *DescribeSpotAdviceOutput {
	s.AvailableSpotResources = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeSpotAdviceOutput) SetNextToken(v string) *DescribeSpotAdviceOutput {
	s.NextToken = &v
	return s
}

type GpuForDescribeSpotAdviceInput struct {
	_ struct{} `type:"structure"`

	Count *int32 `type:"int32"`

	ProductName *string `type:"string"`
}

// String returns the string representation
func (s GpuForDescribeSpotAdviceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GpuForDescribeSpotAdviceInput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *GpuForDescribeSpotAdviceInput) SetCount(v int32) *GpuForDescribeSpotAdviceInput {
	s.Count = &v
	return s
}

// SetProductName sets the ProductName field's value.
func (s *GpuForDescribeSpotAdviceInput) SetProductName(v string) *GpuForDescribeSpotAdviceInput {
	s.ProductName = &v
	return s
}
