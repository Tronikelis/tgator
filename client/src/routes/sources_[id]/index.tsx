import { Card, Group, Input, Loading, Pagination, Stack, Text } from "solid-daisy";
import { For, createSignal } from "solid-js";
import { useParams } from "@solidjs/router";

import useMessages from "hooks/swr/useMessages";
import useSource from "hooks/swr/useSource";
import usePage from "hooks/usePage";

export default function SourcesId() {
    const params = useParams();
    const sourceId = () => params.id;

    const [search, setSearch] = createSignal("");
    const [page, setPage] = usePage([search]);

    const [messages] = useMessages(
        () => {
            const id = sourceId();
            if (!id) return;
            return { sourceId: id, search: search(), page: page() };
        },
        { refreshInterval: 5e3, keepPreviousData: true }
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

            <Group>
                <Input
                    placeholder="Search"
                    bordered
                    value={search()}
                    onInput={e => setSearch(e.target.value)}
                />

                {messages.isLoading() && <Loading />}
            </Group>

            <Pagination
                value={page()}
                setValue={setPage}
                total={messages.data.v?.Pages || 0}
            />

            <Stack class="gap-0">
                <For each={messages.data.v?.Data}>
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
