export type PagedList<T> = {
    page: number
    pageSize: number
    total: number
    items: T[]
}