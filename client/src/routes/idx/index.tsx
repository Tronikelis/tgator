import useSWR from "solid-swr";

export default function Idx() {
    const { data } = useSWR(() => "/sources");
    return (
        <div class="grid gap-8 grid-cols-4">

        </div>
    )
}
