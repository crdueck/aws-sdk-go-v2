// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of Recommendation (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_Recommendation.html)
// objects that contain recommendations for a profiling group for a given time
// period. A list of Anomaly (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_Anomaly.html)
// objects that contains details about anomalies detected in the profiling group
// for the same time period is also returned.
func (c *Client) GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) {
	if params == nil {
		params = &GetRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecommendations", params, optFns, c.addOperationGetRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the GetRecommendationsRequest.
type GetRecommendationsInput struct {

	// The start time of the profile to get analysis data about. You must specify
	// startTime and endTime . This is specified using the ISO 8601 format. For
	// example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020
	// 1:15:02 PM UTC.
	//
	// This member is required.
	EndTime *time.Time

	// The name of the profiling group to get analysis data about.
	//
	// This member is required.
	ProfilingGroupName *string

	// The end time of the profile to get analysis data about. You must specify
	// startTime and endTime . This is specified using the ISO 8601 format. For
	// example, 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020
	// 1:15:02 PM UTC.
	//
	// This member is required.
	StartTime *time.Time

	// The language used to provide analysis. Specify using a string that is one of
	// the following BCP 47 language codes.
	//   - de-DE - German, Germany
	//   - en-GB - English, United Kingdom
	//   - en-US - English, United States
	//   - es-ES - Spanish, Spain
	//   - fr-FR - French, France
	//   - it-IT - Italian, Italy
	//   - ja-JP - Japanese, Japan
	//   - ko-KR - Korean, Republic of Korea
	//   - pt-BR - Portugese, Brazil
	//   - zh-CN - Chinese, China
	//   - zh-TW - Chinese, Taiwan
	Locale *string

	noSmithyDocumentSerde
}

// The structure representing the GetRecommendationsResponse.
type GetRecommendationsOutput struct {

	// The list of anomalies that the analysis has found for this profile.
	//
	// This member is required.
	Anomalies []types.Anomaly

	// The end time of the profile the analysis data is about. This is specified using
	// the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z represents 1
	// millisecond past June 1, 2020 1:15:02 PM UTC.
	//
	// This member is required.
	ProfileEndTime *time.Time

	// The start time of the profile the analysis data is about. This is specified
	// using the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z represents 1
	// millisecond past June 1, 2020 1:15:02 PM UTC.
	//
	// This member is required.
	ProfileStartTime *time.Time

	// The name of the profiling group the analysis data is about.
	//
	// This member is required.
	ProfilingGroupName *string

	// The list of recommendations that the analysis found for this profile.
	//
	// This member is required.
	Recommendations []types.Recommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecommendations{}, middleware.After)
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
	if err = addOpGetRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "GetRecommendations",
	}
}
