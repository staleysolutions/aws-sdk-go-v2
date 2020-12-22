// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a set of one or more game sessions. Request a specific game session or
// request all game sessions on a fleet. Alternatively, use SearchGameSessions to
// request a set of active game sessions that are filtered by certain criteria. To
// retrieve protection policy settings for game sessions, use
// DescribeGameSessionDetails. To get game sessions, specify one of the following:
// game session ID, fleet ID, or alias ID. You can filter this request by game
// session status. Use the pagination parameters to retrieve results as a set of
// sequential pages. If successful, a GameSession object is returned for each game
// session matching the request. Available in Amazon GameLift Local.
//
// *
// CreateGameSession
//
// * DescribeGameSessions
//
// * DescribeGameSessionDetails
//
// *
// SearchGameSessions
//
// * UpdateGameSession
//
// * GetGameSessionLogUrl
//
// * Game session
// placements
//
// * StartGameSessionPlacement
//
// * DescribeGameSessionPlacement
//
// *
// StopGameSessionPlacement
func (c *Client) DescribeGameSessions(ctx context.Context, params *DescribeGameSessionsInput, optFns ...func(*Options)) (*DescribeGameSessionsOutput, error) {
	if params == nil {
		params = &DescribeGameSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameSessions", params, optFns, addOperationDescribeGameSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeGameSessionsInput struct {

	// A unique identifier for an alias associated with the fleet to retrieve all game
	// sessions for. You can use either the alias ID or ARN value.
	AliasId *string

	// A unique identifier for a fleet to retrieve all game sessions for. You can use
	// either the fleet ID or ARN value.
	FleetId *string

	// A unique identifier for the game session to retrieve.
	GameSessionId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Game session status to filter results on. Possible game session statuses include
	// ACTIVE, TERMINATED, ACTIVATING, and TERMINATING (the last two are transitory).
	StatusFilter *string
}

// Represents the returned data in response to a request operation.
type DescribeGameSessionsOutput struct {

	// A collection of objects containing game session properties for each session
	// matching the request.
	GameSessions []types.GameSession

	// Token that indicates where to resume retrieving results on the next call to this
	// operation. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGameSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameSessions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameSessions(options.Region), middleware.Before); err != nil {
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

// DescribeGameSessionsAPIClient is a client that implements the
// DescribeGameSessions operation.
type DescribeGameSessionsAPIClient interface {
	DescribeGameSessions(context.Context, *DescribeGameSessionsInput, ...func(*Options)) (*DescribeGameSessionsOutput, error)
}

var _ DescribeGameSessionsAPIClient = (*Client)(nil)

// DescribeGameSessionsPaginatorOptions is the paginator options for
// DescribeGameSessions
type DescribeGameSessionsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeGameSessionsPaginator is a paginator for DescribeGameSessions
type DescribeGameSessionsPaginator struct {
	options   DescribeGameSessionsPaginatorOptions
	client    DescribeGameSessionsAPIClient
	params    *DescribeGameSessionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeGameSessionsPaginator returns a new DescribeGameSessionsPaginator
func NewDescribeGameSessionsPaginator(client DescribeGameSessionsAPIClient, params *DescribeGameSessionsInput, optFns ...func(*DescribeGameSessionsPaginatorOptions)) *DescribeGameSessionsPaginator {
	options := DescribeGameSessionsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeGameSessionsInput{}
	}

	return &DescribeGameSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeGameSessionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeGameSessions page.
func (p *DescribeGameSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeGameSessionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeGameSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeGameSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameSessions",
	}
}
