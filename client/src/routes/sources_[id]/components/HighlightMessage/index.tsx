import { For, JSX, createMemo } from "solid-js";

import useDebouncedValue from "hooks/useDebouncedValue";

type Props = {
    message: string;
    highlight: string;
    render?: (message: string) => JSX.Element;
};

export default function HighlightMessage(props: Props) {
    const message = () => props.message;

    const _highlight = () => props.highlight;
    const highlight = useDebouncedValue(_highlight);

    const elements = createMemo<JSX.Element[]>(() => {
        const m = message().toLowerCase();
        const h = highlight().toLowerCase();

        const mOrig = message();

        if (!m || !h) return [mOrig];

        // so first position is 0
        let index = -h.length;
        let lastIndexEnd = 0;

        const final: JSX.Element[] = [];

        while ((index = m.indexOf(h, index + h.length)) !== -1) {
            // add message in between last index and current index
            const msg = mOrig.slice(lastIndexEnd, index);

            // add current index highlight
            const highlighted = mOrig.slice(index, (lastIndexEnd = index + h.length));

            final.push(msg, props.render?.(highlighted));
        }

        final.push(mOrig.slice(lastIndexEnd));

        return final;
    });

    return <For each={elements()}>{item => item}</For>;
}
