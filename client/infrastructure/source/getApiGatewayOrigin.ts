import { Fn, Stack } from "aws-cdk-lib";
import { OriginProtocolPolicy } from "aws-cdk-lib/aws-cloudfront";
import { HttpOrigin } from "aws-cdk-lib/aws-cloudfront-origins";
import { StringParameter } from "aws-cdk-lib/aws-ssm";

/**
 *  This function retrieves the API Gateway URL from the CloudFormation stack output.
 *  It constructs an HttpOrigin using the URL's hostname and pathname.
 *  @throws   {Error} - If the API Gateway URL is not defined.
 *  @returns  {HttpOrigin} - The constructed HttpOrigin for the API Gateway.
 */
export function getApiGatewayOrigin(stack: Stack): HttpOrigin {
  const apiGateWayUrlName = process.env.PRIVATE_API_GATEWAY_URL_NAME;
  if (!apiGateWayUrlName) {
    throw new Error('PRIVATE_API_GATEWAY_URL_NAME is not defined');
  }

  const stage = process.env.PUBLIC_STAGE;
  if (!stage) {
    throw new Error('PUBLIC_STAGE is not defined');
  }

  const apiGatewayUrl = StringParameter.valueFromLookup(stack, apiGateWayUrlName);
  const withoutProtocol = Fn.select(1, Fn.split('://', apiGatewayUrl));
  const withoutPath = Fn.select(0, Fn.split('/', withoutProtocol));

  return new HttpOrigin(withoutPath, {
    protocolPolicy: OriginProtocolPolicy.HTTPS_ONLY,
    originPath: `/${stage}`,
  });
}