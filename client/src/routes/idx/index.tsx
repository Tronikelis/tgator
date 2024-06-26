import { For, createSignal } from "solid-js";
import urlbat from "urlbat";

import useSources from "hooks/swr/useSources";
import { Button, Group, Input, Link, Stack } from "solid-daisy";

export default function Idx() {
    const [{ data }, { create }] = useSources();

    const [name, setName] = createSignal("");

    const onCreateSource = async () => {
        await create.trigger({ name: name() });
        create.populateCache();
    };

    return (
        <Stack>
            <Group>
                <Input
                    placeholder="Name"
                    bordered
                    value={name()}
                    onInput={e => setName(e.target.value)}
                />
                <Button onClick={onCreateSource}>Create</Button>
            </Group>

            <Group class="gap-8">
                <For each={data.v}>
                    {source => (
                        <div>
                            <Link
                                href={urlbat("/sources/:id", { id: source.ID })}
                                class="font-mono"
                            >
                                {source.Name}
                            </Link>
                        </div>
                    )}
                </For>
            </Group>
        </Stack>
    );
}
