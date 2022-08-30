// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Specifies information about the inference scheduler being used, including name,
// model, status, and associated metadata
func (c *Client) DescribeInferenceScheduler(ctx context.Context, params *DescribeInferenceSchedulerInput, optFns ...func(*Options)) (*DescribeInferenceSchedulerOutput, error) {
	if params == nil {
		params = &DescribeInferenceSchedulerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInferenceScheduler", params, optFns, c.addOperationDescribeInferenceSchedulerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInferenceSchedulerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInferenceSchedulerInput struct {

	// The name of the inference scheduler being described.
	//
	// This member is required.
	InferenceSchedulerName *string

	noSmithyDocumentSerde
}

type DescribeInferenceSchedulerOutput struct {

	// Specifies the time at which the inference scheduler was created.
	CreatedAt *time.Time

	// A period of time (in minutes) by which inference on the data is delayed after
	// the data starts. For instance, if you select an offset delay time of five
	// minutes, inference will not begin on the data until the first data measurement
	// after the five minute mark. For example, if five minutes is selected, the
	// inference scheduler will wake up at the configured frequency with the additional
	// five minute delay time to check the customer S3 bucket. The customer can upload
	// data at the same frequency and they don't need to stop and restart the scheduler
	// when uploading new data.
	DataDelayOffsetInMinutes *int64

	// Specifies configuration information for the input data for the inference
	// scheduler, including delimiter, format, and dataset location.
	DataInputConfiguration *types.InferenceInputConfiguration

	// Specifies information for the output results for the inference scheduler,
	// including the output S3 location.
	DataOutputConfiguration *types.InferenceOutputConfiguration

	// Specifies how often data is uploaded to the source S3 bucket for the input data.
	// This value is the length of time between data uploads. For instance, if you
	// select 5 minutes, Amazon Lookout for Equipment will upload the real-time data to
	// the source bucket once every 5 minutes. This frequency also determines how often
	// Amazon Lookout for Equipment starts a scheduled inference on your data. In this
	// example, it starts once every 5 minutes.
	DataUploadFrequency types.DataUploadFrequency

	// The Amazon Resource Name (ARN) of the inference scheduler being described.
	InferenceSchedulerArn *string

	// The name of the inference scheduler being described.
	InferenceSchedulerName *string

	// Indicates whether the latest execution for the inference scheduler was Anomalous
	// (anomalous events found) or Normal (no anomalous events found).
	LatestInferenceResult types.LatestInferenceResult

	// The Amazon Resource Name (ARN) of the ML model of the inference scheduler being
	// described.
	ModelArn *string

	// The name of the ML model of the inference scheduler being described.
	ModelName *string

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source for the inference scheduler being described.
	RoleArn *string

	// Provides the identifier of the KMS key used to encrypt inference scheduler data
	// by Amazon Lookout for Equipment.
	ServerSideKmsKeyId *string

	// Indicates the status of the inference scheduler.
	Status types.InferenceSchedulerStatus

	// Specifies the time at which the inference scheduler was last updated, if it was.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInferenceSchedulerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeInferenceScheduler{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeInferenceScheduler{}, middleware.After)
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
	if err = addOpDescribeInferenceSchedulerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInferenceScheduler(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInferenceScheduler(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "DescribeInferenceScheduler",
	}
}
