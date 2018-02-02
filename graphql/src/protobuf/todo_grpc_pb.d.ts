// package: todo
// file: todo.proto

import * as grpc from "grpc";
import * as todo_pb from "./todo_pb";

interface ITodoService extends grpc.ServiceDefinition {
    getTodos: IGetTodos;
    createTodo: ICreateTodo;
    updateTodo: IUpdateTodo;
}

interface IGetTodos {
    path: string; // "/todo.Todo/GetTodos"
    requestStream: boolean; // false
    responseStream: boolean; // true
    requestType: todo_pb.TodoFilter,
    responseType: todo_pb.TodoResponse,
    requestSerialize: (arg: todo_pb.TodoFilter) => Buffer;
    requestDeserialize: (buffer: Uint8Array) => todo_pb.TodoFilter;
    responseSerialize: (arg: todo_pb.TodoResponse) => Buffer;
    responseDeserialize: (buffer: Uint8Array) => todo_pb.TodoResponse;
}
interface ICreateTodo {
    path: string; // "/todo.Todo/CreateTodo"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestType: todo_pb.TodoCreateRequest,
    responseType: todo_pb.TodoResponse,
    requestSerialize: (arg: todo_pb.TodoCreateRequest) => Buffer;
    requestDeserialize: (buffer: Uint8Array) => todo_pb.TodoCreateRequest;
    responseSerialize: (arg: todo_pb.TodoResponse) => Buffer;
    responseDeserialize: (buffer: Uint8Array) => todo_pb.TodoResponse;
}
interface IUpdateTodo {
    path: string; // "/todo.Todo/UpdateTodo"
    requestStream: boolean; // false
    responseStream: boolean; // false
    requestType: todo_pb.TodoUpdateRequest,
    responseType: todo_pb.TodoResponse,
    requestSerialize: (arg: todo_pb.TodoUpdateRequest) => Buffer;
    requestDeserialize: (buffer: Uint8Array) => todo_pb.TodoUpdateRequest;
    responseSerialize: (arg: todo_pb.TodoResponse) => Buffer;
    responseDeserialize: (buffer: Uint8Array) => todo_pb.TodoResponse;
}

export interface ITodoClient {
    getTodos(request: todo_pb.TodoFilter, metadata?: grpc.Metadata): grpc.ClientReadableStream;
    createTodo(request: todo_pb.TodoCreateRequest, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
    createTodo(request: todo_pb.TodoCreateRequest, metadata: grpc.Metadata, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
    updateTodo(request: todo_pb.TodoUpdateRequest, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
    updateTodo(request: todo_pb.TodoUpdateRequest, metadata: grpc.Metadata, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
}

export const TodoService: ITodoService;
export class TodoClient extends grpc.Client {
    constructor(address: string, credentials: any, options?: object);
    public getTodos(request: todo_pb.TodoFilter, metadata?: grpc.Metadata): grpc.ClientReadableStream;
    public createTodo(request: todo_pb.TodoCreateRequest, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
    public createTodo(request: todo_pb.TodoCreateRequest, metadata: grpc.Metadata, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
    public updateTodo(request: todo_pb.TodoUpdateRequest, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
    public updateTodo(request: todo_pb.TodoUpdateRequest, metadata: grpc.Metadata, callback: (error: Error | null, response: todo_pb.TodoResponse) => void): grpc.ClientUnaryCall;
}
