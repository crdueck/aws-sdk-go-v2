// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a log pattern to a LogPatternSet .
func (c *Client) UpdateLogPattern(ctx context.Context, params *UpdateLogPatternInput, optFns ...func(*Options)) (*UpdateLogPatternOutput, error) {
	if params == nil {
		params = &UpdateLogPatternInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLogPattern", params, optFns, c.addOperationUpdateLogPatternMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLogPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLogPatternInput struct {

	// The name of the log pattern.
	//
	// This member is required.
	PatternName *string

	// The name of the log pattern set.
	//
	// This member is required.
	PatternSetName *string

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string

	// The log pattern. The pattern must be DFA compatible. Patterns that utilize
	// forward lookahead or backreference constructions are not supported.
	Pattern *string

	// Rank of the log pattern. Must be a value between 1 and 1,000,000 . The patterns
	// are sorted by rank, so we recommend that you set your highest priority patterns
	// with the lowest rank. A pattern of rank 1 will be the first to get matched to a
	// log line. A pattern of rank 1,000,000 will be last to get matched. When you
	// configure custom log patterns from the console, a Low severity pattern
	// translates to a 750,000 rank. A Medium severity pattern translates to a 500,000
	// rank. And a High severity pattern translates to a 250,000 rank. Rank values
	// less than 1 or greater than 1,000,000 are reserved for AWS-provided patterns.
	Rank int32

	noSmithyDocumentSerde
}

type UpdateLogPatternOutput struct {

	// The successfully created log pattern.
	LogPattern *types.LogPattern

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLogPatternMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLogPattern{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLogPattern{}, middleware.After)
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
	if err = addOpUpdateLogPatternValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLogPattern(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLogPattern(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "UpdateLogPattern",
	}
}
