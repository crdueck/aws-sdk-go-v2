// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns Amazon EC2 instance recommendations. Compute Optimizer generates
// recommendations for Amazon Elastic Compute Cloud (Amazon EC2) instances that
// meet a specific set of requirements. For more information, see the Supported
// resources and requirements (https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html)
// in the Compute Optimizer User Guide.
func (c *Client) GetEC2InstanceRecommendations(ctx context.Context, params *GetEC2InstanceRecommendationsInput, optFns ...func(*Options)) (*GetEC2InstanceRecommendationsOutput, error) {
	if params == nil {
		params = &GetEC2InstanceRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEC2InstanceRecommendations", params, optFns, c.addOperationGetEC2InstanceRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEC2InstanceRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEC2InstanceRecommendationsInput struct {

	// The ID of the Amazon Web Services account for which to return instance
	// recommendations. If your account is the management account of an organization,
	// use this parameter to specify the member account for which you want to return
	// instance recommendations. Only one account ID can be specified per request.
	AccountIds []string

	// An array of objects to specify a filter that returns a more specific list of
	// instance recommendations.
	Filters []types.Filter

	// The Amazon Resource Name (ARN) of the instances for which to return
	// recommendations.
	InstanceArns []string

	// The maximum number of instance recommendations to return with a single request.
	// To retrieve the remaining results, make another request with the returned
	// nextToken value.
	MaxResults *int32

	// The token to advance to the next page of instance recommendations.
	NextToken *string

	// An object to specify the preferences for the Amazon EC2 instance
	// recommendations to return in the response.
	RecommendationPreferences *types.RecommendationPreferences

	noSmithyDocumentSerde
}

type GetEC2InstanceRecommendationsOutput struct {

	// An array of objects that describe errors of the request. For example, an error
	// is returned if you request recommendations for an instance of an unsupported
	// instance family.
	Errors []types.GetRecommendationError

	// An array of objects that describe instance recommendations.
	InstanceRecommendations []types.InstanceRecommendation

	// The token to use to advance to the next page of instance recommendations. This
	// value is null when there are no more pages of instance recommendations to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEC2InstanceRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEC2InstanceRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEC2InstanceRecommendations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEC2InstanceRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEC2InstanceRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetEC2InstanceRecommendations",
	}
}
