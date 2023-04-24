// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon OpenSearch Service-managed VPC endpoint.
func (c *Client) CreateVpcEndpoint(ctx context.Context, params *CreateVpcEndpointInput, optFns ...func(*Options)) (*CreateVpcEndpointOutput, error) {
	if params == nil {
		params = &CreateVpcEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcEndpoint", params, optFns, c.addOperationCreateVpcEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the CreateVpcEndpointRequest operation.
type CreateVpcEndpointInput struct {

	// The Amazon Resource Name (ARN) of the domain to grant access to.
	//
	// This member is required.
	DomainArn *string

	// Options to specify the subnets and security groups for the endpoint.
	//
	// This member is required.
	VpcOptions *types.VPCOptions

	// Unique, case-sensitive identifier to ensure idempotency of the request.
	ClientToken *string

	noSmithyDocumentSerde
}

// Container for response parameters to the CreateVpcEndpoint operation. Contains
// the configuration and status of the VPC Endpoint being created.
type CreateVpcEndpointOutput struct {

	// Information about the newly created VPC endpoint.
	//
	// This member is required.
	VpcEndpoint *types.VpcEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpcEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVpcEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVpcEndpoint{}, middleware.After)
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
	if err = addOpCreateVpcEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpcEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreateVpcEndpoint",
	}
}
