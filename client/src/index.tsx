import { render } from "solid-js/web";
import { Router, Route } from "@solidjs/router";
import { SWROptionsProvider } from "solid-swr";

import Idx from "routes/idx";
import SourceId from "routes/sources_[id]";
import Root from "components/Root";

import { api } from "utils/api";

import "./main.css";

function App() {
    return (
        <SWROptionsProvider value={{ fetcher: key => api.get(key).then(x => x.data) }}>
            <Root>
                <Router>
                    <Route path="/" component={Idx} />
                    <Route path="/sources/:id" component={SourceId} />
                </Router>
            </Root>
        </SWROptionsProvider>
    );
}

render(() => <App />, document.getElementById("solid-root")!);
