import useSWR from "solid-swr";
import { MessageDTO } from "types/dto/message";
import { PaginationDTO } from "types/dto/pagination";
import { SwrArg } from "types/swr";
import { createSwrKey } from "utils/swr";

type Arg = {
    sourceId: string | number;
};

export default function useMessages(arg: SwrArg<Arg>) {
    const key = createSwrKey("/sources/:sourceId/messages", arg);

    const swr = useSWR<PaginationDTO<MessageDTO[]>>(key);

    return swr;
}
