// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// InferSNOMEDCT detects possible medical concepts as entities and links them to
// codes from the Systematized Nomenclature of Medicine, Clinical Terms (SNOMED-CT)
// ontology
func (c *Client) InferSNOMEDCT(ctx context.Context, params *InferSNOMEDCTInput, optFns ...func(*Options)) (*InferSNOMEDCTOutput, error) {
	if params == nil {
		params = &InferSNOMEDCTInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InferSNOMEDCT", params, optFns, c.addOperationInferSNOMEDCTMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InferSNOMEDCTOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InferSNOMEDCTInput struct {

	// The input text to be analyzed using InferSNOMEDCT. The text should be a string
	// with 1 to 10000 characters.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type InferSNOMEDCTOutput struct {

	// The collection of medical concept entities extracted from the input text and
	// their associated information. For each entity, the response provides the entity
	// text, the entity category, where the entity text begins and ends, and the level
	// of confidence that Comprehend Medical has in the detection and analysis.
	// Attributes and traits of the entity are also returned.
	//
	// This member is required.
	Entities []types.SNOMEDCTEntity

	// The number of characters in the input request documentation.
	Characters *types.Characters

	// The version of the model used to analyze the documents, in the format n.n.n You
	// can use this information to track the model used for a particular batch of
	// documents.
	ModelVersion *string

	// If the result of the request is truncated, the pagination token can be used to
	// fetch the next page of entities.
	PaginationToken *string

	// The details of the SNOMED-CT revision, including the edition, language, and
	// version date.
	SNOMEDCTDetails *types.SNOMEDCTDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInferSNOMEDCTMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpInferSNOMEDCT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpInferSNOMEDCT{}, middleware.After)
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
	if err = addOpInferSNOMEDCTValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInferSNOMEDCT(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInferSNOMEDCT(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "InferSNOMEDCT",
	}
}
