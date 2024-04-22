// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes IPAM resource discoveries. A resource discovery is an IPAM component
// that enables IPAM to manage and monitor resources that belong to the owning
// account.
func (c *Client) DescribeIpamResourceDiscoveries(ctx context.Context, params *DescribeIpamResourceDiscoveriesInput, optFns ...func(*Options)) (*DescribeIpamResourceDiscoveriesOutput, error) {
	if params == nil {
		params = &DescribeIpamResourceDiscoveriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIpamResourceDiscoveries", params, optFns, c.addOperationDescribeIpamResourceDiscoveriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIpamResourceDiscoveriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIpamResourceDiscoveriesInput struct {

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The resource discovery filters.
	Filters []types.Filter

	// The IPAM resource discovery IDs.
	IpamResourceDiscoveryIds []string

	// The maximum number of resource discoveries to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeIpamResourceDiscoveriesOutput struct {

	// The resource discoveries.
	IpamResourceDiscoveries []types.IpamResourceDiscovery

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeIpamResourceDiscoveriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeIpamResourceDiscoveries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeIpamResourceDiscoveries{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIpamResourceDiscoveries(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeIpamResourceDiscoveriesAPIClient is a client that implements the
// DescribeIpamResourceDiscoveries operation.
type DescribeIpamResourceDiscoveriesAPIClient interface {
	DescribeIpamResourceDiscoveries(context.Context, *DescribeIpamResourceDiscoveriesInput, ...func(*Options)) (*DescribeIpamResourceDiscoveriesOutput, error)
}

var _ DescribeIpamResourceDiscoveriesAPIClient = (*Client)(nil)

// DescribeIpamResourceDiscoveriesPaginatorOptions is the paginator options for
// DescribeIpamResourceDiscoveries
type DescribeIpamResourceDiscoveriesPaginatorOptions struct {
	// The maximum number of resource discoveries to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeIpamResourceDiscoveriesPaginator is a paginator for
// DescribeIpamResourceDiscoveries
type DescribeIpamResourceDiscoveriesPaginator struct {
	options   DescribeIpamResourceDiscoveriesPaginatorOptions
	client    DescribeIpamResourceDiscoveriesAPIClient
	params    *DescribeIpamResourceDiscoveriesInput
	nextToken *string
	firstPage bool
}

// NewDescribeIpamResourceDiscoveriesPaginator returns a new
// DescribeIpamResourceDiscoveriesPaginator
func NewDescribeIpamResourceDiscoveriesPaginator(client DescribeIpamResourceDiscoveriesAPIClient, params *DescribeIpamResourceDiscoveriesInput, optFns ...func(*DescribeIpamResourceDiscoveriesPaginatorOptions)) *DescribeIpamResourceDiscoveriesPaginator {
	if params == nil {
		params = &DescribeIpamResourceDiscoveriesInput{}
	}

	options := DescribeIpamResourceDiscoveriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeIpamResourceDiscoveriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeIpamResourceDiscoveriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeIpamResourceDiscoveries page.
func (p *DescribeIpamResourceDiscoveriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeIpamResourceDiscoveriesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeIpamResourceDiscoveries(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeIpamResourceDiscoveries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeIpamResourceDiscoveries",
	}
}
