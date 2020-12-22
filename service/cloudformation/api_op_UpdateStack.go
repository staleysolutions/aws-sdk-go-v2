// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a stack as specified in the template. After the call completes
// successfully, the stack update starts. You can check the status of the stack via
// the DescribeStacks action. To get a copy of the template for an existing stack,
// you can use the GetTemplate action. For more information about creating an
// update template, updating a stack, and monitoring the progress of the update,
// see Updating a Stack
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html).
func (c *Client) UpdateStack(ctx context.Context, params *UpdateStackInput, optFns ...func(*Options)) (*UpdateStackOutput, error) {
	if params == nil {
		params = &UpdateStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStack", params, optFns, addOperationUpdateStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for an UpdateStack action.
type UpdateStackInput struct {

	// The name or unique stack ID of the stack to update.
	//
	// This member is required.
	StackName *string

	// In some cases, you must explicitly acknowledge that your stack template contains
	// certain capabilities in order for AWS CloudFormation to update the stack.
	//
	// *
	// CAPABILITY_IAM and CAPABILITY_NAMED_IAM Some stack templates might include
	// resources that can affect permissions in your AWS account; for example, by
	// creating new AWS Identity and Access Management (IAM) users. For those stacks,
	// you must explicitly acknowledge this by specifying one of these capabilities.
	// The following IAM resources require you to specify either the CAPABILITY_IAM or
	// CAPABILITY_NAMED_IAM capability.
	//
	// * If you have IAM resources, you can specify
	// either capability.
	//
	// * If you have IAM resources with custom names, you must
	// specify CAPABILITY_NAMED_IAM.
	//
	// * If you don't specify either of these
	// capabilities, AWS CloudFormation returns an InsufficientCapabilities error.
	//
	// If
	// your stack template contains these resources, we recommend that you review all
	// permissions associated with them and edit their permissions if necessary.
	//
	// *
	// AWS::IAM::AccessKey
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html)
	//
	// *
	// AWS::IAM::Group
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html)
	//
	// *
	// AWS::IAM::InstanceProfile
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)
	//
	// *
	// AWS::IAM::Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html)
	//
	// *
	// AWS::IAM::Role
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html)
	//
	// *
	// AWS::IAM::User
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html)
	//
	// *
	// AWS::IAM::UserToGroupAddition
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html)
	//
	// For
	// more information, see Acknowledging IAM Resources in AWS CloudFormation
	// Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	//
	// *
	// CAPABILITY_AUTO_EXPAND Some template contain macros. Macros perform custom
	// processing on templates; this can include simple actions like find-and-replace
	// operations, all the way to extensive transformations of entire templates.
	// Because of this, users typically create a change set from the processed
	// template, so that they can review the changes resulting from the macros before
	// actually updating the stack. If your stack template contains one or more macros,
	// and you choose to update a stack directly from the processed template, without
	// first reviewing the resulting changes in a change set, you must acknowledge this
	// capability. This includes the AWS::Include
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html)
	// and AWS::Serverless
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html)
	// transforms, which are macros hosted by AWS CloudFormation. If you want to update
	// a stack from a stack template that contains macros and nested stacks, you must
	// update the stack directly from the template using this capability. You should
	// only update stacks directly from a stack template that contains macros if you
	// know what processing the macro performs. Each macro relies on an underlying
	// Lambda service function for processing stack templates. Be aware that the Lambda
	// function owner can update the function operation without AWS CloudFormation
	// being notified. For more information, see Using AWS CloudFormation Macros to
	// Perform Custom Processing on Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html).
	Capabilities []types.Capability

	// A unique identifier for this UpdateStack request. Specify this token if you plan
	// to retry requests so that AWS CloudFormation knows that you're not attempting to
	// update a stack with the same name. You might retry UpdateStack requests to
	// ensure that AWS CloudFormation successfully received them. All events triggered
	// by a given stack operation are assigned the same client request token, which you
	// can use to track operations. For example, if you execute a CreateStack operation
	// with the token token1, then all the StackEvents generated by that operation will
	// have ClientRequestToken set as token1. In the console, stack operations display
	// the client request token on the Events tab. Stack operations that are initiated
	// from the console use the token format Console-StackOperation-ID, which helps you
	// easily identify the stack operation . For example, if you create a stack using
	// the console, each stack event would be assigned the same token in the following
	// format: Console-CreateStack-7f59c3cf-00d2-40c7-b2ff-e75db0987002.
	ClientRequestToken *string

	// Amazon Simple Notification Service topic Amazon Resource Names (ARNs) that AWS
	// CloudFormation associates with the stack. Specify an empty list to remove all
	// notification topics.
	NotificationARNs []string

	// A list of Parameter structures that specify input parameters for the stack. For
	// more information, see the Parameter
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_Parameter.html)
	// data type.
	Parameters []types.Parameter

	// The template resource types that you have permissions to work with for this
	// update stack action, such as AWS::EC2::Instance, AWS::EC2::*, or
	// Custom::MyCustomInstance. If the list of resource types doesn't include a
	// resource that you're updating, the stack update fails. By default, AWS
	// CloudFormation grants permissions to all resource types. AWS Identity and Access
	// Management (IAM) uses this parameter for AWS CloudFormation-specific condition
	// keys in IAM policies. For more information, see Controlling Access with AWS
	// Identity and Access Management
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html).
	ResourceTypes []string

	// The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM)
	// role that AWS CloudFormation assumes to update the stack. AWS CloudFormation
	// uses the role's credentials to make calls on your behalf. AWS CloudFormation
	// always uses this role for all future operations on the stack. As long as users
	// have permission to operate on the stack, AWS CloudFormation uses this role even
	// if the users don't have permission to pass it. Ensure that the role grants least
	// privilege. If you don't specify a value, AWS CloudFormation uses the role that
	// was previously associated with the stack. If no role is available, AWS
	// CloudFormation uses a temporary session that is generated from your user
	// credentials.
	RoleARN *string

	// The rollback triggers for AWS CloudFormation to monitor during stack creation
	// and updating operations, and for the specified monitoring period afterwards.
	RollbackConfiguration *types.RollbackConfiguration

	// Structure containing a new stack policy body. You can specify either the
	// StackPolicyBody or the StackPolicyURL parameter, but not both. You might update
	// the stack policy, for example, in order to protect a new resource that you
	// created during a stack update. If you do not specify a stack policy, the current
	// policy that is associated with the stack is unchanged.
	StackPolicyBody *string

	// Structure containing the temporary overriding stack policy body. You can specify
	// either the StackPolicyDuringUpdateBody or the StackPolicyDuringUpdateURL
	// parameter, but not both. If you want to update protected resources, specify a
	// temporary overriding stack policy during this update. If you do not specify a
	// stack policy, the current policy that is associated with the stack will be used.
	StackPolicyDuringUpdateBody *string

	// Location of a file containing the temporary overriding stack policy. The URL
	// must point to a policy (max size: 16KB) located in an S3 bucket in the same
	// Region as the stack. You can specify either the StackPolicyDuringUpdateBody or
	// the StackPolicyDuringUpdateURL parameter, but not both. If you want to update
	// protected resources, specify a temporary overriding stack policy during this
	// update. If you do not specify a stack policy, the current policy that is
	// associated with the stack will be used.
	StackPolicyDuringUpdateURL *string

	// Location of a file containing the updated stack policy. The URL must point to a
	// policy (max size: 16KB) located in an S3 bucket in the same Region as the stack.
	// You can specify either the StackPolicyBody or the StackPolicyURL parameter, but
	// not both. You might update the stack policy, for example, in order to protect a
	// new resource that you created during a stack update. If you do not specify a
	// stack policy, the current policy that is associated with the stack is unchanged.
	StackPolicyURL *string

	// Key-value pairs to associate with this stack. AWS CloudFormation also propagates
	// these tags to supported resources in the stack. You can specify a maximum number
	// of 50 tags. If you don't specify this parameter, AWS CloudFormation doesn't
	// modify the stack's tags. If you specify an empty value, AWS CloudFormation
	// removes all associated tags.
	Tags []types.Tag

	// Structure containing the template body with a minimum length of 1 byte and a
	// maximum length of 51,200 bytes. (For more information, go to Template Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.) Conditional: You must specify only one of
	// the following parameters: TemplateBody, TemplateURL, or set the
	// UsePreviousTemplate to true.
	TemplateBody *string

	// Location of file containing the template body. The URL must point to a template
	// that is located in an Amazon S3 bucket. For more information, go to Template
	// Anatomy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide. Conditional: You must specify only one of
	// the following parameters: TemplateBody, TemplateURL, or set the
	// UsePreviousTemplate to true.
	TemplateURL *string

	// Reuse the existing template that is associated with the stack that you are
	// updating. Conditional: You must specify only one of the following parameters:
	// TemplateBody, TemplateURL, or set the UsePreviousTemplate to true.
	UsePreviousTemplate *bool
}

// The output for an UpdateStack action.
type UpdateStackOutput struct {

	// Unique identifier of the stack.
	StackId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateStack{}, middleware.After)
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
	if err = addOpUpdateStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "UpdateStack",
	}
}
