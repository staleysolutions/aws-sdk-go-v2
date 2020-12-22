// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an Amazon SageMaker job.
func (c *Client) DescribeAutoMLJob(ctx context.Context, params *DescribeAutoMLJobInput, optFns ...func(*Options)) (*DescribeAutoMLJobOutput, error) {
	if params == nil {
		params = &DescribeAutoMLJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutoMLJob", params, optFns, addOperationDescribeAutoMLJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoMLJobInput struct {

	// Request information about a job using that job's unique name.
	//
	// This member is required.
	AutoMLJobName *string
}

type DescribeAutoMLJobOutput struct {

	// Returns the job's ARN.
	//
	// This member is required.
	AutoMLJobArn *string

	// Returns the name of a job.
	//
	// This member is required.
	AutoMLJobName *string

	// Returns the job's AutoMLJobSecondaryStatus.
	//
	// This member is required.
	AutoMLJobSecondaryStatus types.AutoMLJobSecondaryStatus

	// Returns the job's AutoMLJobStatus.
	//
	// This member is required.
	AutoMLJobStatus types.AutoMLJobStatus

	// Returns the job's creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// Returns the job's input data config.
	//
	// This member is required.
	InputDataConfig []types.AutoMLChannel

	// Returns the job's last modified time.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// Returns the job's output data config.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that has read permission to the input data location and write permission to
	// the output data location in Amazon S3.
	//
	// This member is required.
	RoleArn *string

	// Returns information on the job's artifacts found in AutoMLJobArtifacts.
	AutoMLJobArtifacts *types.AutoMLJobArtifacts

	// Returns the job's config.
	AutoMLJobConfig *types.AutoMLJobConfig

	// Returns the job's objective.
	AutoMLJobObjective *types.AutoMLJobObjective

	// Returns the job's BestCandidate.
	BestCandidate *types.AutoMLCandidate

	// Returns the job's end time.
	EndTime *time.Time

	// Returns the job's FailureReason.
	FailureReason *string

	// Returns the job's output from GenerateCandidateDefinitionsOnly.
	GenerateCandidateDefinitionsOnly bool

	// Returns the job's problem type.
	ProblemType types.ProblemType

	// This contains ProblemType, AutoMLJobObjective and CompletionCriteria. They're
	// auto-inferred values, if not provided by you. If you do provide them, then
	// they'll be the same as provided.
	ResolvedAttributes *types.ResolvedAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAutoMLJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutoMLJob{}, middleware.After)
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
	if err = addOpDescribeAutoMLJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoMLJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAutoMLJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeAutoMLJob",
	}
}
