// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the current runtime configuration for the specified fleet, which tells
// Amazon GameLift how to launch server processes on instances in the fleet. You
// can update a fleet's runtime configuration at any time after the fleet is
// created; it does not need to be in an ACTIVE status. To update runtime
// configuration, specify the fleet ID and provide a RuntimeConfiguration object
// with an updated set of server process configurations. Each instance in a Amazon
// GameLift fleet checks regularly for an updated runtime configuration and changes
// how it launches server processes to comply with the latest version. Existing
// server processes are not affected by the update; runtime configuration changes
// are applied gradually as existing processes shut down and new processes are
// launched during Amazon GameLift's normal process recycling activity. Learn more
// Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
// * CreateFleet
//
// * ListFleets
//
// * DeleteFleet
//
// *
// DescribeFleetAttributes
//
// * Update fleets:
//
// * UpdateFleetAttributes
//
// *
// UpdateFleetCapacity
//
// * UpdateFleetPortSettings
//
// * UpdateRuntimeConfiguration
//
// *
// StartFleetActions or StopFleetActions
func (c *Client) UpdateRuntimeConfiguration(ctx context.Context, params *UpdateRuntimeConfigurationInput, optFns ...func(*Options)) (*UpdateRuntimeConfigurationOutput, error) {
	if params == nil {
		params = &UpdateRuntimeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRuntimeConfiguration", params, optFns, addOperationUpdateRuntimeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRuntimeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type UpdateRuntimeConfigurationInput struct {

	// A unique identifier for a fleet to update runtime configuration for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// Instructions for launching server processes on each instance in the fleet.
	// Server processes run either a custom game build executable or a Realtime Servers
	// script. The runtime configuration lists the types of server processes to run on
	// an instance and includes the following configuration settings: the server
	// executable or launch script file, launch parameters, and the number of processes
	// to run concurrently on each instance. A CreateFleet request must include a
	// runtime configuration with at least one server process configuration.
	//
	// This member is required.
	RuntimeConfiguration *types.RuntimeConfiguration
}

// Represents the returned data in response to a request operation.
type UpdateRuntimeConfigurationOutput struct {

	// The runtime configuration currently in force. If the update was successful, this
	// object matches the one in the request.
	RuntimeConfiguration *types.RuntimeConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRuntimeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRuntimeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRuntimeConfiguration{}, middleware.After)
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
	if err = addOpUpdateRuntimeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRuntimeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRuntimeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateRuntimeConfiguration",
	}
}
