// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the metadata for a notebook.
func (c *Client) UpdateNotebookMetadata(ctx context.Context, params *UpdateNotebookMetadataInput, optFns ...func(*Options)) (*UpdateNotebookMetadataOutput, error) {
	if params == nil {
		params = &UpdateNotebookMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNotebookMetadata", params, optFns, c.addOperationUpdateNotebookMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNotebookMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotebookMetadataInput struct {

	// The name to update the notebook to.
	//
	// This member is required.
	Name *string

	// The ID of the notebook to update the metadata for.
	//
	// This member is required.
	NotebookId *string

	// A unique case-sensitive string used to ensure the request to create the
	// notebook is idempotent (executes only once). This token is listed as not
	// required because Amazon Web Services SDKs (for example the Amazon Web Services
	// SDK for Java) auto-generate the token for you. If you are not using the Amazon
	// Web Services SDK or the Amazon Web Services CLI, you must provide this token or
	// the action will fail.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type UpdateNotebookMetadataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNotebookMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNotebookMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNotebookMetadata{}, middleware.After)
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
	if err = addOpUpdateNotebookMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotebookMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNotebookMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "UpdateNotebookMetadata",
	}
}
