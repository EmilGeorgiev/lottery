/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

export interface MsgEnterLottery {
  creator: string;
  bet: number;
  denom: string;
}

export interface MsgEnterLotteryResponse {
  registeredUsers: number;
}

function createBaseMsgEnterLottery(): MsgEnterLottery {
  return { creator: "", bet: 0, denom: "" };
}

export const MsgEnterLottery = {
  encode(message: MsgEnterLottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.bet !== 0) {
      writer.uint32(16).uint64(message.bet);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnterLottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnterLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.bet = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnterLottery {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      bet: isSet(object.bet) ? Number(object.bet) : 0,
      denom: isSet(object.denom) ? String(object.denom) : "",
    };
  },

  toJSON(message: MsgEnterLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.bet !== undefined && (obj.bet = Math.round(message.bet));
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnterLottery>, I>>(object: I): MsgEnterLottery {
    const message = createBaseMsgEnterLottery();
    message.creator = object.creator ?? "";
    message.bet = object.bet ?? 0;
    message.denom = object.denom ?? "";
    return message;
  },
};

function createBaseMsgEnterLotteryResponse(): MsgEnterLotteryResponse {
  return { registeredUsers: 0 };
}

export const MsgEnterLotteryResponse = {
  encode(message: MsgEnterLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.registeredUsers !== 0) {
      writer.uint32(8).uint64(message.registeredUsers);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnterLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnterLotteryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registeredUsers = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnterLotteryResponse {
    return { registeredUsers: isSet(object.registeredUsers) ? Number(object.registeredUsers) : 0 };
  },

  toJSON(message: MsgEnterLotteryResponse): unknown {
    const obj: any = {};
    message.registeredUsers !== undefined && (obj.registeredUsers = Math.round(message.registeredUsers));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnterLotteryResponse>, I>>(object: I): MsgEnterLotteryResponse {
    const message = createBaseMsgEnterLotteryResponse();
    message.registeredUsers = object.registeredUsers ?? 0;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  EnterLottery(request: MsgEnterLottery): Promise<MsgEnterLotteryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.EnterLottery = this.EnterLottery.bind(this);
  }
  EnterLottery(request: MsgEnterLottery): Promise<MsgEnterLotteryResponse> {
    const data = MsgEnterLottery.encode(request).finish();
    const promise = this.rpc.request("emilgeorgiev.lottery.lottery.Msg", "EnterLottery", data);
    return promise.then((data) => MsgEnterLotteryResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
