import { performance } from 'node:perf_hooks';
import fetch from 'node-fetch';
import { GetNavigationResponse } from '../../stroeer/navigation/v1/navigation_service_pb';

const validateEnvVar = (env: string | undefined, name: string): string => {
  if (!env || env.length < 1) {
    throw new Error(`You need to provide an env var "${name}"`);
  }

  return env;
};

const getAPIEndpoint = () => {
  const apiEndpoint = process.env['API_ENDPOINT'];

  return validateEnvVar(apiEndpoint, 'API_ENDPOINT');
};

function fetchData(acceptHeader: 'json' | 'protobuf' = 'json') {
  return fetch(getAPIEndpoint(), {
    headers: {
      'user-agent': 'tapir_node_buf_demo',
      Accept: `application/${acceptHeader}`,
    },
  }).then((response) => {
    if (!response.ok) {
      throw new Error(
        `The response is not ok. Status: ${response.status}. Text: ${response.statusText}`
      );
    }
    return response;
  });
}

const fetchJSON = () => {
  const start = performance.now();
  fetchData()
    .then((response) => response.json())
    .then((data: any) => {
      console.log('success: json');
      GetNavigationResponse.fromJson(data);
      const end = performance.now();

      console.log('json time', end - start);
    })
    .catch((error) => {
      console.error('There was an error while fetching the data', error);
    });
};

const fetchBinary = () => {
  const start = performance.now();
  fetchData('protobuf')
    .then((response) => response.buffer())
    .then((data) => {
      console.log('success: binary');
      GetNavigationResponse.fromBinary(data);
      const end = performance.now();

      console.log('binary time', end - start);
    })
    .catch((error) => {
      console.error('There was an error while fetching the data', error);
    });
};

fetchJSON();
fetchBinary();
