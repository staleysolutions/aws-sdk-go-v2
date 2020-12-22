// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves IDs and annotations for traces available for a specified time frame
// using an optional filter. To get the full traces, pass the trace IDs to
// BatchGetTraces. A filter expression can target traced requests that hit specific
// service nodes or edges, have errors, or come from a known user. For example, the
// following filter expression targets traces that pass through api.example.com:
// service("api.example.com") This filter expression finds traces that have an
// annotation named account with the value 12345: annotation.account = "12345" For
// a full list of indexed fields and keywords that you can use in filter
// expressions, see Using Filter Expressions
// (https://docs.aws.amazon.com/xray/latest/devguide/xray-console-filters.html) in
// the AWS X-Ray Developer Guide.
func (c *Client) GetTraceSummaries(ctx context.Context, params *GetTraceSummariesInput, optFns ...func(*Options)) (*GetTraceSummariesOutput, error) {
	if params == nil {
		params = &GetTraceSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTraceSummaries", params, optFns, addOperationGetTraceSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTraceSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTraceSummariesInput struct {

	// The end of the time frame for which to retrieve traces.
	//
	// This member is required.
	EndTime *time.Time

	// The start of the time frame for which to retrieve traces.
	//
	// This member is required.
	StartTime *time.Time

	// Specify a filter expression to retrieve trace summaries for services or requests
	// that meet certain requirements.
	FilterExpression *string

	// Specify the pagination token returned by a previous request to retrieve the next
	// page of results.
	NextToken *string

	// Set to true to get summaries for only a subset of available traces.
	Sampling *bool

	// A parameter to indicate whether to enable sampling on trace summaries. Input
	// parameters are Name and Value.
	SamplingStrategy *types.SamplingStrategy

	// A parameter to indicate whether to query trace summaries by TraceId or Event
	// time.
	TimeRangeType types.TimeRangeType
}

type GetTraceSummariesOutput struct {

	// The start time of this page of results.
	ApproximateTime *time.Time

	// If the requested time frame contained more than one page of results, you can use
	// this token to retrieve the next page. The first page contains the most recent
	// results, closest to the end of the time frame.
	NextToken *string

	// Trace IDs and annotations for traces that were found in the specified time
	// frame.
	TraceSummaries []types.TraceSummary

	// The total number of traces processed, including traces that did not match the
	// specified filter expression.
	TracesProcessedCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTraceSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTraceSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTraceSummaries{}, middleware.After)
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
	if err = addOpGetTraceSummariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTraceSummaries(options.Region), middleware.Before); err != nil {
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

// GetTraceSummariesAPIClient is a client that implements the GetTraceSummaries
// operation.
type GetTraceSummariesAPIClient interface {
	GetTraceSummaries(context.Context, *GetTraceSummariesInput, ...func(*Options)) (*GetTraceSummariesOutput, error)
}

var _ GetTraceSummariesAPIClient = (*Client)(nil)

// GetTraceSummariesPaginatorOptions is the paginator options for GetTraceSummaries
type GetTraceSummariesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTraceSummariesPaginator is a paginator for GetTraceSummaries
type GetTraceSummariesPaginator struct {
	options   GetTraceSummariesPaginatorOptions
	client    GetTraceSummariesAPIClient
	params    *GetTraceSummariesInput
	nextToken *string
	firstPage bool
}

// NewGetTraceSummariesPaginator returns a new GetTraceSummariesPaginator
func NewGetTraceSummariesPaginator(client GetTraceSummariesAPIClient, params *GetTraceSummariesInput, optFns ...func(*GetTraceSummariesPaginatorOptions)) *GetTraceSummariesPaginator {
	options := GetTraceSummariesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetTraceSummariesInput{}
	}

	return &GetTraceSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTraceSummariesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetTraceSummaries page.
func (p *GetTraceSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTraceSummariesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetTraceSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTraceSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "GetTraceSummaries",
	}
}
