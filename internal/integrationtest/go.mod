module github.com/aws/aws-sdk-go-v2/internal/integrationtest

require (
	github.com/aws/aws-sdk-go-v2 v0.24.1-0.20200924224914-965b6782bf3d
	github.com/aws/aws-sdk-go-v2/config v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/acm v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/apigateway v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/appstream v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/athena v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/autoscaling v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/batch v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/cloudformation v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/cloudfront v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/cloudsearch v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v0.0.0-20200925205449-f9c5cc24907e
	//github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v0.0.0-20200925205449-f9c5cc24907e
	//github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/docdb v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/dynamodb v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/ec2 v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/ecr v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/ecs v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/efs v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/elasticache v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/emr v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/eventbridge v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/firehose v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/gamelift v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/glacier v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/glue v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/health v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/iam v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/inspector v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/iot v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/kinesis v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/kms v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/lambda v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/lightsail v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/neptune v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/opsworks v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/pinpointemail v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/polly v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/rds v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/redshift v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/rekognition v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/route53 v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/route53domains v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/route53resolver v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/s3 v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/ses v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/sfn v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/shield v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/sms v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/snowball v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/sns v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/sqs v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/ssm v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/sts v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/support v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/waf v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/wafregional v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/wafv2 v0.0.0-20200925205449-f9c5cc24907e
	github.com/aws/aws-sdk-go-v2/service/workspaces v0.0.0-20200925205449-f9c5cc24907e
	github.com/awslabs/smithy-go v0.0.0-20200924210334-28773c6e7960
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)

go 1.15
