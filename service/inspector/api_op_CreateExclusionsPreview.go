// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts the generation of an exclusions preview for the specified assessment
// template. The exclusions preview lists the potential exclusions
// (ExclusionPreview) that Inspector can detect before it runs the assessment.
func (c *Client) CreateExclusionsPreview(ctx context.Context, params *CreateExclusionsPreviewInput, optFns ...func(*Options)) (*CreateExclusionsPreviewOutput, error) {
	if params == nil {
		params = &CreateExclusionsPreviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExclusionsPreview", params, optFns, c.addOperationCreateExclusionsPreviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExclusionsPreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExclusionsPreviewInput struct {

	// The ARN that specifies the assessment template for which you want to create an
	// exclusions preview.
	//
	// This member is required.
	AssessmentTemplateArn *string

	noSmithyDocumentSerde
}

type CreateExclusionsPreviewOutput struct {

	// Specifies the unique identifier of the requested exclusions preview. You can
	// use the unique identifier to retrieve the exclusions preview when running the
	// GetExclusionsPreview API.
	//
	// This member is required.
	PreviewToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExclusionsPreviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateExclusionsPreview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateExclusionsPreview{}, middleware.After)
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
	if err = addOpCreateExclusionsPreviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExclusionsPreview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateExclusionsPreview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "CreateExclusionsPreview",
	}
}
