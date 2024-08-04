import urlbat from "urlbat";
import { Link } from "solid-daisy";
import { createMemo } from "solid-js";
import { useParams } from "@solidjs/router";

import safeJsonPretty from "utils/safeJsonPretty";
import { ChunkNodeRenderer } from "utils/classes/ChunkNode";

type Props = {
    message: string;
    highlight: string;
};

function matchingStrings(parent: string, target: string): [number, number][] {
    let index = -target.length;

    const matched: [number, number][] = [];

    while ((index = parent.indexOf(target, index + target.length)) !== -1) {
        matched.push([index, index + target.length]);
    }

    return matched;
}

export default function Message(props: Props) {
    const params = useParams();

    const message = () => safeJsonPretty(props.message);

    const rendered = createMemo(() => {
        const m = message();

        const renderer = new ChunkNodeRenderer(m);

        const quotes = matchingStrings(m, '"');

        for (let i = 0; i < quotes.length; i += 2) {
            const from = quotes[i]?.[0];
            const to = quotes[i + 1]?.[0];

            if (from != null && to != null) {
                renderer.addNode(
                    (x, raw) => (
                        <Link
                            href={urlbat("/sources/:id", { id: params.id, search: raw })}
                            class="underline"
                        >
                            {x}
                        </Link>
                    ),
                    from + 1,
                    to
                );
            }
        }

        if (props.highlight) {
            for (const hl of matchingStrings(m, props.highlight)) {
                renderer.addNode(x => <span class="text-red-600">{x}</span>, ...hl);
            }
        }

        return renderer.render();
    });

    return <pre>{rendered()}</pre>;
}
