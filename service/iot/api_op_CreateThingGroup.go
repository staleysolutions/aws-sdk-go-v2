// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a thing group. This is a control plane operation. See Authorization
// (https://docs.aws.amazon.com/iot/latest/developerguide/iot-authorization.html)
// for information about authorizing control plane actions.
func (c *Client) CreateThingGroup(ctx context.Context, params *CreateThingGroupInput, optFns ...func(*Options)) (*CreateThingGroupOutput, error) {
	if params == nil {
		params = &CreateThingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateThingGroup", params, optFns, addOperationCreateThingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateThingGroupInput struct {

	// The thing group name to create.
	//
	// This member is required.
	ThingGroupName *string

	// The name of the parent thing group.
	ParentGroupName *string

	// Metadata which can be used to manage the thing group.
	Tags []types.Tag

	// The thing group properties.
	ThingGroupProperties *types.ThingGroupProperties
}

type CreateThingGroupOutput struct {

	// The thing group ARN.
	ThingGroupArn *string

	// The thing group ID.
	ThingGroupId *string

	// The thing group name.
	ThingGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateThingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateThingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThingGroup{}, middleware.After)
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
	if err = addOpCreateThingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateThingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateThingGroup",
	}
}
