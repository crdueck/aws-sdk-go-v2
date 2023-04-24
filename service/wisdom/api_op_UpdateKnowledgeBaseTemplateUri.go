// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the template URI of a knowledge base. This is only supported for
// knowledge bases of type EXTERNAL. Include a single variable in ${variable}
// format; this interpolated by Wisdom using ingested content. For example, if you
// ingest a Salesforce article, it has an Id value, and you can set the template
// URI to
// https://myInstanceName.lightning.force.com/lightning/r/Knowledge__kav/*${Id}*/view
// .
func (c *Client) UpdateKnowledgeBaseTemplateUri(ctx context.Context, params *UpdateKnowledgeBaseTemplateUriInput, optFns ...func(*Options)) (*UpdateKnowledgeBaseTemplateUriOutput, error) {
	if params == nil {
		params = &UpdateKnowledgeBaseTemplateUriInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKnowledgeBaseTemplateUri", params, optFns, c.addOperationUpdateKnowledgeBaseTemplateUriMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKnowledgeBaseTemplateUriOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKnowledgeBaseTemplateUriInput struct {

	// The identifier of the knowledge base. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	// The template URI to update.
	//
	// This member is required.
	TemplateUri *string

	noSmithyDocumentSerde
}

type UpdateKnowledgeBaseTemplateUriOutput struct {

	// The knowledge base to update.
	KnowledgeBase *types.KnowledgeBaseData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateKnowledgeBaseTemplateUriMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateKnowledgeBaseTemplateUri{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateKnowledgeBaseTemplateUri{}, middleware.After)
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
	if err = addOpUpdateKnowledgeBaseTemplateUriValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKnowledgeBaseTemplateUri(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateKnowledgeBaseTemplateUri(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "UpdateKnowledgeBaseTemplateUri",
	}
}
