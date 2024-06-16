import { For } from "solid-js";
import urlbat from "urlbat";

import useSources from "hooks/swr/useSources";

export default function Idx() {
    const { data } = useSources();

    return (
        <div class="grid gap-8 grid-cols-4">
            <For each={data.v}>
                {source => (
                    <div>
                        <a href={urlbat("/sources/:id", { id: source.ID })} class="font-mono">
                            {source.Ip}
                        </a>
                    </div>
                )}
            </For>
        </div>
    );
}
