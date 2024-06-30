export type PaginationDTO<T extends any[]> = {
    Page: number;
    Offset: number;
    Limit: number;
    Pages: number;
    Count: number;
    Data: T;
};
