// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify a setting for a DB cluster. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request.
func (c *Client) ModifyDBCluster(ctx context.Context, params *ModifyDBClusterInput, optFns ...func(*Options)) (*ModifyDBClusterOutput, error) {
	if params == nil {
		params = &ModifyDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBCluster", params, optFns, addOperationModifyDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBClusterInput struct {

	// The DB cluster identifier for the cluster being modified. This parameter is not
	// case-sensitive. Constraints:
	//
	// * Must match the identifier of an existing
	// DBCluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	// A value that specifies whether the modifications in this request and any pending
	// modifications are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the DB cluster. If this parameter is set
	// to false, changes to the DB cluster are applied during the next maintenance
	// window. The ApplyImmediately parameter only affects the NewDBClusterIdentifier
	// and MasterUserPassword values. If you set the ApplyImmediately parameter value
	// to false, then changes to the NewDBClusterIdentifier and MasterUserPassword
	// values are applied during the next maintenance window. All other changes are
	// applied immediately, regardless of the value of the ApplyImmediately parameter.
	// Default: false
	ApplyImmediately bool

	// The number of days for which automated backups are retained. You must specify a
	// minimum value of 1. Default: 1 Constraints:
	//
	// * Must be a value from 1 to 35
	BackupRetentionPeriod *int32

	// The configuration setting for the log types to be enabled for export to
	// CloudWatch Logs for a specific DB cluster.
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration

	// The name of the DB cluster parameter group to use for the DB cluster.
	DBClusterParameterGroupName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool

	// True to enable mapping of AWS Identity and Access Management (IAM) accounts to
	// database accounts, and otherwise false. Default: false
	EnableIAMDatabaseAuthentication *bool

	// The version number of the database engine to which you want to upgrade. Changing
	// this parameter results in an outage. The change is applied during the next
	// maintenance window unless the ApplyImmediately parameter is set to true. For a
	// list of valid engine versions, see Engine Releases for Amazon Neptune
	// (https://docs.aws.amazon.com/neptune/latest/userguide/engine-releases.html), or
	// call DescribeDBEngineVersions
	// (https://docs.aws.amazon.com/neptune/latest/userguide/api-other-apis.html#DescribeDBEngineVersions).
	EngineVersion *string

	// The new password for the master database user. This password can contain any
	// printable ASCII character except "/", """, or "@". Constraints: Must contain
	// from 8 to 41 characters.
	MasterUserPassword *string

	// The new DB cluster identifier for the DB cluster when renaming a DB cluster.
	// This value is stored as a lowercase string. Constraints:
	//
	// * Must contain from 1
	// to 63 letters, numbers, or hyphens
	//
	// * The first character must be a letter
	//
	// *
	// Cannot end with a hyphen or contain two consecutive hyphens
	//
	// Example:
	// my-cluster2
	NewDBClusterIdentifier *string

	// (Not supported by Neptune)
	OptionGroupName *string

	// The port number on which the DB cluster accepts connections. Constraints: Value
	// must be 1150-65535 Default: The same port as the original DB cluster.
	Port *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region. Constraints:
	//
	// * Must be in the format hh24:mi-hh24:mi.
	//
	// * Must be in
	// Universal Coordinated Time (UTC).
	//
	// * Must not conflict with the preferred
	// maintenance window.
	//
	// * Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each AWS
	// Region, occurring on a random day of the week. Valid Days: Mon, Tue, Wed, Thu,
	// Fri, Sat, Sun. Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// A list of VPC security groups that the DB cluster will belong to.
	VpcSecurityGroupIds []string
}

type ModifyDBClusterOutput struct {

	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters action.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBCluster{}, middleware.After)
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
	if err = addOpModifyDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBCluster",
	}
}
