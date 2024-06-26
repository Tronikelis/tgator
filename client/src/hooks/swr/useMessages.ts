import useSWR, { Options } from "solid-swr";

import { SourceDTO, PaginationDTO, MessageDTO } from "types/dto";
import { SwrArg } from "types/swr";
import { createSwrKey } from "utils/swr";

type Arg = {
    page?: number;
    sourceId: string | number;
    search?: string;
    orderBy?: "asc" | "desc";
};

type Res = PaginationDTO<(MessageDTO & { Source: SourceDTO })[]>;

export default function useMessages(arg: SwrArg<Arg>, options?: Options<Res, unknown>) {
    const key = createSwrKey("/sources/:sourceId/messages", arg);

    const swr = useSWR<Res>(key, options);

    return [swr] as const;
}
