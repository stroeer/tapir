// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var service_pb = require('./service_pb.js');
var content_pb = require('./content_pb.js');

function serialize_ArticleRequest(arg) {
  if (!(arg instanceof service_pb.ArticleRequest)) {
    throw new Error('Expected argument of type ArticleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ArticleRequest(buffer_arg) {
  return service_pb.ArticleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ArticleResponse(arg) {
  if (!(arg instanceof service_pb.ArticleResponse)) {
    throw new Error('Expected argument of type ArticleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ArticleResponse(buffer_arg) {
  return service_pb.ArticleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ArticleServiceService = exports.ArticleServiceService = {
  byId: {
    path: '/ArticleService/ById',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.ArticleRequest,
    responseType: service_pb.ArticleResponse,
    requestSerialize: serialize_ArticleRequest,
    requestDeserialize: deserialize_ArticleRequest,
    responseSerialize: serialize_ArticleResponse,
    responseDeserialize: deserialize_ArticleResponse,
  },
};

exports.ArticleServiceClient = grpc.makeGenericClientConstructor(ArticleServiceService);
