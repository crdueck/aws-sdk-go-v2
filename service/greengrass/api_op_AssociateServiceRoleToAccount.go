// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a role with your account. AWS IoT Greengrass will use the role to
// access your Lambda functions and AWS IoT resources. This is necessary for
// deployments to succeed. The role must have at least minimum permissions in the
// policy ”AWSGreengrassResourceAccessRolePolicy”.
func (c *Client) AssociateServiceRoleToAccount(ctx context.Context, params *AssociateServiceRoleToAccountInput, optFns ...func(*Options)) (*AssociateServiceRoleToAccountOutput, error) {
	if params == nil {
		params = &AssociateServiceRoleToAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateServiceRoleToAccount", params, optFns, c.addOperationAssociateServiceRoleToAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateServiceRoleToAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateServiceRoleToAccountInput struct {

	// The ARN of the service role you wish to associate with your account.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type AssociateServiceRoleToAccountOutput struct {

	// The time when the service role was associated with the account.
	AssociatedAt *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateServiceRoleToAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateServiceRoleToAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateServiceRoleToAccount{}, middleware.After)
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
	if err = addOpAssociateServiceRoleToAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateServiceRoleToAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateServiceRoleToAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "AssociateServiceRoleToAccount",
	}
}
