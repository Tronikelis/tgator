import useSWR from "solid-swr";
import { MessageDTO } from "types/dto/message";
import { PaginationDTO } from "types/dto/pagination";
import { SwrArg } from "types/swr";
import { createSwrKey } from "utils/swr";

type Arg = {
    sourceId: number;
};

export default function useMessages(arg: SwrArg<Arg>) {
    const key = createSwrKey("/messages/:sourceId", arg);

    const swr = useSWR<PaginationDTO<MessageDTO[]>>(key);

    return swr;
}
