import { ConnectionOptions, connect } from 'tls';
import { createHash } from 'crypto';

export async function getGitHubThumbprint(): Promise<string> {
  return new Promise<string>((resolve, reject) => {
    const options: ConnectionOptions = {
      host: 'token.actions.githubusercontent.com',
      port: 443,
      // We disable certificate checks just to retrieve the certificate.
      rejectUnauthorized: false,
    };

    const socket = connect(options, () => {
      try {
        // Retrieve a detailed certificate chain.
        const cert = socket.getPeerCertificate(true);
        if (!cert || Object.keys(cert).length === 0) {
          return reject(new Error('No certificate received.'));
        }

        const fingerprint = createHash('sha1').update(cert.raw).digest('hex');
        resolve(fingerprint);
      } catch (error) {
        reject(error);
      } finally {
        socket.end();
      }
    });

    socket.on('error', reject);
  });
}