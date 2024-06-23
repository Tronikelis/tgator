import { useParams } from "@solidjs/router";

import useMessages from "hooks/swr/useMessages";
import useSource from "hooks/swr/useSource";
import { Card, Stack, Text } from "solid-daisy";
import { For } from "solid-js";

export default function SourcesId() {
    const params = useParams();
    const sourceId = () => params.id;

    const [{ data: messages }] = useMessages(
        () => {
            const id = sourceId();
            if (!id) return;
            return { sourceId: id };
        },
        { refreshInterval: 5e3 }
    );

    const [{ data: source }] = useSource(() => {
        const id = sourceId();
        if (!id) return;
        return { id };
    });

    return (
        <Stack class="gap-4">
            <Card>
                <Stack>
                    <Text bold>Source</Text>
                    <Text>id: {source.v?.ID}</Text>
                    <Text>ip: {source.v?.Ip}</Text>
                </Stack>
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
