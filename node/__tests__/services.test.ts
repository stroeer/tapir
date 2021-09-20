import { expect, test } from '@jest/globals';
import CoreArticleService from '../stroeer/core/v1/core_article_service_grpc_pb';
import ArticlePageService from '../stroeer/page/article/v1/article_page_service_grpc_pb';
import SectionPageService from '../stroeer/page/section/v1/section_page_service_grpc_pb';
import FragmentStageService from '../stroeer/fragment/v1/stage_service_grpc_pb';

test('gRPC core article service client is generated and matches the snapshot', () => {
  expect(CoreArticleService).toMatchInlineSnapshot(`
Object {
  "ArticleServiceClient": [Function],
  "ArticleServiceService": Object {
    "getArticle": Object {
      "path": "/stroeer.core.v1.ArticleService/GetArticle",
      "requestDeserialize": [Function],
      "requestSerialize": [Function],
      "requestStream": false,
      "requestType": [Function],
      "responseDeserialize": [Function],
      "responseSerialize": [Function],
      "responseStream": false,
      "responseType": [Function],
    },
    "listArticles": Object {
      "path": "/stroeer.core.v1.ArticleService/ListArticles",
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

test('gRPC article page service client is generated and matches the snapshot', () => {
  expect(ArticlePageService).toMatchInlineSnapshot(`
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

test('gRPC section page service client is generated and matches the snapshot', () => {
  expect(SectionPageService).toMatchInlineSnapshot(`
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

test('gRPC fragment stage service client is generated and matches the snapshot', () => {
  expect(FragmentStageService).toMatchInlineSnapshot(`
Object {
  "StageServiceClient": [Function],
  "StageServiceService": Object {
    "getStages": Object {
      "path": "/stroeer.fragment.v1.StageService/GetStages",
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
