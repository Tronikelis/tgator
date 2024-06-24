import { Accessor, createEffect, createSignal, on } from "solid-js";

export default function usePage(deps: Accessor<any>[] = []) {
    const [value, setValue] = createSignal(0);

    createEffect(
        on(deps, () => {
            setValue(0);
        })
    );

    return [value, setValue] as const;
}
