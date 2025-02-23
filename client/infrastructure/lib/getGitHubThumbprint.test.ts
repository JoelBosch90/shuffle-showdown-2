import { createHash } from 'crypto';
import { connect } from 'tls';
import { getGitHubThumbprint } from './getGitHubThumbprint';

jest.mock('tls', () => ({
  ...jest.requireActual('tls'),
  connect: jest.fn(),
}));

describe('getGitHubThumbprint', () => {
  const socket = {
    getPeerCertificate: jest.fn(),
    end: jest.fn(),
    on: jest.fn(),
  };

  beforeEach(() => {
    (connect as jest.Mock).mockImplementation((_options, callback) => {
      setImmediate(callback); // Call callback asynchronously to mimic real behavior.

      return socket;
    });
  })

  afterEach(() => {
    jest.resetAllMocks();
  });

  it('returns fingerprint when certificate.raw exists', async () => {
    const rawCertificate = Buffer.from('abc');
    const expected = createHash('sha1').update(rawCertificate).digest('hex');
    socket.getPeerCertificate.mockReturnValue({ raw: rawCertificate });

    const fingerprint = await getGitHubThumbprint();

    expect(fingerprint).toBe(expected);
    expect(socket.getPeerCertificate).toHaveBeenCalledWith(true);
    expect(socket.end).toHaveBeenCalled();
  });

  it('returns undefined when certificate.raw is missing', async () => {
    socket.getPeerCertificate.mockReturnValue({});

    const result = await getGitHubThumbprint();

    expect(result).toBeUndefined();
  });

  it('returns undefined when the socket connection fails', async () => {
    (connect as jest.Mock).mockImplementation((_options, _callback) => ({
      on: (_event: string, callback: (_error?: Error) => void) =>
        setImmediate(() => callback(new Error('Connection failed'))),
      end: jest.fn(),
    }));

    const result = await getGitHubThumbprint();

    expect(result).toBeUndefined();
  });
});