import { OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { getApiGatewayOrigin } from './getApiGatewayOrigin';
import { Stack } from 'aws-cdk-lib';
import { StringParameter } from 'aws-cdk-lib/aws-ssm';

const valueFromLookupSpy = jest.spyOn(StringParameter, 'valueFromLookup');

describe('getApiGatewayOrigin', () => {
  const mockApiGatewayUrl = 'MockApiGatewayUrl.execute-api.us-east-1.amazonaws.com';
  const mockApiGatewayUrlName = 'MockApiGatewayUrl';
  const mockStage = 'blah';
  const testStack = new Stack();

  beforeEach(() => {
    jest.clearAllMocks();

    process.env.PRIVATE_API_GATEWAY_URL_NAME = mockApiGatewayUrlName;
    process.env.PUBLIC_STAGE = mockStage;
    valueFromLookupSpy.mockReturnValue(`https://${mockApiGatewayUrl}`);
  });

  it('throws an error if PRIVATE_API_GATEWAY_URL_NAME is not defined', () => {
    delete process.env.PRIVATE_API_GATEWAY_URL_NAME;

    expect(() => getApiGatewayOrigin(testStack)).toThrow('PRIVATE_API_GATEWAY_URL_NAME is not defined');
  });

  it('throws an error if PUBLIC_STAGE is not defined', () => {
    delete process.env.PUBLIC_STAGE;

    expect(() => getApiGatewayOrigin(testStack)).toThrow('PUBLIC_STAGE is not defined');
  });

  it('returns a valid HttpOrigin', () => {
    const origin = getApiGatewayOrigin(testStack);

    const bindResult = origin.bind({} as any, {} as any);
    expect(bindResult.originProperty?.domainName).toMatch(mockApiGatewayUrl);
    expect((bindResult.originProperty?.customOriginConfig as any)?.originProtocolPolicy).toBe(OriginProtocolPolicy.HTTPS_ONLY);
  });

  it('uses the stage to set the originPath', () => {
    const origin = getApiGatewayOrigin(testStack);

    const bindResult = origin.bind({} as any, {} as any);
    expect(bindResult.originProperty?.originPath).toEqual(`/${mockStage}`);
  });
});