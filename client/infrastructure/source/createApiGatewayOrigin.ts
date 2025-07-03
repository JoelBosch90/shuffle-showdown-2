import { OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { HttpOrigin } from 'aws-cdk-lib/aws-cloudfront-origins';

export const createApiGatewayOrigin = (): HttpOrigin => {
  const apiGatewayUrl = process.env.API_GATEWAY_URL;
  if (!apiGatewayUrl) {
    throw new Error('API Gateway URL is not defined');
  }

  return new HttpOrigin(apiGatewayUrl.replace('https://', '').replace('/api', ''), {
    protocolPolicy: OriginProtocolPolicy.HTTPS_ONLY,
  });
};