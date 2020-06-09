import { expect, test } from '@jest/globals';
import WebArticleService from '../stroeer/web/article/v1/web_article_service_grpc_pb';
import WebSectionService from '../stroeer/web/section/v1/web_section_service_grpc_pb';

test('gRPC article service client is generated and matches the snapshot', () => {
  expect(WebArticleService).toMatchInlineSnapshot(`
    Object {
      "WebArticleServiceClient": [Function],
      "WebArticleServiceService": Object {
        "getWebArticlePage": Object {
          "path": "/stroeer.web.article.v1.WebArticleService/GetWebArticlePage",
          "requestDeserialize": [Function],
          "requestSerialize": [Function],
          "requestStream": false,
          "requestType": [Function],
          "responseDeserialize": [Function],
          "responseSerialize": [Function],
          "responseStream": false,
          "responseType": [Function],
        },
      },
    }
  `);
});

test('gRPC section service client is generated and matches the snapshot', () => {
  expect(WebSectionService).toMatchInlineSnapshot(`
    Object {
      "WebSectionServiceClient": [Function],
      "WebSectionServiceService": Object {
        "getWebSectionPage": Object {
          "path": "/stroeer.web.section.v1.WebSectionService/GetWebSectionPage",
          "requestDeserialize": [Function],
          "requestSerialize": [Function],
          "requestStream": false,
          "requestType": [Function],
          "responseDeserialize": [Function],
          "responseSerialize": [Function],
          "responseStream": false,
          "responseType": [Function],
        },
      },
    }
  `);
});
