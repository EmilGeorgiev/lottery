/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

export interface MsgEnterLottery {
  creator: string;
  bet: number;
  denom: string;
}

export interface MsgEnterLotteryResponse {
  registered_users: number;
}

const baseMsgEnterLottery: object = { creator: "", bet: 0, denom: "" };

export const MsgEnterLottery = {
  encode(message: MsgEnterLottery, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgEnterLottery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEnterLottery } as MsgEnterLottery;
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
    const message = { ...baseMsgEnterLottery } as MsgEnterLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = Number(object.bet);
    } else {
      message.bet = 0;
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    return message;
  },

  toJSON(message: MsgEnterLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.bet !== undefined && (obj.bet = message.bet);
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEnterLottery>): MsgEnterLottery {
    const message = { ...baseMsgEnterLottery } as MsgEnterLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.bet !== undefined && object.bet !== null) {
      message.bet = object.bet;
    } else {
      message.bet = 0;
    }
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    return message;
  },
};

const baseMsgEnterLotteryResponse: object = { registered_users: 0 };

export const MsgEnterLotteryResponse = {
  encode(
    message: MsgEnterLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.registered_users !== 0) {
      writer.uint32(8).uint64(message.registered_users);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEnterLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgEnterLotteryResponse,
    } as MsgEnterLotteryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.registered_users = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnterLotteryResponse {
    const message = {
      ...baseMsgEnterLotteryResponse,
    } as MsgEnterLotteryResponse;
    if (
      object.registered_users !== undefined &&
      object.registered_users !== null
    ) {
      message.registered_users = Number(object.registered_users);
    } else {
      message.registered_users = 0;
    }
    return message;
  },

  toJSON(message: MsgEnterLotteryResponse): unknown {
    const obj: any = {};
    message.registered_users !== undefined &&
      (obj.registered_users = message.registered_users);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgEnterLotteryResponse>
  ): MsgEnterLotteryResponse {
    const message = {
      ...baseMsgEnterLotteryResponse,
    } as MsgEnterLotteryResponse;
    if (
      object.registered_users !== undefined &&
      object.registered_users !== null
    ) {
      message.registered_users = object.registered_users;
    } else {
      message.registered_users = 0;
    }
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
  }
  EnterLottery(request: MsgEnterLottery): Promise<MsgEnterLotteryResponse> {
    const data = MsgEnterLottery.encode(request).finish();
    const promise = this.rpc.request(
      "emilgeorgiev.lottery.lottery.Msg",
      "EnterLottery",
      data
    );
    return promise.then((data) =>
      MsgEnterLotteryResponse.decode(new Reader(data))
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
