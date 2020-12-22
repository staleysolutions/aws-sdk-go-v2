// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the emergency calling configuration details from the specified Amazon
// Chime Voice Connector.
func (c *Client) DeleteVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *DeleteVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*DeleteVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	if params == nil {
		params = &DeleteVoiceConnectorEmergencyCallingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVoiceConnectorEmergencyCallingConfiguration", params, optFns, addOperationDeleteVoiceConnectorEmergencyCallingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVoiceConnectorEmergencyCallingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVoiceConnectorEmergencyCallingConfigurationInput struct {

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type DeleteVoiceConnectorEmergencyCallingConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteVoiceConnectorEmergencyCallingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
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
	if err = addOpDeleteVoiceConnectorEmergencyCallingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVoiceConnectorEmergencyCallingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVoiceConnectorEmergencyCallingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteVoiceConnectorEmergencyCallingConfiguration",
	}
}
