// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var todo_pb = require('./todo_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_todo_TodoCreateRequest(arg) {
  if (!(arg instanceof todo_pb.TodoCreateRequest)) {
    throw new Error('Expected argument of type todo.TodoCreateRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_TodoCreateRequest(buffer_arg) {
  return todo_pb.TodoCreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_TodoFilter(arg) {
  if (!(arg instanceof todo_pb.TodoFilter)) {
    throw new Error('Expected argument of type todo.TodoFilter');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_TodoFilter(buffer_arg) {
  return todo_pb.TodoFilter.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_TodoResponse(arg) {
  if (!(arg instanceof todo_pb.TodoResponse)) {
    throw new Error('Expected argument of type todo.TodoResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_TodoResponse(buffer_arg) {
  return todo_pb.TodoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todo_TodoUpdateRequest(arg) {
  if (!(arg instanceof todo_pb.TodoUpdateRequest)) {
    throw new Error('Expected argument of type todo.TodoUpdateRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todo_TodoUpdateRequest(buffer_arg) {
  return todo_pb.TodoUpdateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var TodoService = exports.TodoService = {
  getTodos: {
    path: '/todo.Todo/GetTodos',
    requestStream: false,
    responseStream: true,
    requestType: todo_pb.TodoFilter,
    responseType: todo_pb.TodoResponse,
    requestSerialize: serialize_todo_TodoFilter,
    requestDeserialize: deserialize_todo_TodoFilter,
    responseSerialize: serialize_todo_TodoResponse,
    responseDeserialize: deserialize_todo_TodoResponse,
  },
  createTodo: {
    path: '/todo.Todo/CreateTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.TodoCreateRequest,
    responseType: todo_pb.TodoResponse,
    requestSerialize: serialize_todo_TodoCreateRequest,
    requestDeserialize: deserialize_todo_TodoCreateRequest,
    responseSerialize: serialize_todo_TodoResponse,
    responseDeserialize: deserialize_todo_TodoResponse,
  },
  updateTodo: {
    path: '/todo.Todo/UpdateTodo',
    requestStream: false,
    responseStream: false,
    requestType: todo_pb.TodoUpdateRequest,
    responseType: todo_pb.TodoResponse,
    requestSerialize: serialize_todo_TodoUpdateRequest,
    requestDeserialize: deserialize_todo_TodoUpdateRequest,
    responseSerialize: serialize_todo_TodoResponse,
    responseDeserialize: deserialize_todo_TodoResponse,
  },
};

exports.TodoClient = grpc.makeGenericClientConstructor(TodoService);
