import {SchedulerStatus} from "./schedulerStatus.ts";

export type PollingInformation = {
    status: SchedulerStatus,
    percent: number
}