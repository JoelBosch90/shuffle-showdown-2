import { OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { getApiGatewayOrigin } from './getApiGatewayOrigin';

describe('getApiGatewayOrigin', () => {
  const mockApiGatewayUrlName = 'MockApiGatewayUrl';

  beforeEach(() => {
    jest.clearAllMocks();

    process.env.API_GATEWAY_URL_NAME = mockApiGatewayUrlName;
  });

  it('should throw an error if API_GATEWAY_URL_NAME is not defined', () => {
    delete process.env.API_GATEWAY_URL_NAME;

    expect(() => getApiGatewayOrigin()).toThrow('API_GATEWAY_URL_NAME is not defined');
  });

  it('should return a valid HttpOrigin', () => {
    const origin = getApiGatewayOrigin();

    const bindResult = origin.bind({} as any, {} as any);
    expect(bindResult.originProperty?.domainName).toMatch(/\$\{Token\[TOKEN\.\d+\]\}/);
    expect((bindResult.originProperty?.customOriginConfig as any)?.originProtocolPolicy).toBe(OriginProtocolPolicy.HTTPS_ONLY);
  });
});