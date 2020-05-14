const GrpcArticleServiceClient = require('../v1/article/service_grpc_pb');
const GrpcSectionServiceClient = require('../v1/section/service_grpc_pb');

test('gRPC article service client is generated and matches the snapshot', () => {
  expect(GrpcArticleServiceClient).toMatchInlineSnapshot(`
    Object {
      "ArticleServiceClient": [Function],
      "ArticleServiceService": Object {
        "getArticle": Object {
          "path": "/v1.article.ArticleService/GetArticle",
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
      "SectionServiceClient": [Function],
      "SectionServiceService": Object {
        "getSection": Object {
          "path": "/v1.section.SectionService/GetSection",
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
