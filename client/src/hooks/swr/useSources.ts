import useSWR from "solid-swr";

export default function useSources() {
    const swr = useSWR(() => "/sources");
    return swr;
}
