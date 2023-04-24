// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation is retired. Calling this operation does not change AQUA
// configuration. Amazon Redshift automatically determines whether to use AQUA
// (Advanced Query Accelerator).
func (c *Client) ModifyAquaConfiguration(ctx context.Context, params *ModifyAquaConfigurationInput, optFns ...func(*Options)) (*ModifyAquaConfigurationOutput, error) {
	if params == nil {
		params = &ModifyAquaConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyAquaConfiguration", params, optFns, c.addOperationModifyAquaConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyAquaConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyAquaConfigurationInput struct {

	// The identifier of the cluster to be modified.
	//
	// This member is required.
	ClusterIdentifier *string

	// This parameter is retired. Amazon Redshift automatically determines whether to
	// use AQUA (Advanced Query Accelerator).
	AquaConfigurationStatus types.AquaConfigurationStatus

	noSmithyDocumentSerde
}

type ModifyAquaConfigurationOutput struct {

	// This parameter is retired. Amazon Redshift automatically determines whether to
	// use AQUA (Advanced Query Accelerator).
	AquaConfiguration *types.AquaConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyAquaConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyAquaConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyAquaConfiguration{}, middleware.After)
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
	if err = addOpModifyAquaConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyAquaConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyAquaConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyAquaConfiguration",
	}
}
