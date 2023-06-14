// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var proto_some_pb = require('../proto/some_pb.js');
var proto_nested_nested_pb = require('../proto/nested/nested_pb.js');
var gogoproto_gogo_pb = require('../gogoproto/gogo_pb.js');

function serialize_foo_Some(arg) {
  if (!(arg instanceof proto_some_pb.Some)) {
    throw new Error('Expected argument of type foo.Some');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_foo_Some(buffer_arg) {
  return proto_some_pb.Some.deserializeBinary(new Uint8Array(buffer_arg));
}


var SomeServiceService = exports.SomeServiceService = {
  getSome: {
    path: '/foo.SomeService/GetSome',
    requestStream: false,
    responseStream: false,
    requestType: proto_some_pb.Some,
    responseType: proto_some_pb.Some,
    requestSerialize: serialize_foo_Some,
    requestDeserialize: deserialize_foo_Some,
    responseSerialize: serialize_foo_Some,
    responseDeserialize: deserialize_foo_Some,
  },
};

exports.SomeServiceClient = grpc.makeGenericClientConstructor(SomeServiceService);
