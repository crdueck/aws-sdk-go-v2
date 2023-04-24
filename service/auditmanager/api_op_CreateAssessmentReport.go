// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an assessment report for the specified assessment.
func (c *Client) CreateAssessmentReport(ctx context.Context, params *CreateAssessmentReportInput, optFns ...func(*Options)) (*CreateAssessmentReportOutput, error) {
	if params == nil {
		params = &CreateAssessmentReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAssessmentReport", params, optFns, c.addOperationCreateAssessmentReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAssessmentReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAssessmentReportInput struct {

	// The identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	// The name of the new assessment report.
	//
	// This member is required.
	Name *string

	// The description of the assessment report.
	Description *string

	// A SQL statement that represents an evidence finder query. Provide this
	// parameter when you want to generate an assessment report from the results of an
	// evidence finder search query. When you use this parameter, Audit Manager
	// generates a one-time report using only the evidence from the query output. This
	// report does not include any assessment evidence that was manually added to a
	// report using the console (https://docs.aws.amazon.com/audit-manager/latest/userguide/generate-assessment-report.html#generate-assessment-report-include-evidence)
	// , or associated with a report using the API (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_BatchAssociateAssessmentReportEvidence.html)
	// . To use this parameter, the enablementStatus (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_EvidenceFinderEnablement.html#auditmanager-Type-EvidenceFinderEnablement-enablementStatus)
	// of evidence finder must be ENABLED . For examples and help resolving
	// queryStatement validation exceptions, see Troubleshooting evidence finder issues (https://docs.aws.amazon.com/audit-manager/latest/userguide/evidence-finder-issues.html#querystatement-exceptions)
	// in the Audit Manager User Guide.
	QueryStatement *string

	noSmithyDocumentSerde
}

type CreateAssessmentReportOutput struct {

	// The new assessment report that the CreateAssessmentReport API returned.
	AssessmentReport *types.AssessmentReport

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAssessmentReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAssessmentReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAssessmentReport{}, middleware.After)
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
	if err = addOpCreateAssessmentReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAssessmentReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAssessmentReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "CreateAssessmentReport",
	}
}
