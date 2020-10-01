import { expect, test } from '@jest/globals';
import ArticleService from '../stroeer/pages/article/v1/article_service_grpc_pb';
import SectionService from '../stroeer/pages/section/v1/section_service_grpc_pb';

test('gRPC article service client is generated and matches the snapshot', () => {
  expect(ArticleService).toMatchInlineSnapshot(`
    Object {
      "ArticlePageServiceClient": [Function],
      "ArticlePageServiceService": Object {
        "getArticlePage": Object {
          "path": "/stroeer.pages.article.v1.ArticlePageService/GetArticlePage",
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
  expect(SectionService).toMatchInlineSnapshot(`
    Object {
      "SectionPageServiceClient": [Function],
      "SectionPageServiceService": Object {
        "getSectionPage": Object {
          "path": "/stroeer.pages.section.v1.SectionPageService/GetSectionPage",
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
