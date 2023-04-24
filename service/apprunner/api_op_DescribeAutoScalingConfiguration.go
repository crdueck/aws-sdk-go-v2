// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return a full description of an App Runner automatic scaling configuration
// resource.
func (c *Client) DescribeAutoScalingConfiguration(ctx context.Context, params *DescribeAutoScalingConfigurationInput, optFns ...func(*Options)) (*DescribeAutoScalingConfigurationOutput, error) {
	if params == nil {
		params = &DescribeAutoScalingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutoScalingConfiguration", params, optFns, c.addOperationDescribeAutoScalingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutoScalingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoScalingConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the App Runner auto scaling configuration
	// that you want a description for. The ARN can be a full auto scaling
	// configuration ARN, or a partial ARN ending with either .../name  or
	// .../name/revision . If a revision isn't specified, the latest active revision
	// is described.
	//
	// This member is required.
	AutoScalingConfigurationArn *string

	noSmithyDocumentSerde
}

type DescribeAutoScalingConfigurationOutput struct {

	// A full description of the App Runner auto scaling configuration that you
	// specified in this request.
	//
	// This member is required.
	AutoScalingConfiguration *types.AutoScalingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAutoScalingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeAutoScalingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeAutoScalingConfiguration{}, middleware.After)
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
	if err = addOpDescribeAutoScalingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoScalingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAutoScalingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "DescribeAutoScalingConfiguration",
	}
}
