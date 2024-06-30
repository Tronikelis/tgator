import { useBeforeLeave, useSearchParams } from "@solidjs/router";
import { createEffect, createSignal, on, onMount } from "solid-js";

type Arg<T> = {
    key: string;
    fromQuery: (query: string) => T;
    def: T;
};

export default function useUrlSignal<T extends string | number | boolean>({
    key,
    def,
    fromQuery,
}: Arg<T>) {
    const [params, setParams] = useSearchParams();

    const getParam = () => params[key];
    const setParam = (value: string) => setParams({ [key]: value });

    const [value, setValue] = createSignal<T>(def);

    onMount(() => {
        const param = getParam();
        if (param) setValue(() => fromQuery(param));
    });

    useBeforeLeave(ev => {
        if (typeof ev.to !== "string") return;

        const [, query] = ev.to.split("?");

        if (!query) {
            setValue(() => def);
            return;
        }

        if (!new URLSearchParams(query).get(key)) {
            setValue(() => def);
        }
    });

    createEffect(on(value, value => setParam(String(value)), { defer: true }));

    return [value, setValue] as const;
}
