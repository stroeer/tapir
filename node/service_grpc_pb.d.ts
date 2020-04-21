// GENERATED CODE -- DO NOT EDIT!

// package: 
// file: service.proto

import * as service_pb from "./service_pb";
import * as grpc from "grpc";

interface IArticleServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  byId: grpc.MethodDefinition<service_pb.ArticleRequest, service_pb.ArticleResponse>;
}

export const ArticleServiceService: IArticleServiceService;

export class ArticleServiceClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  byId(argument: service_pb.ArticleRequest, callback: grpc.requestCallback<service_pb.ArticleResponse>): grpc.ClientUnaryCall;
  byId(argument: service_pb.ArticleRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<service_pb.ArticleResponse>): grpc.ClientUnaryCall;
  byId(argument: service_pb.ArticleRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<service_pb.ArticleResponse>): grpc.ClientUnaryCall;
}
