import { JSX } from "solid-js";

type Props = {
    message: string;
    highlight: string;
    render?: (message: string) => JSX.Element;
};
export default function HighlightMessage(props: Props) {
    const hlIndex = () => {
        if (!props.highlight) return -1;
        return props.message.indexOf(props.highlight);
    };

    const hlPart = () => {
        if (hlIndex() === -1) return;
        return props.render?.(props.highlight) || <span>{props.highlight}</span>;
    };

    const all = () => {
        const index = hlIndex();
        if (index === -1) return props.message;

        return [
            props.message.slice(0, index),
            hlPart(),
            props.message.slice(index + props.highlight.length),
        ];
    };

    return <>{all()}</>;
}
