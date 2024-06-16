import { useParams } from "@solidjs/router";
import useMessages from "hooks/swr/useMessages";
import { Card, Stack, Text } from "solid-daisy";
import { For } from "solid-js";

export default function SourcesId() {
    const params = useParams();
    const sourceId = () => params.id;

    const { data: messages } = useMessages(() => {
        const s = sourceId();
        if (!s) return;
        return { sourceId: s };
    });

    return (
        <Stack class="gap-4 p-12">
            <Card>
                <Text>source: {sourceId()}</Text>
            </Card>

            <Stack class="gap-0">
                <For each={messages.v?.Data}>
                    {msg => (
                        <Card class="rounded-none">
                            <Text class="font-mono">{msg.Raw}</Text>
                        </Card>
                    )}
                </For>
            </Stack>
        </Stack>
    );
}
