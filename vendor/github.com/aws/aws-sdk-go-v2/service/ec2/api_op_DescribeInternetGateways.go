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

// Describes one or more of your internet gateways.
func (c *Client) DescribeInternetGateways(ctx context.Context, params *DescribeInternetGatewaysInput, optFns ...func(*Options)) (*DescribeInternetGatewaysOutput, error) {
	if params == nil {
		params = &DescribeInternetGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInternetGateways", params, optFns, addOperationDescribeInternetGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInternetGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInternetGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * attachment.state - The current state of the attachment
	// between the gateway and the VPC (available). Present only if a VPC is
	// attached.
	//
	// * attachment.vpc-id - The ID of an attached VPC.
	//
	// *
	// internet-gateway-id - The ID of the Internet gateway.
	//
	// * owner-id - The ID of
	// the AWS account that owns the internet gateway.
	//
	// * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	// * tag-key - The key of a tag
	// assigned to the resource. Use this filter to find all resources assigned a tag
	// with a specific key, regardless of the tag value.
	Filters []types.Filter

	// One or more internet gateway IDs. Default: Describes all your internet gateways.
	InternetGatewayIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeInternetGatewaysOutput struct {

	// Information about one or more internet gateways.
	InternetGateways []types.InternetGateway

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInternetGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInternetGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInternetGateways{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInternetGateways(options.Region), middleware.Before); err != nil {
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

// DescribeInternetGatewaysAPIClient is a client that implements the
// DescribeInternetGateways operation.
type DescribeInternetGatewaysAPIClient interface {
	DescribeInternetGateways(context.Context, *DescribeInternetGatewaysInput, ...func(*Options)) (*DescribeInternetGatewaysOutput, error)
}

var _ DescribeInternetGatewaysAPIClient = (*Client)(nil)

// DescribeInternetGatewaysPaginatorOptions is the paginator options for
// DescribeInternetGateways
type DescribeInternetGatewaysPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInternetGatewaysPaginator is a paginator for DescribeInternetGateways
type DescribeInternetGatewaysPaginator struct {
	options   DescribeInternetGatewaysPaginatorOptions
	client    DescribeInternetGatewaysAPIClient
	params    *DescribeInternetGatewaysInput
	nextToken *string
	firstPage bool
}

// NewDescribeInternetGatewaysPaginator returns a new
// DescribeInternetGatewaysPaginator
func NewDescribeInternetGatewaysPaginator(client DescribeInternetGatewaysAPIClient, params *DescribeInternetGatewaysInput, optFns ...func(*DescribeInternetGatewaysPaginatorOptions)) *DescribeInternetGatewaysPaginator {
	if params == nil {
		params = &DescribeInternetGatewaysInput{}
	}

	options := DescribeInternetGatewaysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInternetGatewaysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInternetGatewaysPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeInternetGateways page.
func (p *DescribeInternetGatewaysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInternetGatewaysOutput, error) {
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

	result, err := p.client.DescribeInternetGateways(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeInternetGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInternetGateways",
	}
}
