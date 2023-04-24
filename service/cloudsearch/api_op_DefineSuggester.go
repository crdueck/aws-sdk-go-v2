// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures a suggester for a domain. A suggester enables you to display
// possible matches before users finish typing their queries. When you configure a
// suggester, you must specify the name of the text field you want to search for
// possible matches and a unique name for the suggester. For more information, see
// Getting Search Suggestions (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DefineSuggester(ctx context.Context, params *DefineSuggesterInput, optFns ...func(*Options)) (*DefineSuggesterOutput, error) {
	if params == nil {
		params = &DefineSuggesterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DefineSuggester", params, optFns, c.addOperationDefineSuggesterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DefineSuggesterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DefineSuggester operation. Specifies the
// name of the domain you want to update and the suggester configuration.
type DefineSuggesterInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string

	// Configuration information for a search suggester. Each suggester has a unique
	// name and specifies the text field you want to use for suggestions. The following
	// options can be configured for a suggester: FuzzyMatching , SortExpression .
	//
	// This member is required.
	Suggester *types.Suggester

	noSmithyDocumentSerde
}

// The result of a DefineSuggester request. Contains the status of the
// newly-configured suggester.
type DefineSuggesterOutput struct {

	// The value of a Suggester and its current status.
	//
	// This member is required.
	Suggester *types.SuggesterStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDefineSuggesterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDefineSuggester{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDefineSuggester{}, middleware.After)
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
	if err = addOpDefineSuggesterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDefineSuggester(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDefineSuggester(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DefineSuggester",
	}
}
