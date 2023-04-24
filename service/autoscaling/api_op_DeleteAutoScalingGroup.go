// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified Auto Scaling group. If the group has instances or scaling
// activities in progress, you must specify the option to force the deletion in
// order for it to succeed. The force delete operation will also terminate the EC2
// instances. If the group has a warm pool, the force delete option also deletes
// the warm pool. To remove instances from the Auto Scaling group before deleting
// it, call the DetachInstances API with the list of instances and the option to
// decrement the desired capacity. This ensures that Amazon EC2 Auto Scaling does
// not launch replacement instances. To terminate all instances before deleting the
// Auto Scaling group, call the UpdateAutoScalingGroup API and set the minimum
// size and desired capacity of the Auto Scaling group to zero. If the group has
// scaling policies, deleting the group deletes the policies, the underlying alarm
// actions, and any alarm that no longer has an associated action. For more
// information, see Delete your Auto Scaling infrastructure (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-process-shutdown.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DeleteAutoScalingGroup(ctx context.Context, params *DeleteAutoScalingGroupInput, optFns ...func(*Options)) (*DeleteAutoScalingGroupOutput, error) {
	if params == nil {
		params = &DeleteAutoScalingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAutoScalingGroup", params, optFns, c.addOperationDeleteAutoScalingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAutoScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAutoScalingGroupInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Specifies that the group is to be deleted along with all instances associated
	// with the group, without waiting for all instances to be terminated. This action
	// also deletes any outstanding lifecycle actions associated with the group.
	ForceDelete *bool

	noSmithyDocumentSerde
}

type DeleteAutoScalingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAutoScalingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAutoScalingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAutoScalingGroup{}, middleware.After)
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
	if err = addOpDeleteAutoScalingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAutoScalingGroup(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDeleteAutoScalingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DeleteAutoScalingGroup",
	}
}
