// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about Aurora global database clusters. This API supports
// pagination. For more information on Amazon Aurora, see  What Is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. This action only applies to Aurora DB clusters.
func (c *Client) DescribeGlobalClusters(ctx context.Context, params *DescribeGlobalClustersInput, optFns ...func(*Options)) (*DescribeGlobalClustersOutput, error) {
	if params == nil {
		params = &DescribeGlobalClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGlobalClusters", params, optFns, addOperationDescribeGlobalClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGlobalClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalClustersInput struct {

	// A filter that specifies one or more global DB clusters to describe. Supported
	// filters:
	//
	// * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	// Resource Names (ARNs). The results list will only include information about the
	// DB clusters identified by these ARNs.
	Filters []types.Filter

	// The user-supplied DB cluster identifier. If this parameter is specified,
	// information from only the specific DB cluster is returned. This parameter isn't
	// case-sensitive. Constraints:
	//
	// * If supplied, must match an existing
	// DBClusterIdentifier.
	GlobalClusterIdentifier *string

	// An optional pagination token provided by a previous DescribeGlobalClusters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeGlobalClustersOutput struct {

	// The list of global clusters returned by this request.
	GlobalClusters []types.GlobalCluster

	// An optional pagination token provided by a previous DescribeGlobalClusters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGlobalClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeGlobalClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeGlobalClusters{}, middleware.After)
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
	if err = addOpDescribeGlobalClustersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalClusters(options.Region), middleware.Before); err != nil {
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

// DescribeGlobalClustersAPIClient is a client that implements the
// DescribeGlobalClusters operation.
type DescribeGlobalClustersAPIClient interface {
	DescribeGlobalClusters(context.Context, *DescribeGlobalClustersInput, ...func(*Options)) (*DescribeGlobalClustersOutput, error)
}

var _ DescribeGlobalClustersAPIClient = (*Client)(nil)

// DescribeGlobalClustersPaginatorOptions is the paginator options for
// DescribeGlobalClusters
type DescribeGlobalClustersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeGlobalClustersPaginator is a paginator for DescribeGlobalClusters
type DescribeGlobalClustersPaginator struct {
	options   DescribeGlobalClustersPaginatorOptions
	client    DescribeGlobalClustersAPIClient
	params    *DescribeGlobalClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeGlobalClustersPaginator returns a new DescribeGlobalClustersPaginator
func NewDescribeGlobalClustersPaginator(client DescribeGlobalClustersAPIClient, params *DescribeGlobalClustersInput, optFns ...func(*DescribeGlobalClustersPaginatorOptions)) *DescribeGlobalClustersPaginator {
	options := DescribeGlobalClustersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeGlobalClustersInput{}
	}

	return &DescribeGlobalClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeGlobalClustersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeGlobalClusters page.
func (p *DescribeGlobalClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeGlobalClustersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeGlobalClusters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeGlobalClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeGlobalClusters",
	}
}
