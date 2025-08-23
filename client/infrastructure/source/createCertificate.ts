import { Stack } from 'aws-cdk-lib';
import { Certificate, ICertificate } from 'aws-cdk-lib/aws-certificatemanager';

/**
 *  This function creates a certificate for the website bucket.
 *  It retrieves the certificate ARN from environment variables and returns the certificate.
 *  @param    {Stack} stack - The CDK stack in which to create the certificate.
 *  @throws   {Error} - If the PRIVATE_SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN environment variable is not defined.
 *  @returns  {ICertificate} - The created certificate.
 */
export const createCertificate = (stack: Stack): ICertificate => {
  const certificateArn = process.env.PRIVATE_SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN;
  if (!certificateArn) {
    throw new Error('Certificate ARN is not defined in environment variables.');
  }

  return Certificate.fromCertificateArn(stack, 'WebsiteCertificate', certificateArn);
};