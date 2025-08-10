import { OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { HttpOrigin } from 'aws-cdk-lib/aws-cloudfront-origins';

/**
 *  This function retrieves the website domain from environment variables and returns an HttpOrigin.
 *  @throws   {Error} - If the SHUFFLE_SHOWDOWN_DOMAIN environment variable is not defined.
 *  @returns  {HttpOrigin} - The created website origin.
 */
export const getWebsiteDomain = (): HttpOrigin => {
  const domain = process.env.SHUFFLE_SHOWDOWN_DOMAIN;
  if (!domain) {
    throw new Error('Website domain is not defined');
  }

  return new HttpOrigin(domain, { protocolPolicy: OriginProtocolPolicy.HTTPS_ONLY });
};