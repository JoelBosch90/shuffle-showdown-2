import { Fn } from "aws-cdk-lib";
import { OriginProtocolPolicy } from "aws-cdk-lib/aws-cloudfront";
import { HttpOrigin } from "aws-cdk-lib/aws-cloudfront-origins";

/**
 *  This function retrieves the API Gateway URL from the CloudFormation stack output.
 *  It constructs an HttpOrigin using the URL's hostname and pathname.
 *  @throws   {Error} - If the API Gateway URL is not defined.
 *  @returns  {HttpOrigin} - The constructed HttpOrigin for the API Gateway.
 */
export function getApiGatewayOrigin(): HttpOrigin {
  const apiGateWayUrlName = process.env.API_GATEWAY_URL_NAME;
  if (!apiGateWayUrlName) {
    throw new Error('API_GATEWAY_URL_NAME is not defined');
  }

  // Turns the URL into a token to be resolved at deployment time.
  const apiDomain = Fn.select(1, Fn.split('://', Fn.select(0, Fn.split('/', Fn.importValue(apiGateWayUrlName)))));

  return new HttpOrigin(apiDomain, {
    protocolPolicy: OriginProtocolPolicy.HTTPS_ONLY,
  });
}