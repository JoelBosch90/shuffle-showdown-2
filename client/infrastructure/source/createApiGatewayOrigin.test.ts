import { Construct } from "constructs";
import { createApiGatewayOrigin } from "./createApiGatewayOrigin";
import { App } from "aws-cdk-lib";
import { OriginProtocolPolicy } from "aws-cdk-lib/aws-cloudfront";
import { CfnDistribution } from "aws-cdk-lib/aws-cloudfront";

describe('createApiGatewayOrigin', () => {
  const mockApp = new App();
  const mockConstructName = 'MockConstruct';
  const mockConstruct = new Construct(mockApp, mockConstructName);
  const mockOriginId = 'MockOriginId';

  it('throws an error if API Gateway URL is not defined', () => {
    delete process.env.API_GATEWAY_URL;

    expect(() => createApiGatewayOrigin()).toThrow('API Gateway URL is not defined');
  });

  it('creates an HttpOrigin with the correct URL', () => {
    process.env.API_GATEWAY_URL = 'https://api.example.com/api';

    const origin = createApiGatewayOrigin();

    const bindConfig = origin.bind(mockConstruct, { originId: mockOriginId });
    expect(bindConfig.originProperty).toBeDefined();
    expect(bindConfig.originProperty?.id).toBe(mockOriginId);
    expect(bindConfig.originProperty?.domainName).toBe('api.example.com');
    expect(bindConfig.originProperty?.originPath).toBeUndefined();
  });

  it('creates an HttpOrigin with the HTTPS configuration', () => {
    process.env.API_GATEWAY_URL = 'https://api.example.com/api';

    const origin = createApiGatewayOrigin();

    const bindConfig = origin.bind(mockConstruct, { originId: mockOriginId });
    expect((bindConfig.originProperty?.customOriginConfig as CfnDistribution.CustomOriginConfigProperty).originProtocolPolicy).toBe(OriginProtocolPolicy.HTTPS_ONLY);
  });
});