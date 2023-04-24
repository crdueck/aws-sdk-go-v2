// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an access policy for the specified container to restrict the users and
// clients that can access it. For information about the data that is included in
// an access policy, see the AWS Identity and Access Management User Guide (https://aws.amazon.com/documentation/iam/)
// . For this release of the REST API, you can create only one policy for a
// container. If you enter PutContainerPolicy twice, the second command modifies
// the existing policy.
func (c *Client) PutContainerPolicy(ctx context.Context, params *PutContainerPolicyInput, optFns ...func(*Options)) (*PutContainerPolicyOutput, error) {
	if params == nil {
		params = &PutContainerPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutContainerPolicy", params, optFns, c.addOperationPutContainerPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutContainerPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutContainerPolicyInput struct {

	// The name of the container.
	//
	// This member is required.
	ContainerName *string

	// The contents of the policy, which includes the following:
	//   - One Version tag
	//   - One Statement tag that contains the standard tags for the policy.
	//
	// This member is required.
	Policy *string

	noSmithyDocumentSerde
}

type PutContainerPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutContainerPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutContainerPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutContainerPolicy{}, middleware.After)
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
	if err = addOpPutContainerPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutContainerPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutContainerPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediastore",
		OperationName: "PutContainerPolicy",
	}
}
