// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports security findings generated by a finding provider into Security Hub.
// This action is requested by the finding provider to import its findings into
// Security Hub. BatchImportFindings must be called by one of the following:
//   - The Amazon Web Services account that is associated with a finding if you
//     are using the default product ARN (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-custom-providers.html#securityhub-custom-providers-bfi-reqs)
//     or are a partner sending findings from within a customer's Amazon Web Services
//     account. In these cases, the identifier of the account that you are calling
//     BatchImportFindings from needs to be the same as the AwsAccountId attribute
//     for the finding.
//   - An Amazon Web Services account that Security Hub has allow-listed for an
//     official partner integration. In this case, you can call BatchImportFindings
//     from the allow-listed account and send findings from different customer accounts
//     in the same batch.
//
// The maximum allowed size for a finding is 240 Kb. An error is returned for any
// finding larger than 240 Kb. After a finding is created, BatchImportFindings
// cannot be used to update the following finding fields and objects, which
// Security Hub customers use to manage their investigation workflow.
//   - Note
//   - UserDefinedFields
//   - VerificationState
//   - Workflow
//
// Finding providers also should not use BatchImportFindings to update the
// following attributes.
//   - Confidence
//   - Criticality
//   - RelatedFindings
//   - Severity
//   - Types
//
// Instead, finding providers use FindingProviderFields to provide values for
// these attributes.
func (c *Client) BatchImportFindings(ctx context.Context, params *BatchImportFindingsInput, optFns ...func(*Options)) (*BatchImportFindingsOutput, error) {
	if params == nil {
		params = &BatchImportFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchImportFindings", params, optFns, c.addOperationBatchImportFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchImportFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchImportFindingsInput struct {

	// A list of findings to import. To successfully import a finding, it must follow
	// the Amazon Web Services Security Finding Format (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html)
	// . Maximum of 100 findings per request.
	//
	// This member is required.
	Findings []types.AwsSecurityFinding

	noSmithyDocumentSerde
}

type BatchImportFindingsOutput struct {

	// The number of findings that failed to import.
	//
	// This member is required.
	FailedCount int32

	// The number of findings that were successfully imported.
	//
	// This member is required.
	SuccessCount int32

	// The list of findings that failed to import.
	FailedFindings []types.ImportFindingsError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchImportFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchImportFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchImportFindings{}, middleware.After)
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
	if err = addOpBatchImportFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchImportFindings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchImportFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchImportFindings",
	}
}
