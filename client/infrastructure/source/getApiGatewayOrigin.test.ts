import { OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { getApiGatewayOrigin } from './getApiGatewayOrigin';

describe('getApiGatewayOrigin', () => {
  const mockApiGatewayUrlName = 'MockApiGatewayUrl';
  const mockStage = 'blah';

  beforeEach(() => {
    jest.clearAllMocks();

    process.env.API_GATEWAY_URL_NAME = mockApiGatewayUrlName;
    process.env.STAGE = mockStage;
  });

  it('throws an error if API_GATEWAY_URL_NAME is not defined', () => {
    delete process.env.API_GATEWAY_URL_NAME;

    expect(() => getApiGatewayOrigin()).toThrow('API_GATEWAY_URL_NAME is not defined');
  });

  it('throws an error if STAGE is not defined', () => {
    delete process.env.STAGE;

    expect(() => getApiGatewayOrigin()).toThrow('STAGE is not defined');
  });

  it('returns a valid HttpOrigin', () => {
    const origin = getApiGatewayOrigin();

    const bindResult = origin.bind({} as any, {} as any);
    expect(bindResult.originProperty?.domainName).toMatch(/\$\{Token\[TOKEN\.\d+\]\}/);
    expect((bindResult.originProperty?.customOriginConfig as any)?.originProtocolPolicy).toBe(OriginProtocolPolicy.HTTPS_ONLY);
  });

  it('uses the stage to set the originPath', () => {
    const origin = getApiGatewayOrigin();

    const bindResult = origin.bind({} as any, {} as any);
    expect(bindResult.originProperty?.originPath).toEqual(`/${mockStage}`);
  });
});