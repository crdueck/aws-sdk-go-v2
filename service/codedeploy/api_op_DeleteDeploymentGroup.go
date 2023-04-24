// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a deployment group.
func (c *Client) DeleteDeploymentGroup(ctx context.Context, params *DeleteDeploymentGroupInput, optFns ...func(*Options)) (*DeleteDeploymentGroupOutput, error) {
	if params == nil {
		params = &DeleteDeploymentGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDeploymentGroup", params, optFns, c.addOperationDeleteDeploymentGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDeploymentGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteDeploymentGroup operation.
type DeleteDeploymentGroupInput struct {

	// The name of an CodeDeploy application associated with the IAM user or Amazon
	// Web Services account.
	//
	// This member is required.
	ApplicationName *string

	// The name of a deployment group for the specified application.
	//
	// This member is required.
	DeploymentGroupName *string

	noSmithyDocumentSerde
}

// Represents the output of a DeleteDeploymentGroup operation.
type DeleteDeploymentGroupOutput struct {

	// If the output contains no data, and the corresponding deployment group
	// contained at least one Auto Scaling group, CodeDeploy successfully removed all
	// corresponding Auto Scaling lifecycle event hooks from the Amazon EC2 instances
	// in the Auto Scaling group. If the output contains data, CodeDeploy could not
	// remove some Auto Scaling lifecycle event hooks from the Amazon EC2 instances in
	// the Auto Scaling group.
	HooksNotCleanedUp []types.AutoScalingGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDeploymentGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDeploymentGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDeploymentGroup{}, middleware.After)
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
	if err = addOpDeleteDeploymentGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDeploymentGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDeploymentGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "DeleteDeploymentGroup",
	}
}
