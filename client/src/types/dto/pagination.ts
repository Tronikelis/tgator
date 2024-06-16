export type PaginationDTO<T extends any[]> = {
    Offset: number;
    Limit: number;
    Data: T;
};
