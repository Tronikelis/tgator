import useSWR from "solid-swr";
import { SourceDTO } from "types/dto";
import { SwrArg } from "types/swr";
import { createSwrKey } from "utils/swr";

type Arg = {
    id: string | number;
};

export default function useSource(arg: SwrArg<Arg>) {
    const key = createSwrKey("/sources/:id", arg);

    const swr = useSWR<SourceDTO>(key);

    return [swr] as const;
}
