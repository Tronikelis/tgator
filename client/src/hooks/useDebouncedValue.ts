import { Accessor, createEffect, createSignal, onCleanup } from "solid-js";

export default function useDebouncedValue<T>(
    value: Accessor<T>,
    ms: Accessor<number> = () => 400
) {
    const [debounced, setDebounced] = createSignal(value());

    createEffect(() => {
        const v = value();
        const cb = () => {
            setDebounced(() => v);
        };

        const timeout = setTimeout(cb, ms());
        onCleanup(() => {
            clearTimeout(timeout);
        });
    });

    return debounced;
}
