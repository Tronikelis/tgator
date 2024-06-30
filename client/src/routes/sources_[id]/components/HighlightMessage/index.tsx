import { For, JSX, createMemo } from "solid-js";

type Props = {
    message: string;
    highlight: string;
    render?: (message: string) => JSX.Element;
};

export default function HighlightMessage(props: Props) {
    const elements = createMemo<JSX.Element[]>(() => {
        const message = props.message.toLowerCase();
        const highlight = props.highlight.toLowerCase();

        if (!message || !highlight) return [props.message];

        // so first position is 0
        let index = -highlight.length;
        let lastIndexEnd = 0;

        const final: JSX.Element[] = [];

        while ((index = message.indexOf(highlight, index + highlight.length)) !== -1) {
            // add message in between last index and current index
            const msg = props.message.slice(lastIndexEnd, index);

            // add current index highlight
            const highlighted = props.message.slice(
                index,
                (lastIndexEnd = index + props.highlight.length)
            );

            final.push(msg, props.render?.(highlighted));
        }

        final.push(props.message.slice(lastIndexEnd));

        return final;
    });

    return <For each={elements()}>{item => item}</For>;
}
