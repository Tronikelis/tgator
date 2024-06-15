import useSWR from "solid-swr";

export default function Idx() {
    const { data } = useSWR(() => "/messages");
    return <pre>{JSON.stringify(data.v, null, 2)}</pre>;
}
