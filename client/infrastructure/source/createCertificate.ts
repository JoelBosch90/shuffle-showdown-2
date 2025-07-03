import { Stack } from 'aws-cdk-lib';
import { Certificate, ICertificate } from 'aws-cdk-lib/aws-certificatemanager';

export const createCertificate = (stack: Stack): ICertificate => {
  const certificateArn = process.env.SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN;
  if (!certificateArn) {
    throw new Error('Certificate ARN is not defined in environment variables.');
  }

  return Certificate.fromCertificateArn(stack, 'WebsiteCertificate', certificateArn);
};