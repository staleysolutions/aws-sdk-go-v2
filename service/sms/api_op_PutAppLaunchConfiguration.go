// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the launch configuration for the specified application.
func (c *Client) PutAppLaunchConfiguration(ctx context.Context, params *PutAppLaunchConfigurationInput, optFns ...func(*Options)) (*PutAppLaunchConfigurationOutput, error) {
	if params == nil {
		params = &PutAppLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAppLaunchConfiguration", params, optFns, addOperationPutAppLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAppLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAppLaunchConfigurationInput struct {

	// The ID of the application.
	AppId *string

	// Indicates whether the application is configured to launch automatically after
	// replication is complete.
	AutoLaunch *bool

	// The name of service role in the customer's account that AWS CloudFormation uses
	// to launch the application.
	RoleName *string

	// Information about the launch configurations for server groups in the
	// application.
	ServerGroupLaunchConfigurations []types.ServerGroupLaunchConfiguration
}

type PutAppLaunchConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAppLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAppLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAppLaunchConfiguration{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAppLaunchConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAppLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "PutAppLaunchConfiguration",
	}
}
