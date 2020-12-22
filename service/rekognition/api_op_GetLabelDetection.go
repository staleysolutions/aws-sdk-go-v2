// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the label detection results of a Amazon Rekognition Video analysis started
// by StartLabelDetection. The label detection operation is started by a call to
// StartLabelDetection which returns a job identifier (JobId). When the label
// detection operation finishes, Amazon Rekognition publishes a completion status
// to the Amazon Simple Notification Service topic registered in the initial call
// to StartlabelDetection. To get the results of the label detection operation,
// first check that the status value published to the Amazon SNS topic is
// SUCCEEDED. If so, call GetLabelDetection and pass the job identifier (JobId)
// from the initial call to StartLabelDetection. GetLabelDetection returns an array
// of detected labels (Labels) sorted by the time the labels were detected. You can
// also sort by the label name by specifying NAME for the SortBy input parameter.
// The labels returned include the label name, the percentage confidence in the
// accuracy of the detected label, and the time the label was detected in the
// video. The returned labels also include bounding box information for common
// objects, a hierarchical taxonomy of detected labels, and the version of the
// label model used for detection. Use MaxResults parameter to limit the number of
// labels returned. If there are more results than specified in MaxResults, the
// value of NextToken in the operation response contains a pagination token for
// getting the next set of results. To get the next page of results, call
// GetlabelDetection and populate the NextToken request parameter with the token
// value returned from the previous call to GetLabelDetection.
func (c *Client) GetLabelDetection(ctx context.Context, params *GetLabelDetectionInput, optFns ...func(*Options)) (*GetLabelDetectionOutput, error) {
	if params == nil {
		params = &GetLabelDetectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLabelDetection", params, optFns, addOperationGetLabelDetectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLabelDetectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLabelDetectionInput struct {

	// Job identifier for the label detection operation for which you want results
	// returned. You get the job identifer from an initial call to StartlabelDetection.
	//
	// This member is required.
	JobId *string

	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more labels to
	// retrieve), Amazon Rekognition Video returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of labels.
	NextToken *string

	// Sort to use for elements in the Labels array. Use TIMESTAMP to sort array
	// elements by the time labels are detected. Use NAME to alphabetically group
	// elements for a label together. Within each label group, the array element are
	// sorted by detection confidence. The default sort is by TIMESTAMP.
	SortBy types.LabelDetectionSortBy
}

type GetLabelDetectionOutput struct {

	// The current status of the label detection job.
	JobStatus types.VideoJobStatus

	// Version number of the label detection model that was used to detect labels.
	LabelModelVersion *string

	// An array of labels detected in the video. Each element contains the detected
	// label and the time, in milliseconds from the start of the video, that the label
	// was detected.
	Labels []types.LabelDetection

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of labels.
	NextToken *string

	// If the job fails, StatusMessage provides a descriptive error message.
	StatusMessage *string

	// Information about a video that Amazon Rekognition Video analyzed. Videometadata
	// is returned in every page of paginated responses from a Amazon Rekognition video
	// operation.
	VideoMetadata *types.VideoMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLabelDetectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLabelDetection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLabelDetection{}, middleware.After)
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
	if err = addOpGetLabelDetectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLabelDetection(options.Region), middleware.Before); err != nil {
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

// GetLabelDetectionAPIClient is a client that implements the GetLabelDetection
// operation.
type GetLabelDetectionAPIClient interface {
	GetLabelDetection(context.Context, *GetLabelDetectionInput, ...func(*Options)) (*GetLabelDetectionOutput, error)
}

var _ GetLabelDetectionAPIClient = (*Client)(nil)

// GetLabelDetectionPaginatorOptions is the paginator options for GetLabelDetection
type GetLabelDetectionPaginatorOptions struct {
	// Maximum number of results to return per paginated call. The largest value you
	// can specify is 1000. If you specify a value greater than 1000, a maximum of 1000
	// results is returned. The default value is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetLabelDetectionPaginator is a paginator for GetLabelDetection
type GetLabelDetectionPaginator struct {
	options   GetLabelDetectionPaginatorOptions
	client    GetLabelDetectionAPIClient
	params    *GetLabelDetectionInput
	nextToken *string
	firstPage bool
}

// NewGetLabelDetectionPaginator returns a new GetLabelDetectionPaginator
func NewGetLabelDetectionPaginator(client GetLabelDetectionAPIClient, params *GetLabelDetectionInput, optFns ...func(*GetLabelDetectionPaginatorOptions)) *GetLabelDetectionPaginator {
	options := GetLabelDetectionPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetLabelDetectionInput{}
	}

	return &GetLabelDetectionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetLabelDetectionPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetLabelDetection page.
func (p *GetLabelDetectionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetLabelDetectionOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetLabelDetection(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetLabelDetection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "GetLabelDetection",
	}
}
