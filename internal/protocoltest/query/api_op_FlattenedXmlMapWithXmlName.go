// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Flattened maps with @xmlName
func (c *Client) FlattenedXmlMapWithXmlName(ctx context.Context, params *FlattenedXmlMapWithXmlNameInput, optFns ...func(*Options)) (*FlattenedXmlMapWithXmlNameOutput, error) {
	if params == nil {
		params = &FlattenedXmlMapWithXmlNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FlattenedXmlMapWithXmlName", params, optFns, addOperationFlattenedXmlMapWithXmlNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FlattenedXmlMapWithXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FlattenedXmlMapWithXmlNameInput struct {
}

type FlattenedXmlMapWithXmlNameOutput struct {
	MyMap map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationFlattenedXmlMapWithXmlNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFlattenedXmlMapWithXmlName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFlattenedXmlMapWithXmlName{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFlattenedXmlMapWithXmlName(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFlattenedXmlMapWithXmlName(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "FlattenedXmlMapWithXmlName",
	}
}
