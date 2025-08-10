const mockFnImportValue = jest.fn();

import { OriginProtocolPolicy } from 'aws-cdk-lib/aws-cloudfront';
import { getApiGatewayOrigin } from './getApiGatewayOrigin';

jest.mock('aws-cdk-lib', () => ({ Fn: { importValue: mockFnImportValue } }));

describe('getApiGatewayOrigin', () => {
  const mockDomainName = 'api.example.com';
  const mockPathName = '/path';
  const mockApiGatewayUrl = `https://${mockDomainName}${mockPathName}`;

  beforeEach(() => {
    jest.clearAllMocks();

    mockFnImportValue.mockReturnValue(mockApiGatewayUrl);
    process.env.API_GATEWAY_URL_NAME = 'MockApiGatewayUrl';
  });

  it('should throw an error if API_GATEWAY_URL_NAME is not defined', () => {
    delete process.env.API_GATEWAY_URL_NAME;

    expect(() => getApiGatewayOrigin()).toThrow('API_GATEWAY_URL_NAME is not defined');
  });

  it('should throw an error if the API Gateway URL is empty', () => {
    mockFnImportValue.mockReturnValue('');

    expect(() => getApiGatewayOrigin()).toThrow('API Gateway URL cannot be found');
  });

  it('should return a valid HttpOrigin', () => {
    const origin = getApiGatewayOrigin();

    const bindResult = origin.bind({} as any, {} as any);
    expect(bindResult.originProperty?.domainName).toBe(mockDomainName);
    expect(bindResult.originProperty?.originPath).toBe(mockPathName);
    expect((bindResult.originProperty?.customOriginConfig as any)?.originProtocolPolicy).toBe(OriginProtocolPolicy.HTTPS_ONLY);
  });
});