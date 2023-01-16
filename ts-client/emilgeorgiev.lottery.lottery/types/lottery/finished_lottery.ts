/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { EnterLotteryTx } from "./lottery";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

export interface FinishedLottery {
  index: string;
  winner: string;
  reward: number;
  enterLotteryTxs: EnterLotteryTx[];
  winnerIndex: number;
}

function createBaseFinishedLottery(): FinishedLottery {
  return { index: "", winner: "", reward: 0, enterLotteryTxs: [], winnerIndex: 0 };
}

export const FinishedLottery = {
  encode(message: FinishedLottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.winner !== "") {
      writer.uint32(18).string(message.winner);
    }
    if (message.reward !== 0) {
      writer.uint32(24).uint64(message.reward);
    }
    for (const v of message.enterLotteryTxs) {
      EnterLotteryTx.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.winnerIndex !== 0) {
      writer.uint32(40).uint64(message.winnerIndex);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FinishedLottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFinishedLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.winner = reader.string();
          break;
        case 3:
          message.reward = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.enterLotteryTxs.push(EnterLotteryTx.decode(reader, reader.uint32()));
          break;
        case 5:
          message.winnerIndex = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FinishedLottery {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      winner: isSet(object.winner) ? String(object.winner) : "",
      reward: isSet(object.reward) ? Number(object.reward) : 0,
      enterLotteryTxs: Array.isArray(object?.enterLotteryTxs)
        ? object.enterLotteryTxs.map((e: any) => EnterLotteryTx.fromJSON(e))
        : [],
      winnerIndex: isSet(object.winnerIndex) ? Number(object.winnerIndex) : 0,
    };
  },

  toJSON(message: FinishedLottery): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.winner !== undefined && (obj.winner = message.winner);
    message.reward !== undefined && (obj.reward = Math.round(message.reward));
    if (message.enterLotteryTxs) {
      obj.enterLotteryTxs = message.enterLotteryTxs.map((e) => e ? EnterLotteryTx.toJSON(e) : undefined);
    } else {
      obj.enterLotteryTxs = [];
    }
    message.winnerIndex !== undefined && (obj.winnerIndex = Math.round(message.winnerIndex));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FinishedLottery>, I>>(object: I): FinishedLottery {
    const message = createBaseFinishedLottery();
    message.index = object.index ?? "";
    message.winner = object.winner ?? "";
    message.reward = object.reward ?? 0;
    message.enterLotteryTxs = object.enterLotteryTxs?.map((e) => EnterLotteryTx.fromPartial(e)) || [];
    message.winnerIndex = object.winnerIndex ?? 0;
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
