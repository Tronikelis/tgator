import { Accessor, createEffect, on } from "solid-js";

import useUrlSignal from "./useUrlSignal";

export default function usePage(deps: Accessor<any>[] = []) {
    const [value, setValue] = useUrlSignal<number>({
        key: "page",
        def: 0,
        fromQuery: q => parseInt(q),
    });

    createEffect(
        on(deps, () => {
            setValue(0);
        })
    );

    return [value, setValue] as const;
}
