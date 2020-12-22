// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of dashboards for an AWS IoT SiteWise Monitor
// project.
func (c *Client) ListDashboards(ctx context.Context, params *ListDashboardsInput, optFns ...func(*Options)) (*ListDashboardsOutput, error) {
	if params == nil {
		params = &ListDashboardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDashboards", params, optFns, addOperationListDashboardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDashboardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDashboardsInput struct {

	// The ID of the project.
	//
	// This member is required.
	ProjectId *string

	// The maximum number of results to be returned per paginated request. Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListDashboardsOutput struct {

	// A list that summarizes each dashboard in the project.
	//
	// This member is required.
	DashboardSummaries []types.DashboardSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDashboardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDashboards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDashboards{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addEndpointPrefix_opListDashboardsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListDashboardsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDashboards(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListDashboardsMiddleware struct {
}

func (*endpointPrefix_opListDashboardsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListDashboardsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "monitor." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListDashboardsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListDashboardsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListDashboardsAPIClient is a client that implements the ListDashboards
// operation.
type ListDashboardsAPIClient interface {
	ListDashboards(context.Context, *ListDashboardsInput, ...func(*Options)) (*ListDashboardsOutput, error)
}

var _ ListDashboardsAPIClient = (*Client)(nil)

// ListDashboardsPaginatorOptions is the paginator options for ListDashboards
type ListDashboardsPaginatorOptions struct {
	// The maximum number of results to be returned per paginated request. Default: 50
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDashboardsPaginator is a paginator for ListDashboards
type ListDashboardsPaginator struct {
	options   ListDashboardsPaginatorOptions
	client    ListDashboardsAPIClient
	params    *ListDashboardsInput
	nextToken *string
	firstPage bool
}

// NewListDashboardsPaginator returns a new ListDashboardsPaginator
func NewListDashboardsPaginator(client ListDashboardsAPIClient, params *ListDashboardsInput, optFns ...func(*ListDashboardsPaginatorOptions)) *ListDashboardsPaginator {
	options := ListDashboardsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDashboardsInput{}
	}

	return &ListDashboardsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDashboardsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDashboards page.
func (p *ListDashboardsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDashboardsOutput, error) {
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

	result, err := p.client.ListDashboards(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDashboards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListDashboards",
	}
}
