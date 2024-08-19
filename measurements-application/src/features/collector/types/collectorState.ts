export type CollectorStatus = "pending" | "collecting"

export type CollectorState = {
    status: CollectorStatus
    polledPostCount: number
    postCount: number
    pollingPercent: number
    receivedCount: number
    started: string
}