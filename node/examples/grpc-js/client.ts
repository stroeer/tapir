import { GetSectionPageRequest } from '../../stroeer/page/section/v1/section_page_service_pb';
import { SectionPageServiceClient } from '../../stroeer/page/section/v1/section_page_service_grpc_pb';
import { credentials, Metadata } from '@grpc/grpc-js';

import type { ServiceError, CallOptions } from '@grpc/grpc-js';
import type { SectionPage } from '../../stroeer/page/section/v1/section_page_pb';

const host: string = process.env['GRPC_HOST'] ?? '';
const apiToken: string = process.env['GRPC_API_TOKEN'] ?? '';
const ssl = credentials.createSsl();

const client = new SectionPageServiceClient(host, ssl, {
  'grpc.primary_user_agent': 'tapir/node test-client',
});

const createRequest = (sectionPath: string) => {
  const request = new GetSectionPageRequest();
  request.setPage(0);
  request.setSectionPath(sectionPath);
  return request;
};

const callOptions = (): CallOptions => {
  return {
    deadline: Date.now() + 6000,
  };
};

const metaData = new Metadata();
metaData.set('authorization', apiToken);

export const getSectionPage = (
  sectionPath: string,
  cb: (err: ServiceError | null, sectionPage?: SectionPage | undefined) => void,
) => {
  const request = createRequest(sectionPath);
  console.log({ sectionPath });
  console.time('request');
  client.getSectionPage(request, metaData, callOptions(), (err, value) => {
    if (err) {
      cb(err);
    } else {
      cb(null, value?.getSectionPage());
    }
    console.timeEnd('request');
  });
};
