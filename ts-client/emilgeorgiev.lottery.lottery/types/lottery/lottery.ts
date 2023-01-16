/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

export interface Lottery {
  enterLotteryTxs: EnterLotteryTx[];
}

export interface EnterLotteryTx {
  userAddress: string;
  bet: number;
  denom: string;
  datetime: string;
}

function createBaseLottery(): Lottery {
  return { enterLotteryTxs: [] };
}

export const Lottery = {
  encode(message: Lottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.enterLotteryTxs) {
      EnterLotteryTx.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Lottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.enterLotteryTxs.push(EnterLotteryTx.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Lottery {
    return {
      enterLotteryTxs: Array.isArray(object?.enterLotteryTxs)
        ? object.enterLotteryTxs.map((e: any) => EnterLotteryTx.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Lottery): unknown {
    const obj: any = {};
    if (message.enterLotteryTxs) {
      obj.enterLotteryTxs = message.enterLotteryTxs.map((e) => e ? EnterLotteryTx.toJSON(e) : undefined);
    } else {
      obj.enterLotteryTxs = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Lottery>, I>>(object: I): Lottery {
    const message = createBaseLottery();
    message.enterLotteryTxs = object.enterLotteryTxs?.map((e) => EnterLotteryTx.fromPartial(e)) || [];
    return message;
  },
};

function createBaseEnterLotteryTx(): EnterLotteryTx {
  return { userAddress: "", bet: 0, denom: "", datetime: "" };
}

export const EnterLotteryTx = {
  encode(message: EnterLotteryTx, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userAddress !== "") {
      writer.uint32(10).string(message.userAddress);
    }
    if (message.bet !== 0) {
      writer.uint32(16).uint64(message.bet);
    }
    if (message.denom !== "") {
      writer.uint32(26).string(message.denom);
    }
    if (message.datetime !== "") {
      writer.uint32(34).string(message.datetime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EnterLotteryTx {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnterLotteryTx();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAddress = reader.string();
          break;
        case 2:
          message.bet = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.denom = reader.string();
          break;
        case 4:
          message.datetime = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EnterLotteryTx {
    return {
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
      bet: isSet(object.bet) ? Number(object.bet) : 0,
      denom: isSet(object.denom) ? String(object.denom) : "",
      datetime: isSet(object.datetime) ? String(object.datetime) : "",
    };
  },

  toJSON(message: EnterLotteryTx): unknown {
    const obj: any = {};
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    message.bet !== undefined && (obj.bet = Math.round(message.bet));
    message.denom !== undefined && (obj.denom = message.denom);
    message.datetime !== undefined && (obj.datetime = message.datetime);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EnterLotteryTx>, I>>(object: I): EnterLotteryTx {
    const message = createBaseEnterLotteryTx();
    message.userAddress = object.userAddress ?? "";
    message.bet = object.bet ?? 0;
    message.denom = object.denom ?? "";
    message.datetime = object.datetime ?? "";
    return message;
  },
};

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
