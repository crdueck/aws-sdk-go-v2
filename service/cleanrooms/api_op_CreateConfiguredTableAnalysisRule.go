// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new analysis rule for a configured table. Currently, only one
// analysis rule can be created for a given configured table.
func (c *Client) CreateConfiguredTableAnalysisRule(ctx context.Context, params *CreateConfiguredTableAnalysisRuleInput, optFns ...func(*Options)) (*CreateConfiguredTableAnalysisRuleOutput, error) {
	if params == nil {
		params = &CreateConfiguredTableAnalysisRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfiguredTableAnalysisRule", params, optFns, c.addOperationCreateConfiguredTableAnalysisRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfiguredTableAnalysisRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfiguredTableAnalysisRuleInput struct {

	// The entire created configured table analysis rule object.
	//
	// This member is required.
	AnalysisRulePolicy types.ConfiguredTableAnalysisRulePolicy

	// The type of analysis rule. Valid values are AGGREGATION and LIST.
	//
	// This member is required.
	AnalysisRuleType types.ConfiguredTableAnalysisRuleType

	// The identifier for the configured table to create the analysis rule for.
	// Currently accepts the configured table ID.
	//
	// This member is required.
	ConfiguredTableIdentifier *string

	noSmithyDocumentSerde
}

type CreateConfiguredTableAnalysisRuleOutput struct {

	// The entire created analysis rule.
	//
	// This member is required.
	AnalysisRule *types.ConfiguredTableAnalysisRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfiguredTableAnalysisRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfiguredTableAnalysisRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfiguredTableAnalysisRule{}, middleware.After)
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
	if err = addOpCreateConfiguredTableAnalysisRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfiguredTableAnalysisRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfiguredTableAnalysisRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "CreateConfiguredTableAnalysisRule",
	}
}
