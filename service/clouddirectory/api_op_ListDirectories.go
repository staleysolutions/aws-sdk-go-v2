// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists directories created within an account.
func (c *Client) ListDirectories(ctx context.Context, params *ListDirectoriesInput, optFns ...func(*Options)) (*ListDirectoriesOutput, error) {
	if params == nil {
		params = &ListDirectoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDirectories", params, optFns, addOperationListDirectoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDirectoriesInput struct {

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// The state of the directories in the list. Can be either Enabled, Disabled, or
	// Deleted.
	State types.DirectoryState
}

type ListDirectoriesOutput struct {

	// Lists all directories that are associated with your account in pagination
	// fashion.
	//
	// This member is required.
	Directories []types.Directory

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDirectoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDirectories{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDirectories(options.Region), middleware.Before); err != nil {
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

// ListDirectoriesAPIClient is a client that implements the ListDirectories
// operation.
type ListDirectoriesAPIClient interface {
	ListDirectories(context.Context, *ListDirectoriesInput, ...func(*Options)) (*ListDirectoriesOutput, error)
}

var _ ListDirectoriesAPIClient = (*Client)(nil)

// ListDirectoriesPaginatorOptions is the paginator options for ListDirectories
type ListDirectoriesPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDirectoriesPaginator is a paginator for ListDirectories
type ListDirectoriesPaginator struct {
	options   ListDirectoriesPaginatorOptions
	client    ListDirectoriesAPIClient
	params    *ListDirectoriesInput
	nextToken *string
	firstPage bool
}

// NewListDirectoriesPaginator returns a new ListDirectoriesPaginator
func NewListDirectoriesPaginator(client ListDirectoriesAPIClient, params *ListDirectoriesInput, optFns ...func(*ListDirectoriesPaginatorOptions)) *ListDirectoriesPaginator {
	options := ListDirectoriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListDirectoriesInput{}
	}

	return &ListDirectoriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDirectoriesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDirectories page.
func (p *ListDirectoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDirectoriesOutput, error) {
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

	result, err := p.client.ListDirectories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDirectories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListDirectories",
	}
}
