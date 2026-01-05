import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

const client = new grpc.Client();
client.load(['proto'], 'user.proto');

const JWT_TOKEN = 'JWT_TOKEN';

export const options = {
  vus: 100,
  duration: '30s',
};

export default function () {
  client.connect('localhost:8080', {
    plaintext: true,
  });

  const res = client.invoke(
    'user.UserQueryService/GetUserByID',
    { id: 'USER_ID' },
    {
      metadata: {
        Authorization: `Bearer ${JWT_TOKEN}`,
      },
    }
  );

  check(res, {
    'gRPC OK': (r) => r && r.status === grpc.StatusOK,
  });

  client.close();
  sleep(1);
}
