import useSWR from "solid-swr";
import { SourceDTO } from "types/dto";

export default function useSources() {
    const swr = useSWR<SourceDTO[] | null>(() => "/sources");
    return swr;
}
