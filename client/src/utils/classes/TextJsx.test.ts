import { it, expect } from "vitest";
import stringify from "fast-json-stable-stringify";

import ChunkNodeTree from "./TextJsx";
import { children } from "solid-js";

const expectJson = (a: any, b: any) => {
    expect(stringify(a)).toBe(stringify(b));
};

it("works", () => {
    let chunkNodeTree = new ChunkNodeTree(undefined, 0, 10);

    let coords: [number, number][] = [[0, 5]];

    const reset = () => {
        chunkNodeTree = new ChunkNodeTree(undefined, 0, 10);
        coords.forEach(coord => chunkNodeTree.addNode(undefined, ...coord));
    };

    reset();

    expectJson(chunkNodeTree.root, {
        coords: [0, 10],
        children: [
            {
                coords: [0, 5],
                children: [],
            },
        ],
    });

    coords.push([0, 3], [0, 5]);
    reset();

    // expectJson(chunkNodeTree.root, {
    //     coords: [0, 10],
    //     children: [
    //         {
    //             coords: [0, 5],
    //             children: [
    //                 {
    //                     coords: [0, 5],
    //                     children: [
    //                         {
    //                             coords: [0, 3],
    //                             children: [],
    //                         },
    //                     ],
    //                 },
    //             ],
    //         },
    //     ],
    // });

    coords = [
        [0, 5],
        [3, 7],
    ];
    reset();

    expectJson(chunkNodeTree.root, {
        coords: [0, 10],
        children: [
            {
                coords: [0, 5],
                children: [
                    {
                        coords: [3, 5],
                        children: [],
                    },
                    {
                        coords: [0, 3],
                        children: [],
                    },
                ],
            },
        ],
    });
});
