import { JSX, createMemo } from "solid-js";

type Props = {
    message: string;
    highlight: string;
    render?: (message: string) => JSX.Element;
};

export default function HighlightMessage(props: Props) {
    const renderHlPart = (index: number) => {
        const highlight = props.message.slice(index, index + props.highlight.length);
        return props.render?.(highlight) || <span>{highlight}</span>;
    };

    const hlIndexes = createMemo<number[]>(() => {
        const message = props.message.toLowerCase();
        const highlight = props.highlight.toLowerCase();

        if (!message || !highlight) return [];

        const indexes: number[] = [];
        // so first position is 0
        let index = -highlight.length;

        while ((index = message.indexOf(highlight, index + highlight.length)) !== -1) {
            indexes.push(index);
        }

        return indexes;
    });

    const all = createMemo(() => {
        const indexes = hlIndexes();
        if (indexes.length === 0) return props.message;

        const final: JSX.Element[] = [];
        let tmp = "";
        let ptr = 0;

        for (let i = 0; i < props.message.length; i++) {
            const hlIndex = indexes[ptr];

            if (hlIndex === i) {
                final.push(tmp, renderHlPart(hlIndex));
                i += props.highlight.length - 1;

                ptr++;
                tmp = "";

                continue;
            }

            tmp += props.message[i];
        }

        final.push(tmp);

        return final;
    });

    return <>{all()}</>;
}
