import {convertToStringFormat} from "../../../utils/convertToStringFormat.ts";
import {convertToLocalTime} from "../../../utils/convertToLocalTime.ts";

export const tickFormatter = (value: string) => convertToStringFormat(convertToLocalTime(value))