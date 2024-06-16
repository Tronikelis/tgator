import { useParams } from "@solidjs/router";

export default function SourceId() {
    const params = useParams();
    const sourceId = () => params.id;

    return (
        <div>
            <p>source: {sourceId()}</p>
        </div>
    );
}
