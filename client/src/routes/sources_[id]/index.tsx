import { Button, Card, Group, Input, Loading, Pagination, Stack, Text } from "solid-daisy";
import { For, createSignal } from "solid-js";
import { useParams } from "@solidjs/router";

import useMessages from "hooks/swr/useMessages";
import useSource from "hooks/swr/useSource";
import usePage from "hooks/usePage";
import safeJsonPretty from "utils/safeJsonPretty";
import useDebouncedValue from "hooks/useDebouncedValue";

import HighlightMessage from "./components/HighlightMessage";

export default function SourcesId() {
    const params = useParams();
    const sourceId = () => params.id;

    const [search, setSearch] = createSignal("");
    const [page, setPage] = usePage([search]);
    const [orderBy, setOrderBy] = createSignal<"desc" | "asc">("desc");

    const debouncedSearch = useDebouncedValue(search, () => 200);

    const [messages] = useMessages(
        () => {
            const id = sourceId();
            if (!id) return;
            return {
                sourceId: id,
                orderBy: orderBy(),
                page: page(),
                search: debouncedSearch(),
            };
        },
        {
            refreshInterval: 5e3,
            keepPreviousData: true,
        }
    );

    const [{ data: source }] = useSource(() => {
        const id = sourceId();
        if (!id) return;
        return { id };
    });

    const formatDate = (date: string): string => {
        const d = new Date(date);

        return d
            .toLocaleString("lt", {
                timeStyle: "short",
                dateStyle: "short",
            })
            .split(" ")
            .reverse()
            .join(" ");
    };

    const onClickOrderBy = () => {
        const n = orderBy() === "asc" ? "desc" : "asc";
        setOrderBy(n);
    };

    return (
        <Stack class="gap-4">
            <Card>
                <Text>
                    <span class="font-bold">[{source.v?.ID}] </span>
                    Source: {source.v?.Name}
                </Text>
            </Card>

            <Input
                wrapperProps={{ class: "flex-1" }}
                placeholder="Search"
                bordered
                value={search()}
                onInput={e => setSearch(e.target.value)}
                rightSection={messages.isLoading() && <Loading />}
            />

            <Group>
                <Button onClick={onClickOrderBy}>{orderBy().toUpperCase()}</Button>

                <Pagination
                    value={page()}
                    setValue={setPage}
                    total={messages.data.v?.Pages || 0}
                />
            </Group>

            <Stack class="gap-0">
                <For each={messages.data.v?.Data}>
                    {msg => (
                        <Card class="rounded-none">
                            <Group>
                                <Text class="flex-1 font-mono">
                                    <pre>
                                        <HighlightMessage
                                            highlight={search()}
                                            message={safeJsonPretty(msg.Raw)}
                                            render={x => (
                                                <span class="font-bold text-red-600">{x}</span>
                                            )}
                                        />
                                    </pre>
                                </Text>

                                <Text italic>{formatDate(msg.CreatedAt)}</Text>
                            </Group>
                        </Card>
                    )}
                </For>
            </Stack>
        </Stack>
    );
}
