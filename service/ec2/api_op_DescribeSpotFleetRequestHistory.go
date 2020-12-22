// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the events for the specified Spot Fleet request during the specified
// time. Spot Fleet events are delayed by up to 30 seconds before they can be
// described. This ensures that you can query by the last evaluated time and not
// miss a recorded event. Spot Fleet events are available for 48 hours.
func (c *Client) DescribeSpotFleetRequestHistory(ctx context.Context, params *DescribeSpotFleetRequestHistoryInput, optFns ...func(*Options)) (*DescribeSpotFleetRequestHistoryOutput, error) {
	if params == nil {
		params = &DescribeSpotFleetRequestHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSpotFleetRequestHistory", params, optFns, addOperationDescribeSpotFleetRequestHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSpotFleetRequestHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotFleetRequestHistory.
type DescribeSpotFleetRequestHistoryInput struct {

	// The ID of the Spot Fleet request.
	//
	// This member is required.
	SpotFleetRequestId *string

	// The starting date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	//
	// This member is required.
	StartTime *time.Time

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The type of events to describe. By default, all events are described.
	EventType types.EventType

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults int32

	// The token for the next set of results.
	NextToken *string
}

// Contains the output of DescribeSpotFleetRequestHistory.
type DescribeSpotFleetRequestHistoryOutput struct {

	// Information about the events in the history of the Spot Fleet request.
	HistoryRecords []types.HistoryRecord

	// The last date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ). All records up to this time were retrieved. If nextToken
	// indicates that there are more results, this value is not present.
	LastEvaluatedTime *time.Time

	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The ID of the Spot Fleet request.
	SpotFleetRequestId *string

	// The starting date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSpotFleetRequestHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotFleetRequestHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotFleetRequestHistory{}, middleware.After)
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
	if err = addOpDescribeSpotFleetRequestHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotFleetRequestHistory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSpotFleetRequestHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotFleetRequestHistory",
	}
}
