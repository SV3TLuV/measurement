import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import customParseFormat from "dayjs/plugin/customParseFormat";

dayjs.extend(utc);
dayjs.extend(customParseFormat);

export const convertToStringFormat = (datetime: dayjs.Dayjs): string => {
    return datetime.format('DD.MM.YYYY HH:mm')
}
