// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Switches over a blue/green deployment. Before you switch over, production
// traffic is routed to the databases in the blue environment. After you switch
// over, production traffic is routed to the databases in the green environment.
// For more information, see Using Amazon RDS Blue/Green Deployments for database
// updates (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html)
// in the Amazon RDS User Guide and Using Amazon RDS Blue/Green Deployments for
// database updates (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html)
// in the Amazon Aurora User Guide.
func (c *Client) SwitchoverBlueGreenDeployment(ctx context.Context, params *SwitchoverBlueGreenDeploymentInput, optFns ...func(*Options)) (*SwitchoverBlueGreenDeploymentOutput, error) {
	if params == nil {
		params = &SwitchoverBlueGreenDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SwitchoverBlueGreenDeployment", params, optFns, c.addOperationSwitchoverBlueGreenDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SwitchoverBlueGreenDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SwitchoverBlueGreenDeploymentInput struct {

	// The blue/green deployment identifier. Constraints:
	//   - Must match an existing blue/green deployment identifier.
	//
	// This member is required.
	BlueGreenDeploymentIdentifier *string

	// The amount of time, in seconds, for the switchover to complete. The default is
	// 300. If the switchover takes longer than the specified duration, then any
	// changes are rolled back, and no changes are made to the environments.
	SwitchoverTimeout *int32

	noSmithyDocumentSerde
}

type SwitchoverBlueGreenDeploymentOutput struct {

	// Contains the details about a blue/green deployment. For more information, see
	// Using Amazon RDS Blue/Green Deployments for database updates (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/blue-green-deployments.html)
	// in the Amazon RDS User Guide and Using Amazon RDS Blue/Green Deployments for
	// database updates (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html)
	// in the Amazon Aurora User Guide.
	BlueGreenDeployment *types.BlueGreenDeployment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSwitchoverBlueGreenDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSwitchoverBlueGreenDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSwitchoverBlueGreenDeployment{}, middleware.After)
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
	if err = addOpSwitchoverBlueGreenDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSwitchoverBlueGreenDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSwitchoverBlueGreenDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "SwitchoverBlueGreenDeployment",
	}
}
