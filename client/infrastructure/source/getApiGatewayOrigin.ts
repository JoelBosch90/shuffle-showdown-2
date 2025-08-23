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

  const stage = process.env.STAGE;
  if (!stage) {
    throw new Error('STAGE is not defined');
  }

  // Turns the URL into a token to be resolved at deployment time.
  const apiGateWayUrl = Fn.importValue(apiGateWayUrlName);
  const withoutProtocol = Fn.select(1, Fn.split('://', apiGateWayUrl));
  const withoutPath = Fn.select(0, Fn.split('/', withoutProtocol));

  return new HttpOrigin(withoutPath, {
    protocolPolicy: OriginProtocolPolicy.HTTPS_ONLY,
    originPath: `/${stage}`,
  });
}