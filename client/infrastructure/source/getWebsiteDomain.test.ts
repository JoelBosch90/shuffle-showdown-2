import { Construct } from "constructs";
import { getWebsiteDomain } from "./getWebsiteDomain";
import { App } from "aws-cdk-lib";
import { OriginProtocolPolicy } from "aws-cdk-lib/aws-cloudfront";
import { CfnDistribution } from "aws-cdk-lib/aws-cloudfront";

describe('getWebsiteDomain', () => {
  const mockApp = new App();
  const mockConstructName = 'MockConstruct';
  const mockConstruct = new Construct(mockApp, mockConstructName);
  const mockOriginId = 'MockOriginId';

  it('throws an error if the website domain is not defined', () => {
    delete process.env.SHUFFLE_SHOWDOWN_DOMAIN;

    expect(() => getWebsiteDomain()).toThrow('Website domain is not defined');
  });

  it('creates an HttpOrigin with the correct URL', () => {
    process.env.SHUFFLE_SHOWDOWN_DOMAIN = 'example.com';

    const origin = getWebsiteDomain();

    const bindConfig = origin.bind(mockConstruct, { originId: mockOriginId });
    expect(bindConfig.originProperty).toBeDefined();
    expect(bindConfig.originProperty?.id).toBe(mockOriginId);
    expect(bindConfig.originProperty?.domainName).toBe('example.com');
    expect(bindConfig.originProperty?.originPath).toBeUndefined();
  });

  it('creates an HttpOrigin with the HTTPS configuration', () => {
    process.env.SHUFFLE_SHOWDOWN_DOMAIN = 'example.com';

    const origin = getWebsiteDomain();

    const bindConfig = origin.bind(mockConstruct, { originId: mockOriginId });
    expect((bindConfig.originProperty?.customOriginConfig as CfnDistribution.CustomOriginConfigProperty).originProtocolPolicy).toBe(OriginProtocolPolicy.HTTPS_ONLY);
  });
});