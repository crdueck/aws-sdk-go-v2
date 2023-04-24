// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a DB subnet group. The specified database subnet group must not be
// associated with any DB instances.
func (c *Client) DeleteDBSubnetGroup(ctx context.Context, params *DeleteDBSubnetGroupInput, optFns ...func(*Options)) (*DeleteDBSubnetGroupOutput, error) {
	if params == nil {
		params = &DeleteDBSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBSubnetGroup", params, optFns, c.addOperationDeleteDBSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDBSubnetGroupInput struct {

	// The name of the database subnet group to delete. You can't delete the default
	// subnet group. Constraints: Constraints: Must match the name of an existing
	// DBSubnetGroup. Must not be default. Example: mySubnetgroup
	//
	// This member is required.
	DBSubnetGroupName *string

	noSmithyDocumentSerde
}

type DeleteDBSubnetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDBSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBSubnetGroup{}, middleware.After)
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
	if err = addOpDeleteDBSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBSubnetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBSubnetGroup",
	}
}
