import { For, createSignal } from "solid-js";
import urlbat from "urlbat";

import useSources from "hooks/swr/useSources";
import { Button, Group, Input, Link, Stack } from "solid-daisy";

export default function Idx() {
    const [{ data }, { create }] = useSources();

    const [sourceIp, setSourceIp] = createSignal("");

    const onCreateSource = async () => {
        await create.trigger({ ip: sourceIp() });
        create.populateCache();
    };

    return (
        <Stack>
            <Group>
                <Input
                    bordered
                    value={sourceIp()}
                    onInput={e => setSourceIp(e.target.value)}
                />
                <Button onClick={onCreateSource}>create source</Button>
            </Group>

            <Group class="gap-8">
                <For each={data.v}>
                    {source => (
                        <div>
                            <Link
                                href={urlbat("/sources/:id", { id: source.ID })}
                                class="font-mono"
                            >
                                {source.Ip}
                            </Link>
                        </div>
                    )}
                </For>
            </Group>
        </Stack>
    );
}
