// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a web ACL with a regional application resource, to protect the
// resource. A regional application can be an Application Load Balancer (ALB), an
// Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito user
// pool, or an App Runner service. For Amazon CloudFront, don't use this call.
// Instead, use your CloudFront distribution configuration. To associate a web ACL,
// in the CloudFront call UpdateDistribution , set the web ACL ID to the Amazon
// Resource Name (ARN) of the web ACL. For information, see UpdateDistribution (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html)
// in the Amazon CloudFront Developer Guide. When you make changes to web ACLs or
// web ACL components, like rules and rule groups, WAF propagates the changes
// everywhere that the web ACL and its components are stored and used. Your changes
// are applied within seconds, but there might be a brief period of inconsistency
// when the changes have arrived in some places and not in others. So, for example,
// if you change a rule action setting, the action might be the old action in one
// area and the new action in another area. Or if you add an IP address to an IP
// set used in a blocking rule, the new address might briefly be blocked in one
// area while still allowed in another. This temporary inconsistency can occur when
// you first associate a web ACL with an Amazon Web Services resource and when you
// change a web ACL that is already associated with a resource. Generally, any
// inconsistencies of this type last only a few seconds.
func (c *Client) AssociateWebACL(ctx context.Context, params *AssociateWebACLInput, optFns ...func(*Options)) (*AssociateWebACLOutput, error) {
	if params == nil {
		params = &AssociateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWebACL", params, optFns, c.addOperationAssociateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWebACLInput struct {

	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL.
	// The ARN must be in one of the following formats:
	//   - For an Application Load Balancer:
	//   arn:partition:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//   - For an Amazon API Gateway REST API:
	//   arn:partition:apigateway:region::/restapis/api-id/stages/stage-name
	//   - For an AppSync GraphQL API:
	//   arn:partition:appsync:region:account-id:apis/GraphQLApiId
	//   - For an Amazon Cognito user pool:
	//   arn:partition:cognito-idp:region:account-id:userpool/user-pool-id
	//   - For an App Runner service:
	//   arn:partition:apprunner:region:account-id:service/apprunner-service-name/apprunner-service-id
	//
	// This member is required.
	ResourceArn *string

	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with
	// the resource.
	//
	// This member is required.
	WebACLArn *string

	noSmithyDocumentSerde
}

type AssociateWebACLOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateWebACL{}, middleware.After)
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
	if err = addOpAssociateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "AssociateWebACL",
	}
}
