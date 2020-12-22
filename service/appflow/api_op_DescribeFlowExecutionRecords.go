// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Fetches the execution history of the flow.
func (c *Client) DescribeFlowExecutionRecords(ctx context.Context, params *DescribeFlowExecutionRecordsInput, optFns ...func(*Options)) (*DescribeFlowExecutionRecordsOutput, error) {
	if params == nil {
		params = &DescribeFlowExecutionRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFlowExecutionRecords", params, optFns, addOperationDescribeFlowExecutionRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFlowExecutionRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowExecutionRecordsInput struct {

	// The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	//
	// This member is required.
	FlowName *string

	// Specifies the maximum number of items that should be returned in the result set.
	// The default for maxResults is 20 (for all paginated API operations).
	MaxResults *int32

	// The pagination token for the next page of data.
	NextToken *string
}

type DescribeFlowExecutionRecordsOutput struct {

	// Returns a list of all instances when this flow was run.
	FlowExecutions []types.ExecutionRecord

	// The pagination token for the next page of data.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFlowExecutionRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFlowExecutionRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFlowExecutionRecords{}, middleware.After)
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
	if err = addOpDescribeFlowExecutionRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlowExecutionRecords(options.Region), middleware.Before); err != nil {
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

// DescribeFlowExecutionRecordsAPIClient is a client that implements the
// DescribeFlowExecutionRecords operation.
type DescribeFlowExecutionRecordsAPIClient interface {
	DescribeFlowExecutionRecords(context.Context, *DescribeFlowExecutionRecordsInput, ...func(*Options)) (*DescribeFlowExecutionRecordsOutput, error)
}

var _ DescribeFlowExecutionRecordsAPIClient = (*Client)(nil)

// DescribeFlowExecutionRecordsPaginatorOptions is the paginator options for
// DescribeFlowExecutionRecords
type DescribeFlowExecutionRecordsPaginatorOptions struct {
	// Specifies the maximum number of items that should be returned in the result set.
	// The default for maxResults is 20 (for all paginated API operations).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFlowExecutionRecordsPaginator is a paginator for
// DescribeFlowExecutionRecords
type DescribeFlowExecutionRecordsPaginator struct {
	options   DescribeFlowExecutionRecordsPaginatorOptions
	client    DescribeFlowExecutionRecordsAPIClient
	params    *DescribeFlowExecutionRecordsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFlowExecutionRecordsPaginator returns a new
// DescribeFlowExecutionRecordsPaginator
func NewDescribeFlowExecutionRecordsPaginator(client DescribeFlowExecutionRecordsAPIClient, params *DescribeFlowExecutionRecordsInput, optFns ...func(*DescribeFlowExecutionRecordsPaginatorOptions)) *DescribeFlowExecutionRecordsPaginator {
	options := DescribeFlowExecutionRecordsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeFlowExecutionRecordsInput{}
	}

	return &DescribeFlowExecutionRecordsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFlowExecutionRecordsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeFlowExecutionRecords page.
func (p *DescribeFlowExecutionRecordsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFlowExecutionRecordsOutput, error) {
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

	result, err := p.client.DescribeFlowExecutionRecords(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFlowExecutionRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeFlowExecutionRecords",
	}
}
