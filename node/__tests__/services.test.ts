import { expect, test } from '@jest/globals';
import ArticleService from '../stroeer/page/article/v1/article_page_service_grpc_pb';
import SectionService from '../stroeer/page/section/v1/section_page_service_grpc_pb';

test('gRPC article service client is generated and matches the snapshot', () => {
  expect(ArticleService).toMatchInlineSnapshot(`
    Object {
      "ArticlePageServiceClient": [Function],
      "ArticlePageServiceService": Object {
        "getArticlePage": Object {
          "path": "/stroeer.page.article.v1.ArticlePageService/GetArticlePage",
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
          "path": "/stroeer.page.section.v1.SectionPageService/GetSectionPage",
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
