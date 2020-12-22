// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of all approved origins associated with the instance.
func (c *Client) ListApprovedOrigins(ctx context.Context, params *ListApprovedOriginsInput, optFns ...func(*Options)) (*ListApprovedOriginsOutput, error) {
	if params == nil {
		params = &ListApprovedOriginsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApprovedOrigins", params, optFns, addOperationListApprovedOriginsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApprovedOriginsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApprovedOriginsInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximimum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
}

type ListApprovedOriginsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The approved origins.
	Origins []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListApprovedOriginsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApprovedOrigins{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApprovedOrigins{}, middleware.After)
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
	if err = addOpListApprovedOriginsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApprovedOrigins(options.Region), middleware.Before); err != nil {
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

// ListApprovedOriginsAPIClient is a client that implements the ListApprovedOrigins
// operation.
type ListApprovedOriginsAPIClient interface {
	ListApprovedOrigins(context.Context, *ListApprovedOriginsInput, ...func(*Options)) (*ListApprovedOriginsOutput, error)
}

var _ ListApprovedOriginsAPIClient = (*Client)(nil)

// ListApprovedOriginsPaginatorOptions is the paginator options for
// ListApprovedOrigins
type ListApprovedOriginsPaginatorOptions struct {
	// The maximimum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApprovedOriginsPaginator is a paginator for ListApprovedOrigins
type ListApprovedOriginsPaginator struct {
	options   ListApprovedOriginsPaginatorOptions
	client    ListApprovedOriginsAPIClient
	params    *ListApprovedOriginsInput
	nextToken *string
	firstPage bool
}

// NewListApprovedOriginsPaginator returns a new ListApprovedOriginsPaginator
func NewListApprovedOriginsPaginator(client ListApprovedOriginsAPIClient, params *ListApprovedOriginsInput, optFns ...func(*ListApprovedOriginsPaginatorOptions)) *ListApprovedOriginsPaginator {
	options := ListApprovedOriginsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListApprovedOriginsInput{}
	}

	return &ListApprovedOriginsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApprovedOriginsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListApprovedOrigins page.
func (p *ListApprovedOriginsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApprovedOriginsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListApprovedOrigins(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApprovedOrigins(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListApprovedOrigins",
	}
}
