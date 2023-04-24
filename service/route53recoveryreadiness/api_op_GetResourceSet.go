// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays the details about a resource set, including a list of the resources in
// the set.
func (c *Client) GetResourceSet(ctx context.Context, params *GetResourceSetInput, optFns ...func(*Options)) (*GetResourceSetOutput, error) {
	if params == nil {
		params = &GetResourceSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceSet", params, optFns, c.addOperationGetResourceSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceSetInput struct {

	// Name of a resource set.
	//
	// This member is required.
	ResourceSetName *string

	noSmithyDocumentSerde
}

type GetResourceSetOutput struct {

	// The Amazon Resource Name (ARN) for the resource set.
	ResourceSetArn *string

	// The name of the resource set.
	ResourceSetName *string

	// The resource type of the resources in the resource set. Enter one of the
	// following values for resource type: AWS::ApiGateway::Stage,
	// AWS::ApiGatewayV2::Stage, AWS::AutoScaling::AutoScalingGroup,
	// AWS::CloudWatch::Alarm, AWS::EC2::CustomerGateway, AWS::DynamoDB::Table,
	// AWS::EC2::Volume, AWS::ElasticLoadBalancing::LoadBalancer,
	// AWS::ElasticLoadBalancingV2::LoadBalancer, AWS::Lambda::Function,
	// AWS::MSK::Cluster, AWS::RDS::DBCluster, AWS::Route53::HealthCheck,
	// AWS::SQS::Queue, AWS::SNS::Topic, AWS::SNS::Subscription, AWS::EC2::VPC,
	// AWS::EC2::VPNConnection, AWS::EC2::VPNGateway,
	// AWS::Route53RecoveryReadiness::DNSTargetResource
	ResourceSetType *string

	// A list of resource objects.
	Resources []types.Resource

	// A collection of tags associated with a resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceSet{}, middleware.After)
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
	if err = addOpGetResourceSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResourceSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "GetResourceSet",
	}
}
