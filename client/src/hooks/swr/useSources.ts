import useSWR, { useSWRMutation } from "solid-swr";
import { SourceDTO } from "types/dto";
import { SwrArg } from "types/swr";
import { api } from "utils/api";
import { createSwrKey } from "utils/swr";

export default function useSources() {
    const key = createSwrKey("/sources", () => ({}));

    const swr = useSWR<SourceDTO[] | null>(key);
    const actions = useActions(key);

    return [swr, actions] as const;
}

type CreateBody = {
    name: string;
};

function useActions(key: SwrArg<string>) {
    const create = useSWRMutation(
        k => k === key(),
        (arg: CreateBody) => api.post("/sources", arg).then(x => x.data)
    );

    return { create };
}
