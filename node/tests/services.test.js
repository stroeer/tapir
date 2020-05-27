const GrpcArticleServiceClient = require('../stroeer/web/article/v1/web_article_service_grpc_pb');
const GrpcSectionServiceClient = require('../stroeer/web/section/v1/web_section_service_grpc_pb');

test('gRPC article service client is generated and matches the snapshot', () => {
  expect(GrpcArticleServiceClient).toMatchInlineSnapshot(`
    Object {
      "WebArticleServiceClient": [Function],
      "WebArticleServiceService": Object {
        "getArticlePage": Object {
          "path": "/stroeer.web.article.v1.WebArticleService/GetArticlePage",
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
  expect(GrpcSectionServiceClient).toMatchInlineSnapshot(`
    Object {
      "WebSectionServiceClient": [Function],
      "WebSectionServiceService": Object {
        "getSectionPage": Object {
          "path": "/stroeer.web.section.v1.WebSectionService/GetSectionPage",
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
