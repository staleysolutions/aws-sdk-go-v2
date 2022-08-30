// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the label groups.
func (c *Client) ListLabelGroups(ctx context.Context, params *ListLabelGroupsInput, optFns ...func(*Options)) (*ListLabelGroupsOutput, error) {
	if params == nil {
		params = &ListLabelGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLabelGroups", params, optFns, c.addOperationListLabelGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLabelGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLabelGroupsInput struct {

	// The beginning of the name of the label groups to be listed.
	LabelGroupNameBeginsWith *string

	// Specifies the maximum number of label groups to list.
	MaxResults *int32

	// An opaque pagination token indicating where to continue the listing of label
	// groups.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLabelGroupsOutput struct {

	// A summary of the label groups.
	LabelGroupSummaries []types.LabelGroupSummary

	// An opaque pagination token indicating where to continue the listing of label
	// groups.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLabelGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListLabelGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListLabelGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLabelGroups(options.Region), middleware.Before); err != nil {
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

// ListLabelGroupsAPIClient is a client that implements the ListLabelGroups
// operation.
type ListLabelGroupsAPIClient interface {
	ListLabelGroups(context.Context, *ListLabelGroupsInput, ...func(*Options)) (*ListLabelGroupsOutput, error)
}

var _ ListLabelGroupsAPIClient = (*Client)(nil)

// ListLabelGroupsPaginatorOptions is the paginator options for ListLabelGroups
type ListLabelGroupsPaginatorOptions struct {
	// Specifies the maximum number of label groups to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLabelGroupsPaginator is a paginator for ListLabelGroups
type ListLabelGroupsPaginator struct {
	options   ListLabelGroupsPaginatorOptions
	client    ListLabelGroupsAPIClient
	params    *ListLabelGroupsInput
	nextToken *string
	firstPage bool
}

// NewListLabelGroupsPaginator returns a new ListLabelGroupsPaginator
func NewListLabelGroupsPaginator(client ListLabelGroupsAPIClient, params *ListLabelGroupsInput, optFns ...func(*ListLabelGroupsPaginatorOptions)) *ListLabelGroupsPaginator {
	if params == nil {
		params = &ListLabelGroupsInput{}
	}

	options := ListLabelGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLabelGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLabelGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLabelGroups page.
func (p *ListLabelGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLabelGroupsOutput, error) {
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

	result, err := p.client.ListLabelGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLabelGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "ListLabelGroups",
	}
}
