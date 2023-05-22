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

const getAuthorizationHeader = () => {
  const header = process.env['AUTH_HEADER'];

  return validateEnvVar(header, 'AUTH_HEADER');
};

function fetchData() {
  return fetch(getAPIEndpoint(), {
    headers: {
      'user-agent': 'tapir_node_buf_demo',
      Authorization: getAuthorizationHeader(),
    },
  }).then((response) => {
    if (!response.ok) {
      throw new Error(
        `The response is not ok. Status: ${response.status}. Text: ${response.statusText}`
      );
    }
    return response.json();
  });
}

fetchData()
  .then((data) => {
    console.log('success');
    const navigationResponse = GetNavigationResponse.fromJson(data);
    console.log({
      navi: JSON.stringify(navigationResponse.navigationMenu, null, 2),
    });
  })
  .catch((error) => {
    console.error('There was an error while fetching the data', error);
  });
