// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists a history of user activity and any risks detected as part of Amazon
// Cognito advanced security.
func (c *Client) AdminListUserAuthEvents(ctx context.Context, params *AdminListUserAuthEventsInput, optFns ...func(*Options)) (*AdminListUserAuthEventsOutput, error) {
	if params == nil {
		params = &AdminListUserAuthEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminListUserAuthEvents", params, optFns, addOperationAdminListUserAuthEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminListUserAuthEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminListUserAuthEventsInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user pool username or an alias.
	//
	// This member is required.
	Username *string

	// The maximum number of authentication events to return.
	MaxResults *int32

	// A pagination token.
	NextToken *string
}

type AdminListUserAuthEventsOutput struct {

	// The response object. It includes the EventID, EventType, CreationDate,
	// EventRisk, and EventResponse.
	AuthEvents []types.AuthEventType

	// A pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdminListUserAuthEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminListUserAuthEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminListUserAuthEvents{}, middleware.After)
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
	if err = addOpAdminListUserAuthEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminListUserAuthEvents(options.Region), middleware.Before); err != nil {
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

// AdminListUserAuthEventsAPIClient is a client that implements the
// AdminListUserAuthEvents operation.
type AdminListUserAuthEventsAPIClient interface {
	AdminListUserAuthEvents(context.Context, *AdminListUserAuthEventsInput, ...func(*Options)) (*AdminListUserAuthEventsOutput, error)
}

var _ AdminListUserAuthEventsAPIClient = (*Client)(nil)

// AdminListUserAuthEventsPaginatorOptions is the paginator options for
// AdminListUserAuthEvents
type AdminListUserAuthEventsPaginatorOptions struct {
	// The maximum number of authentication events to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// AdminListUserAuthEventsPaginator is a paginator for AdminListUserAuthEvents
type AdminListUserAuthEventsPaginator struct {
	options   AdminListUserAuthEventsPaginatorOptions
	client    AdminListUserAuthEventsAPIClient
	params    *AdminListUserAuthEventsInput
	nextToken *string
	firstPage bool
}

// NewAdminListUserAuthEventsPaginator returns a new
// AdminListUserAuthEventsPaginator
func NewAdminListUserAuthEventsPaginator(client AdminListUserAuthEventsAPIClient, params *AdminListUserAuthEventsInput, optFns ...func(*AdminListUserAuthEventsPaginatorOptions)) *AdminListUserAuthEventsPaginator {
	options := AdminListUserAuthEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &AdminListUserAuthEventsInput{}
	}

	return &AdminListUserAuthEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *AdminListUserAuthEventsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next AdminListUserAuthEvents page.
func (p *AdminListUserAuthEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*AdminListUserAuthEventsOutput, error) {
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

	result, err := p.client.AdminListUserAuthEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opAdminListUserAuthEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminListUserAuthEvents",
	}
}
