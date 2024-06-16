import { useParams } from "@solidjs/router";
import useMessages from "hooks/swr/useMessages";
import { For } from "solid-js";

export default function SourcesId() {
    const params = useParams();
    const sourceId = () => params.id;

    const { data: messages } = useMessages(() => ({ sourceId: sourceId() }));

    return (
        <div class="flex flex-col gap-4 items-center justify-center m-10">
            <p>source: {sourceId()}</p>

            <For each={messages.v?.Data}>
                {msg => (
                    <div class="self-stretch border border-black">
                        <p class="font-mono">{msg.Raw}</p>
                    </div>
                )}
            </For>
        </div>
    );
}
