import { Stack } from 'aws-cdk-lib';
import { createCertificate } from './createCertificate';

describe('createCertificate', () => {
  it('throws an error if certificate ARN is not defined', () => {
    delete process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN;

    const stack = new Stack();

    expect(() => createCertificate(stack)).toThrow('Certificate ARN is not defined in environment variables.');
  });

  it('creates a certificate from the ARN', () => {
    const fakeCertificateArn = 'arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012';
    process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN = fakeCertificateArn;

    const stack = new Stack();
    const certificate = createCertificate(stack);

    expect(certificate.certificateArn).toBe(fakeCertificateArn);
  });

  it('creates a certificate with the correct name', () => {
    const fakeCertificateArn = 'arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012';
    process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN = fakeCertificateArn;

    const stack = new Stack();
    const certificate = createCertificate(stack);

    expect(certificate.node.id).toBe('WebsiteCertificate');
  });
});