// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Inserts or deletes ActivatedRule objects in a WebACL . Each Rule
// identifies web requests that you want to allow, block, or count. When you update
// a WebACL , you specify the following values:
//   - A default action for the WebACL , either ALLOW or BLOCK . AWS WAF performs
//     the default action if a request doesn't match the criteria in any of the Rules
//     in a WebACL .
//   - The Rules that you want to add or delete. If you want to replace one Rule
//     with another, you delete the existing Rule and add the new one.
//   - For each Rule , whether you want AWS WAF to allow requests, block requests,
//     or count requests that match the conditions in the Rule .
//   - The order in which you want AWS WAF to evaluate the Rules in a WebACL . If
//     you add more than one Rule to a WebACL , AWS WAF evaluates each request
//     against the Rules in order based on the value of Priority . (The Rule that has
//     the lowest value for Priority is evaluated first.) When a web request matches
//     all the predicates (such as ByteMatchSets and IPSets ) in a Rule , AWS WAF
//     immediately takes the corresponding action, allow or block, and doesn't evaluate
//     the request against the remaining Rules in the WebACL , if any.
//
// To create and configure a WebACL , perform the following steps:
//   - Create and update the predicates that you want to include in Rules . For
//     more information, see CreateByteMatchSet , UpdateByteMatchSet , CreateIPSet ,
//     UpdateIPSet , CreateSqlInjectionMatchSet , and UpdateSqlInjectionMatchSet .
//   - Create and update the Rules that you want to include in the WebACL . For
//     more information, see CreateRule and UpdateRule .
//   - Create a WebACL . See CreateWebACL .
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of an UpdateWebACL request.
//   - Submit an UpdateWebACL request to specify the Rules that you want to include
//     in the WebACL , to specify the default action, and to associate the WebACL
//     with a CloudFront distribution. The ActivatedRule can be a rule group. If you
//     specify a rule group as your ActivatedRule , you can exclude specific rules
//     from that rule group. If you already have a rule group associated with a web ACL
//     and want to submit an UpdateWebACL request to exclude certain rules from that
//     rule group, you must first remove the rule group from the web ACL, the re-insert
//     it again, specifying the excluded rules. For details, see
//     ActivatedRule$ExcludedRules .
//
// Be aware that if you try to add a RATE_BASED rule to a web ACL without setting
// the rule type when first creating the rule, the UpdateWebACL request will fail
// because the request tries to add a REGULAR rule (the default rule type) with the
// specified ID, which does not exist. For more information about how to use the
// AWS WAF API to allow or block HTTP requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/)
// .
func (c *Client) UpdateWebACL(ctx context.Context, params *UpdateWebACLInput, optFns ...func(*Options)) (*UpdateWebACLOutput, error) {
	if params == nil {
		params = &UpdateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWebACL", params, optFns, c.addOperationUpdateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWebACLInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// The WebACLId of the WebACL that you want to update. WebACLId is returned by
	// CreateWebACL and by ListWebACLs .
	//
	// This member is required.
	WebACLId *string

	// A default action for the web ACL, either ALLOW or BLOCK. AWS WAF performs the
	// default action if a request doesn't match the criteria in any of the rules in a
	// web ACL.
	DefaultAction *types.WafAction

	// An array of updates to make to the WebACL . An array of WebACLUpdate objects
	// that you want to insert into or delete from a WebACL . For more information, see
	// the applicable data types:
	//   - WebACLUpdate : Contains Action and ActivatedRule
	//   - ActivatedRule : Contains Action , OverrideAction , Priority , RuleId , and
	//   Type . ActivatedRule|OverrideAction applies only when updating or adding a
	//   RuleGroup to a WebACL . In this case, you do not use ActivatedRule|Action .
	//   For all other update requests, ActivatedRule|Action is used instead of
	//   ActivatedRule|OverrideAction .
	//   - WafAction : Contains Type
	Updates []types.WebACLUpdate

	noSmithyDocumentSerde
}

type UpdateWebACLOutput struct {

	// The ChangeToken that you used to submit the UpdateWebACL request. You can also
	// use this value to query the status of the request. For more information, see
	// GetChangeTokenStatus .
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWebACL{}, middleware.After)
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
	if err = addOpUpdateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "UpdateWebACL",
	}
}
