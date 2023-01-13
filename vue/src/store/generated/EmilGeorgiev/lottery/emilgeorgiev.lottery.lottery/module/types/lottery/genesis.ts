/* eslint-disable */
import { Params } from "../lottery/params";
import { Lottery } from "../lottery/lottery";
import { SystemInfo } from "../lottery/system_info";
import { FinishedLottery } from "../lottery/finished_lottery";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "emilgeorgiev.lottery.lottery";

/** GenesisState defines the lottery module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  lottery: Lottery | undefined;
  systemInfo: SystemInfo | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  finishedLotteryList: FinishedLottery[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.finishedLotteryList = [];
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
          message.finishedLotteryList.push(
            FinishedLottery.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.finishedLotteryList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.lottery !== undefined && object.lottery !== null) {
      message.lottery = Lottery.fromJSON(object.lottery);
    } else {
      message.lottery = undefined;
    }
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      message.systemInfo = SystemInfo.fromJSON(object.systemInfo);
    } else {
      message.systemInfo = undefined;
    }
    if (
      object.finishedLotteryList !== undefined &&
      object.finishedLotteryList !== null
    ) {
      for (const e of object.finishedLotteryList) {
        message.finishedLotteryList.push(FinishedLottery.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.lottery !== undefined &&
      (obj.lottery = message.lottery
        ? Lottery.toJSON(message.lottery)
        : undefined);
    message.systemInfo !== undefined &&
      (obj.systemInfo = message.systemInfo
        ? SystemInfo.toJSON(message.systemInfo)
        : undefined);
    if (message.finishedLotteryList) {
      obj.finishedLotteryList = message.finishedLotteryList.map((e) =>
        e ? FinishedLottery.toJSON(e) : undefined
      );
    } else {
      obj.finishedLotteryList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.finishedLotteryList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.lottery !== undefined && object.lottery !== null) {
      message.lottery = Lottery.fromPartial(object.lottery);
    } else {
      message.lottery = undefined;
    }
    if (object.systemInfo !== undefined && object.systemInfo !== null) {
      message.systemInfo = SystemInfo.fromPartial(object.systemInfo);
    } else {
      message.systemInfo = undefined;
    }
    if (
      object.finishedLotteryList !== undefined &&
      object.finishedLotteryList !== null
    ) {
      for (const e of object.finishedLotteryList) {
        message.finishedLotteryList.push(FinishedLottery.fromPartial(e));
      }
    }
    return message;
  },
};

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
