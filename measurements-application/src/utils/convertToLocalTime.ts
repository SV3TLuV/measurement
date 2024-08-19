import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import customParseFormat from "dayjs/plugin/customParseFormat";

dayjs.extend(utc);
dayjs.extend(customParseFormat);

export const convertToLocalTime = (datetime: string): dayjs.Dayjs => {
    console.log(dayjs.utc(datetime).local(), datetime)

    return dayjs.utc(datetime).local()
}