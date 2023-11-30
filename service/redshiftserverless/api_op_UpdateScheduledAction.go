// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a scheduled action.
func (c *Client) UpdateScheduledAction(ctx context.Context, params *UpdateScheduledActionInput, optFns ...func(*Options)) (*UpdateScheduledActionOutput, error) {
	if params == nil {
		params = &UpdateScheduledActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateScheduledAction", params, optFns, c.addOperationUpdateScheduledActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateScheduledActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateScheduledActionInput struct {

	// The name of the scheduled action to update to.
	//
	// This member is required.
	ScheduledActionName *string

	// Specifies whether to enable the scheduled action.
	Enabled *bool

	// The end time in UTC of the scheduled action to update.
	EndTime *time.Time

	// The ARN of the IAM role to assume to run the scheduled action. This IAM role
	// must have permission to run the Amazon Redshift Serverless API operation in the
	// scheduled action. This IAM role must allow the Amazon Redshift scheduler to
	// schedule creating snapshots (Principal scheduler.redshift.amazonaws.com) to
	// assume permissions on your behalf. For more information about the IAM role to
	// use with the Amazon Redshift scheduler, see Using Identity-Based Policies for
	// Amazon Redshift (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html)
	// in the Amazon Redshift Cluster Management Guide
	RoleArn *string

	// The schedule for a one-time (at format) or recurring (cron format) scheduled
	// action. Schedule invocations must be separated by at least one hour. Format of
	// at expressions is " at(yyyy-mm-ddThh:mm:ss) ". For example, "
	// at(2016-03-04T17:27:00) ". Format of cron expressions is " cron(Minutes Hours
	// Day-of-month Month Day-of-week Year) ". For example, " cron(0 10 ? * MON *) ".
	// For more information, see Cron Expressions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions)
	// in the Amazon CloudWatch Events User Guide.
	Schedule types.Schedule

	// The descripion of the scheduled action to update to.
	ScheduledActionDescription *string

	// The start time in UTC of the scheduled action to update to.
	StartTime *time.Time

	// A JSON format string of the Amazon Redshift Serverless API operation with input
	// parameters. The following is an example of a target action. "{"CreateSnapshot":
	// {"NamespaceName": "sampleNamespace","SnapshotName": "sampleSnapshot",
	// "retentionPeriod": "1"}}"
	TargetAction types.TargetAction

	noSmithyDocumentSerde
}

type UpdateScheduledActionOutput struct {

	// The ScheduledAction object that was updated.
	ScheduledAction *types.ScheduledActionResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateScheduledActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateScheduledAction"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateScheduledActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateScheduledAction(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateScheduledAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateScheduledAction",
	}
}
