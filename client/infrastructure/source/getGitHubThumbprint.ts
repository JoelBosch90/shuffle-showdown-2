import { ConnectionOptions, connect, type TLSSocket } from 'tls';
import { createHash } from 'crypto';

/**
 *  Gets the GitHub certificate thumbprint.
 *  Based on https://github.com/aws-actions/configure-aws-credentials/issues/357#issuecomment-1011642085
 *  @returns  {Promise<string | undefined>} - The SHA-1 thumbprint of the GitHub Actions certificate, or undefined if it cannot be retrieved.
 */
export async function getGitHubThumbprint(): Promise<string | undefined> {
  const options: ConnectionOptions = {
    host: 'token.actions.githubusercontent.com',
    port: 443,
    // We disable certificate checks just to retrieve the certificate.
    rejectUnauthorized: false,
    timeout: 3000,
  };
  const socket = await new Promise<TLSSocket | undefined>((resolve) => {
    const connection = connect(options, () => resolve(connection));
    connection.on('error', () => { resolve(undefined); });
  });

  if (!socket) {
    return undefined;
  }

  const certificate = socket.getPeerCertificate(true);
  socket.end();

  if (certificate && certificate.raw) {
    return createHash('sha1').update(certificate.raw).digest('hex');
  }

  return undefined;
};
