// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing function definition in the Data Catalog.
func (c *Client) UpdateUserDefinedFunction(ctx context.Context, params *UpdateUserDefinedFunctionInput, optFns ...func(*Options)) (*UpdateUserDefinedFunctionOutput, error) {
	if params == nil {
		params = &UpdateUserDefinedFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserDefinedFunction", params, optFns, c.addOperationUpdateUserDefinedFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserDefinedFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserDefinedFunctionInput struct {

	// The name of the catalog database where the function to be updated is located.
	//
	// This member is required.
	DatabaseName *string

	// A FunctionInput object that redefines the function in the Data Catalog.
	//
	// This member is required.
	FunctionInput *types.UserDefinedFunctionInput

	// The name of the function.
	//
	// This member is required.
	FunctionName *string

	// The ID of the Data Catalog where the function to be updated is located. If none
	// is provided, the Amazon Web Services account ID is used by default.
	CatalogId *string

	noSmithyDocumentSerde
}

type UpdateUserDefinedFunctionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserDefinedFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserDefinedFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserDefinedFunction{}, middleware.After)
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
	if err = addOpUpdateUserDefinedFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserDefinedFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserDefinedFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateUserDefinedFunction",
	}
}
