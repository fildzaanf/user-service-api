import http from 'k6/http';
import { check, sleep } from 'k6';

const JWT_TOKEN = 'JWT_TOKEN';

export const options = {
  vus: 100,
  duration: '30s', 
};

export default function () {
  const res = http.get(
    'http://localhost:8081/users/USER_ID',
    {
      headers: {
        Authorization: `Bearer ${JWT_TOKEN}`,
      },
    }
  );

  check(res, {
    'REST Status 200': (r) => r.status === 200,
  });

  sleep(1);
}
