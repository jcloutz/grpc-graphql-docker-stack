// package: todo
// file: todo.proto

import * as jspb from "google-protobuf";

export class TodoCreateRequest extends jspb.Message { 
    getName(): string;
    setName(value: string): void;

    getComplete(): boolean;
    setComplete(value: boolean): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TodoCreateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: TodoCreateRequest): TodoCreateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TodoCreateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TodoCreateRequest;
    static deserializeBinaryFromReader(message: TodoCreateRequest, reader: jspb.BinaryReader): TodoCreateRequest;
}

export namespace TodoCreateRequest {
    export type AsObject = {
        name: string,
        complete: boolean,
    }
}

export class TodoFilter extends jspb.Message { 
    getId(): string;
    setId(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TodoFilter.AsObject;
    static toObject(includeInstance: boolean, msg: TodoFilter): TodoFilter.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TodoFilter, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TodoFilter;
    static deserializeBinaryFromReader(message: TodoFilter, reader: jspb.BinaryReader): TodoFilter;
}

export namespace TodoFilter {
    export type AsObject = {
        id: string,
    }
}

export class TodoUpdateRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): void;

    getName(): string;
    setName(value: string): void;

    getComplete(): boolean;
    setComplete(value: boolean): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TodoUpdateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: TodoUpdateRequest): TodoUpdateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TodoUpdateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TodoUpdateRequest;
    static deserializeBinaryFromReader(message: TodoUpdateRequest, reader: jspb.BinaryReader): TodoUpdateRequest;
}

export namespace TodoUpdateRequest {
    export type AsObject = {
        id: string,
        name: string,
        complete: boolean,
    }
}

export class TodoResponse extends jspb.Message { 
    getId(): string;
    setId(value: string): void;

    getName(): string;
    setName(value: string): void;

    getComplete(): boolean;
    setComplete(value: boolean): void;

    getCreatedat(): number;
    setCreatedat(value: number): void;

    getUpdatedat(): number;
    setUpdatedat(value: number): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TodoResponse.AsObject;
    static toObject(includeInstance: boolean, msg: TodoResponse): TodoResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TodoResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TodoResponse;
    static deserializeBinaryFromReader(message: TodoResponse, reader: jspb.BinaryReader): TodoResponse;
}

export namespace TodoResponse {
    export type AsObject = {
        id: string,
        name: string,
        complete: boolean,
        createdat: number,
        updatedat: number,
    }
}
