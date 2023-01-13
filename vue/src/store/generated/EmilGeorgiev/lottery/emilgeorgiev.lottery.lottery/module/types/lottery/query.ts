/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../lottery/params";
import { Lottery } from "../lottery/lottery";
import { SystemInfo } from "../lottery/system_info";
import { FinishedLottery } from "../lottery/finished_lottery";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetLotteryRequest {}

export interface QueryGetLotteryResponse {
  Lottery: Lottery | undefined;
}

export interface QueryGetSystemInfoRequest {}

export interface QueryGetSystemInfoResponse {
  SystemInfo: SystemInfo | undefined;
}

export interface QueryGetFinishedLotteryRequest {
  index: string;
}

export interface QueryGetFinishedLotteryResponse {
  finishedLottery: FinishedLottery | undefined;
}

export interface QueryAllFinishedLotteryRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFinishedLotteryResponse {
  finishedLottery: FinishedLottery[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetLotteryRequest: object = {};

export const QueryGetLotteryRequest = {
  encode(_: QueryGetLotteryRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLotteryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetLotteryRequest } as QueryGetLotteryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetLotteryRequest {
    const message = { ...baseQueryGetLotteryRequest } as QueryGetLotteryRequest;
    return message;
  },

  toJSON(_: QueryGetLotteryRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryGetLotteryRequest>): QueryGetLotteryRequest {
    const message = { ...baseQueryGetLotteryRequest } as QueryGetLotteryRequest;
    return message;
  },
};

const baseQueryGetLotteryResponse: object = {};

export const QueryGetLotteryResponse = {
  encode(
    message: QueryGetLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Lottery !== undefined) {
      Lottery.encode(message.Lottery, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLotteryResponse,
    } as QueryGetLotteryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Lottery = Lottery.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLotteryResponse {
    const message = {
      ...baseQueryGetLotteryResponse,
    } as QueryGetLotteryResponse;
    if (object.Lottery !== undefined && object.Lottery !== null) {
      message.Lottery = Lottery.fromJSON(object.Lottery);
    } else {
      message.Lottery = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLotteryResponse): unknown {
    const obj: any = {};
    message.Lottery !== undefined &&
      (obj.Lottery = message.Lottery
        ? Lottery.toJSON(message.Lottery)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLotteryResponse>
  ): QueryGetLotteryResponse {
    const message = {
      ...baseQueryGetLotteryResponse,
    } as QueryGetLotteryResponse;
    if (object.Lottery !== undefined && object.Lottery !== null) {
      message.Lottery = Lottery.fromPartial(object.Lottery);
    } else {
      message.Lottery = undefined;
    }
    return message;
  },
};

const baseQueryGetSystemInfoRequest: object = {};

export const QueryGetSystemInfoRequest = {
  encode(
    _: QueryGetSystemInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },

  toJSON(_: QueryGetSystemInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetSystemInfoRequest>
  ): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },
};

const baseQueryGetSystemInfoResponse: object = {};

export const QueryGetSystemInfoResponse = {
  encode(
    message: QueryGetSystemInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SystemInfo !== undefined) {
      SystemInfo.encode(message.SystemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SystemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromJSON(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSystemInfoResponse): unknown {
    const obj: any = {};
    message.SystemInfo !== undefined &&
      (obj.SystemInfo = message.SystemInfo
        ? SystemInfo.toJSON(message.SystemInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSystemInfoResponse>
  ): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromPartial(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetFinishedLotteryRequest: object = { index: "" };

export const QueryGetFinishedLotteryRequest = {
  encode(
    message: QueryGetFinishedLotteryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetFinishedLotteryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFinishedLotteryRequest,
    } as QueryGetFinishedLotteryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFinishedLotteryRequest {
    const message = {
      ...baseQueryGetFinishedLotteryRequest,
    } as QueryGetFinishedLotteryRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetFinishedLotteryRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFinishedLotteryRequest>
  ): QueryGetFinishedLotteryRequest {
    const message = {
      ...baseQueryGetFinishedLotteryRequest,
    } as QueryGetFinishedLotteryRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetFinishedLotteryResponse: object = {};

export const QueryGetFinishedLotteryResponse = {
  encode(
    message: QueryGetFinishedLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.finishedLottery !== undefined) {
      FinishedLottery.encode(
        message.finishedLottery,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetFinishedLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetFinishedLotteryResponse,
    } as QueryGetFinishedLotteryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.finishedLottery = FinishedLottery.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetFinishedLotteryResponse {
    const message = {
      ...baseQueryGetFinishedLotteryResponse,
    } as QueryGetFinishedLotteryResponse;
    if (
      object.finishedLottery !== undefined &&
      object.finishedLottery !== null
    ) {
      message.finishedLottery = FinishedLottery.fromJSON(
        object.finishedLottery
      );
    } else {
      message.finishedLottery = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetFinishedLotteryResponse): unknown {
    const obj: any = {};
    message.finishedLottery !== undefined &&
      (obj.finishedLottery = message.finishedLottery
        ? FinishedLottery.toJSON(message.finishedLottery)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetFinishedLotteryResponse>
  ): QueryGetFinishedLotteryResponse {
    const message = {
      ...baseQueryGetFinishedLotteryResponse,
    } as QueryGetFinishedLotteryResponse;
    if (
      object.finishedLottery !== undefined &&
      object.finishedLottery !== null
    ) {
      message.finishedLottery = FinishedLottery.fromPartial(
        object.finishedLottery
      );
    } else {
      message.finishedLottery = undefined;
    }
    return message;
  },
};

const baseQueryAllFinishedLotteryRequest: object = {};

export const QueryAllFinishedLotteryRequest = {
  encode(
    message: QueryAllFinishedLotteryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllFinishedLotteryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllFinishedLotteryRequest,
    } as QueryAllFinishedLotteryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFinishedLotteryRequest {
    const message = {
      ...baseQueryAllFinishedLotteryRequest,
    } as QueryAllFinishedLotteryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFinishedLotteryRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFinishedLotteryRequest>
  ): QueryAllFinishedLotteryRequest {
    const message = {
      ...baseQueryAllFinishedLotteryRequest,
    } as QueryAllFinishedLotteryRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllFinishedLotteryResponse: object = {};

export const QueryAllFinishedLotteryResponse = {
  encode(
    message: QueryAllFinishedLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.finishedLottery) {
      FinishedLottery.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllFinishedLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllFinishedLotteryResponse,
    } as QueryAllFinishedLotteryResponse;
    message.finishedLottery = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.finishedLottery.push(
            FinishedLottery.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllFinishedLotteryResponse {
    const message = {
      ...baseQueryAllFinishedLotteryResponse,
    } as QueryAllFinishedLotteryResponse;
    message.finishedLottery = [];
    if (
      object.finishedLottery !== undefined &&
      object.finishedLottery !== null
    ) {
      for (const e of object.finishedLottery) {
        message.finishedLottery.push(FinishedLottery.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllFinishedLotteryResponse): unknown {
    const obj: any = {};
    if (message.finishedLottery) {
      obj.finishedLottery = message.finishedLottery.map((e) =>
        e ? FinishedLottery.toJSON(e) : undefined
      );
    } else {
      obj.finishedLottery = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllFinishedLotteryResponse>
  ): QueryAllFinishedLotteryResponse {
    const message = {
      ...baseQueryAllFinishedLotteryResponse,
    } as QueryAllFinishedLotteryResponse;
    message.finishedLottery = [];
    if (
      object.finishedLottery !== undefined &&
      object.finishedLottery !== null
    ) {
      for (const e of object.finishedLottery) {
        message.finishedLottery.push(FinishedLottery.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Lottery by index. */
  Lottery(request: QueryGetLotteryRequest): Promise<QueryGetLotteryResponse>;
  /** Queries a SystemInfo by index. */
  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse>;
  /** Queries a FinishedLottery by index. */
  FinishedLottery(
    request: QueryGetFinishedLotteryRequest
  ): Promise<QueryGetFinishedLotteryResponse>;
  /** Queries a list of FinishedLottery items. */
  FinishedLotteryAll(
    request: QueryAllFinishedLotteryRequest
  ): Promise<QueryAllFinishedLotteryResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "emilgeorgiev.lottery.lottery.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Lottery(request: QueryGetLotteryRequest): Promise<QueryGetLotteryResponse> {
    const data = QueryGetLotteryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "emilgeorgiev.lottery.lottery.Query",
      "Lottery",
      data
    );
    return promise.then((data) =>
      QueryGetLotteryResponse.decode(new Reader(data))
    );
  }

  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse> {
    const data = QueryGetSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "emilgeorgiev.lottery.lottery.Query",
      "SystemInfo",
      data
    );
    return promise.then((data) =>
      QueryGetSystemInfoResponse.decode(new Reader(data))
    );
  }

  FinishedLottery(
    request: QueryGetFinishedLotteryRequest
  ): Promise<QueryGetFinishedLotteryResponse> {
    const data = QueryGetFinishedLotteryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "emilgeorgiev.lottery.lottery.Query",
      "FinishedLottery",
      data
    );
    return promise.then((data) =>
      QueryGetFinishedLotteryResponse.decode(new Reader(data))
    );
  }

  FinishedLotteryAll(
    request: QueryAllFinishedLotteryRequest
  ): Promise<QueryAllFinishedLotteryResponse> {
    const data = QueryAllFinishedLotteryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "emilgeorgiev.lottery.lottery.Query",
      "FinishedLotteryAll",
      data
    );
    return promise.then((data) =>
      QueryAllFinishedLotteryResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
