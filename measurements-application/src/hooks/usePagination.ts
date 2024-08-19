import {useState} from "react";

type usePaginationReturn = {
    page: number;
    pageSize: number;
    setPage: (page: number) => void;
    setPageSize: (pageSize: number) => void;
}

export const usePagination = (): usePaginationReturn => {
    const [page, setPage] = useState(1);
    const [pageSize, setPageSize] = useState(50);

    return { page, pageSize, setPage, setPageSize }
}