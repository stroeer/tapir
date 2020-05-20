const GrpcArticleServiceClient = require('../stroeer/article/v1/service_grpc_pb');
const GrpcSectionServiceClient = require('../stroeer/section/v1/service_grpc_pb');

test('gRPC article service client is generated and matches the snapshot', () => {
  expect(GrpcArticleServiceClient).toMatchInlineSnapshot(`
    Object {
      "ArticleServiceClient": [Function],
      "ArticleServiceService": Object {
        "getArticle": Object {
          "path": "/stroeer.article.v1.ArticleService/GetArticle",
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
          "path": "/stroeer.section.v1.SectionService/GetSection",
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
