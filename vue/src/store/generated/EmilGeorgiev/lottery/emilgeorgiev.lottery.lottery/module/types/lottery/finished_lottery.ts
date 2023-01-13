/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

export interface FinishedLottery {
  index: string;
  winner: string;
  payment: number;
}

const baseFinishedLottery: object = { index: "", winner: "", payment: 0 };

export const FinishedLottery = {
  encode(message: FinishedLottery, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.winner !== "") {
      writer.uint32(18).string(message.winner);
    }
    if (message.payment !== 0) {
      writer.uint32(24).uint64(message.payment);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): FinishedLottery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseFinishedLottery } as FinishedLottery;
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
          message.payment = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FinishedLottery {
    const message = { ...baseFinishedLottery } as FinishedLottery;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    if (object.payment !== undefined && object.payment !== null) {
      message.payment = Number(object.payment);
    } else {
      message.payment = 0;
    }
    return message;
  },

  toJSON(message: FinishedLottery): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.winner !== undefined && (obj.winner = message.winner);
    message.payment !== undefined && (obj.payment = message.payment);
    return obj;
  },

  fromPartial(object: DeepPartial<FinishedLottery>): FinishedLottery {
    const message = { ...baseFinishedLottery } as FinishedLottery;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    if (object.payment !== undefined && object.payment !== null) {
      message.payment = object.payment;
    } else {
      message.payment = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
