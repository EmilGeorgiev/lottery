/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { FinishedLottery } from "./finished_lottery";
import { Lottery } from "./lottery";
import { Params } from "./params";
import { SystemInfo } from "./system_info";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

/** GenesisState defines the lottery module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  lottery: Lottery | undefined;
  systemInfo:
    | SystemInfo
    | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  finishedLotteryList: FinishedLottery[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, lottery: undefined, systemInfo: undefined, finishedLotteryList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.lottery !== undefined) {
      Lottery.encode(message.lottery, writer.uint32(18).fork()).ldelim();
    }
    if (message.systemInfo !== undefined) {
      SystemInfo.encode(message.systemInfo, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.finishedLotteryList) {
      FinishedLottery.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.lottery = Lottery.decode(reader, reader.uint32());
          break;
        case 3:
          message.systemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        case 4:
          message.finishedLotteryList.push(FinishedLottery.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      lottery: isSet(object.lottery) ? Lottery.fromJSON(object.lottery) : undefined,
      systemInfo: isSet(object.systemInfo) ? SystemInfo.fromJSON(object.systemInfo) : undefined,
      finishedLotteryList: Array.isArray(object?.finishedLotteryList)
        ? object.finishedLotteryList.map((e: any) => FinishedLottery.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.lottery !== undefined && (obj.lottery = message.lottery ? Lottery.toJSON(message.lottery) : undefined);
    message.systemInfo !== undefined
      && (obj.systemInfo = message.systemInfo ? SystemInfo.toJSON(message.systemInfo) : undefined);
    if (message.finishedLotteryList) {
      obj.finishedLotteryList = message.finishedLotteryList.map((e) => e ? FinishedLottery.toJSON(e) : undefined);
    } else {
      obj.finishedLotteryList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.lottery = (object.lottery !== undefined && object.lottery !== null)
      ? Lottery.fromPartial(object.lottery)
      : undefined;
    message.systemInfo = (object.systemInfo !== undefined && object.systemInfo !== null)
      ? SystemInfo.fromPartial(object.systemInfo)
      : undefined;
    message.finishedLotteryList = object.finishedLotteryList?.map((e) => FinishedLottery.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
