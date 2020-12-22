// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// After you deploy a model into production using Amazon SageMaker hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint. For an overview of Amazon SageMaker, see How
// It Works (https://docs.aws.amazon.com/sagemaker/latest/dg/how-it-works.html).
// Amazon SageMaker strips all POST headers except those supported by the API.
// Amazon SageMaker might add additional headers. You should not rely on the
// behavior of headers outside those enumerated in the request syntax. Calls to
// InvokeEndpoint are authenticated by using AWS Signature Version 4. For
// information, see Authenticating Requests (AWS Signature Version 4)
// (http://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon S3 API Reference. A customer's model containers must respond to
// requests within 60 seconds. The model itself can have a maximum processing time
// of 60 seconds before responding to the /invocations. If your model is going to
// take 50-60 seconds of processing time, the SDK socket timeout should be set to
// be 70 seconds. Endpoints are scoped to an individual account, and are not
// public. The URL does not contain the account ID, but Amazon SageMaker determines
// the account ID from the authentication token that is supplied by the caller.
func (c *Client) InvokeEndpoint(ctx context.Context, params *InvokeEndpointInput, optFns ...func(*Options)) (*InvokeEndpointOutput, error) {
	if params == nil {
		params = &InvokeEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeEndpoint", params, optFns, addOperationInvokeEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeEndpointInput struct {

	// Provides input data, in the format specified in the ContentType request header.
	// Amazon SageMaker passes all of the data in the body to the model. For
	// information about the format of the request body, see Common Data
	// Formats-Inference
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// This member is required.
	Body []byte

	// The name of the endpoint that you specified when you created the endpoint using
	// the CreateEndpoint
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpoint.html) API.
	//
	// This member is required.
	EndpointName *string

	// The desired MIME type of the inference in the response.
	Accept *string

	// The MIME type of the input data in the request body.
	ContentType *string

	// Provides additional information about a request for an inference submitted to a
	// model hosted at an Amazon SageMaker endpoint. The information is an opaque value
	// that is forwarded verbatim. You could use this value, for example, to provide an
	// ID that you can use to track a request or to provide other metadata that a
	// service endpoint was programmed to process. The value must consist of no more
	// than 1024 visible US-ASCII characters as specified in Section 3.3.6. Field Value
	// Components (https://tools.ietf.org/html/rfc7230#section-3.2.6) of the Hypertext
	// Transfer Protocol (HTTP/1.1). This feature is currently supported in the AWS
	// SDKs but not in the Amazon SageMaker Python SDK.
	CustomAttributes *string

	// The model to request for inference when invoking a multi-model endpoint.
	TargetModel *string

	// Specify the production variant to send the inference request to when invoking an
	// endpoint that is running two or more variants. Note that this parameter
	// overrides the default behavior for the endpoint, which is to distribute the
	// invocation traffic based on the variant weights.
	TargetVariant *string
}

type InvokeEndpointOutput struct {

	// Includes the inference provided by the model. For information about the format
	// of the response body, see Common Data Formats-Inference
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-inference.html).
	//
	// This member is required.
	Body []byte

	// The MIME type of the inference returned in the response body.
	ContentType *string

	// Provides additional information in the response about the inference returned by
	// a model hosted at an Amazon SageMaker endpoint. The information is an opaque
	// value that is forwarded verbatim. You could use this value, for example, to
	// return an ID received in the CustomAttributes header of a request or other
	// metadata that a service endpoint was programmed to produce. The value must
	// consist of no more than 1024 visible US-ASCII characters as specified in Section
	// 3.3.6. Field Value Components
	// (https://tools.ietf.org/html/rfc7230#section-3.2.6) of the Hypertext Transfer
	// Protocol (HTTP/1.1). If the customer wants the custom attribute returned, the
	// model must set the custom attribute to be included on the way back. This feature
	// is currently supported in the AWS SDKs but not in the Amazon SageMaker Python
	// SDK.
	CustomAttributes *string

	// Identifies the production variant that was invoked.
	InvokedProductionVariant *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInvokeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeEndpoint{}, middleware.After)
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
	if err = addOpInvokeEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "InvokeEndpoint",
	}
}
