// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists AWS Lambda layers
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) and
// shows information about the latest version of each. Specify a runtime identifier
// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) to list only
// layers that indicate that they're compatible with that runtime.
func (c *Client) ListLayers(ctx context.Context, params *ListLayersInput, optFns ...func(*Options)) (*ListLayersOutput, error) {
	if params == nil {
		params = &ListLayersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLayers", params, optFns, addOperationListLayersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLayersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLayersInput struct {

	// A runtime identifier. For example, go1.x.
	CompatibleRuntime types.Runtime

	// A pagination token returned by a previous call.
	Marker *string

	// The maximum number of layers to return.
	MaxItems *int32
}

type ListLayersOutput struct {

	// A list of function layers.
	Layers []types.LayersListItem

	// A pagination token returned when the response doesn't contain all layers.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLayersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLayers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLayers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLayers(options.Region), middleware.Before); err != nil {
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

// ListLayersAPIClient is a client that implements the ListLayers operation.
type ListLayersAPIClient interface {
	ListLayers(context.Context, *ListLayersInput, ...func(*Options)) (*ListLayersOutput, error)
}

var _ ListLayersAPIClient = (*Client)(nil)

// ListLayersPaginatorOptions is the paginator options for ListLayers
type ListLayersPaginatorOptions struct {
	// The maximum number of layers to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLayersPaginator is a paginator for ListLayers
type ListLayersPaginator struct {
	options   ListLayersPaginatorOptions
	client    ListLayersAPIClient
	params    *ListLayersInput
	nextToken *string
	firstPage bool
}

// NewListLayersPaginator returns a new ListLayersPaginator
func NewListLayersPaginator(client ListLayersAPIClient, params *ListLayersInput, optFns ...func(*ListLayersPaginatorOptions)) *ListLayersPaginator {
	options := ListLayersPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListLayersInput{}
	}

	return &ListLayersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLayersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListLayers page.
func (p *ListLayersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLayersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListLayers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListLayers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListLayers",
	}
}
