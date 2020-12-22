// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of AWS Batch jobs. You must specify only one of the following:
//
// *
// a job queue ID to return a list of jobs in that job queue
//
// * a multi-node
// parallel job ID to return a list of that job's nodes
//
// * an array job ID to
// return a list of that job's children
//
// You can filter the results by job status
// with the jobStatus parameter. If you do not specify a status, only RUNNING jobs
// are returned.
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if params == nil {
		params = &ListJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobs", params, optFns, addOperationListJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobsInput struct {

	// The job ID for an array job. Specifying an array job ID with this parameter
	// lists all child jobs from within the specified array.
	ArrayJobId *string

	// The name or full Amazon Resource Name (ARN) of the job queue with which to list
	// jobs.
	JobQueue *string

	// The job status with which to filter jobs in the specified queue. If you do not
	// specify a status, only RUNNING jobs are returned.
	JobStatus types.JobStatus

	// The maximum number of results returned by ListJobs in paginated output. When
	// this parameter is used, ListJobs only returns maxResults results in a single
	// page along with a nextToken response element. The remaining results of the
	// initial request can be seen by sending another ListJobs request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// is not used, then ListJobs returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults int32

	// The job ID for a multi-node parallel job. Specifying a multi-node parallel job
	// ID with this parameter lists all nodes that are associated with the specified
	// job.
	MultiNodeJobId *string

	// The nextToken value returned from a previous paginated ListJobs request where
	// maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string
}

type ListJobsOutput struct {

	// A list of job summaries that match the request.
	//
	// This member is required.
	JobSummaryList []types.JobSummary

	// The nextToken value to include in a future ListJobs request. When the results of
	// a ListJobs request exceed maxResults, this value can be used to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before); err != nil {
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

// ListJobsAPIClient is a client that implements the ListJobs operation.
type ListJobsAPIClient interface {
	ListJobs(context.Context, *ListJobsInput, ...func(*Options)) (*ListJobsOutput, error)
}

var _ ListJobsAPIClient = (*Client)(nil)

// ListJobsPaginatorOptions is the paginator options for ListJobs
type ListJobsPaginatorOptions struct {
	// The maximum number of results returned by ListJobs in paginated output. When
	// this parameter is used, ListJobs only returns maxResults results in a single
	// page along with a nextToken response element. The remaining results of the
	// initial request can be seen by sending another ListJobs request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// is not used, then ListJobs returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobsPaginator is a paginator for ListJobs
type ListJobsPaginator struct {
	options   ListJobsPaginatorOptions
	client    ListJobsAPIClient
	params    *ListJobsInput
	nextToken *string
	firstPage bool
}

// NewListJobsPaginator returns a new ListJobsPaginator
func NewListJobsPaginator(client ListJobsAPIClient, params *ListJobsInput, optFns ...func(*ListJobsPaginatorOptions)) *ListJobsPaginator {
	options := ListJobsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListJobsInput{}
	}

	return &ListJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListJobs page.
func (p *ListJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "ListJobs",
	}
}
