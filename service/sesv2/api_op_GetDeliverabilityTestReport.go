// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve the results of a predictive inbox placement test.
func (c *Client) GetDeliverabilityTestReport(ctx context.Context, params *GetDeliverabilityTestReportInput, optFns ...func(*Options)) (*GetDeliverabilityTestReportOutput, error) {
	if params == nil {
		params = &GetDeliverabilityTestReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeliverabilityTestReport", params, optFns, c.addOperationGetDeliverabilityTestReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeliverabilityTestReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve the results of a predictive inbox placement test.
type GetDeliverabilityTestReportInput struct {

	// A unique string that identifies the predictive inbox placement test.
	//
	// This member is required.
	ReportId *string

	noSmithyDocumentSerde
}

// The results of the predictive inbox placement test.
type GetDeliverabilityTestReportOutput struct {

	// An object that contains the results of the predictive inbox placement test.
	//
	// This member is required.
	DeliverabilityTestReport *types.DeliverabilityTestReport

	// An object that describes how the test email was handled by several email
	// providers, including Gmail, Hotmail, Yahoo, AOL, and others.
	//
	// This member is required.
	IspPlacements []types.IspPlacement

	// An object that specifies how many test messages that were sent during the
	// predictive inbox placement test were delivered to recipients' inboxes, how many
	// were sent to recipients' spam folders, and how many weren't delivered.
	//
	// This member is required.
	OverallPlacement *types.PlacementStatistics

	// An object that contains the message that you sent when you performed this
	// predictive inbox placement test.
	Message *string

	// An array of objects that define the tags (keys and values) that are associated
	// with the predictive inbox placement test.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeliverabilityTestReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeliverabilityTestReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeliverabilityTestReport{}, middleware.After)
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
	if err = addOpGetDeliverabilityTestReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeliverabilityTestReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeliverabilityTestReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDeliverabilityTestReport",
	}
}
