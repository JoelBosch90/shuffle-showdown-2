import { OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { getApiGatewayOrigin } from './getApiGatewayOrigin';

describe('getApiGatewayOrigin', () => {
  const mockApiGatewayUrlName = 'MockApiGatewayUrl';
  const mockStage = 'blah';

  beforeEach(() => {
    jest.clearAllMocks();

    process.env.PRIVATE_API_GATEWAY_URL_NAME = mockApiGatewayUrlName;
    process.env.PUBLIC_STAGE = mockStage;
  });

  it('throws an error if PRIVATE_API_GATEWAY_URL_NAME is not defined', () => {
    delete process.env.PRIVATE_API_GATEWAY_URL_NAME;

    expect(() => getApiGatewayOrigin()).toThrow('PRIVATE_API_GATEWAY_URL_NAME is not defined');
  });

  it('throws an error if PUBLIC_STAGE is not defined', () => {
    delete process.env.PUBLIC_STAGE;

    expect(() => getApiGatewayOrigin()).toThrow('PUBLIC_STAGE is not defined');
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