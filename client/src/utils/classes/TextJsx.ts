type Coords = [from: number, to: number];

type ChunkNode<E> = {
    coords: Coords;

    children: ChunkNode<E>[];

    extra: E;
};

export default class ChunkNodeTree<E = undefined> {
    root: ChunkNode<E>;

    private createNode(extra: E, from: number, to: number): ChunkNode<E> {
        return {
            extra,
            children: [],
            coords: [from, to],
        };
    }

    private splitNode(node: ChunkNode<E>, coords: Coords): [ChunkNode<E>, ...ChunkNode<E>[]] {
        if (this.fitsBeneath(node.coords, coords)) {
            const inside = this.createNode(node.extra, ...coords);
            const result: [ChunkNode<E>, ...ChunkNode<E>[]] = [inside];

            if (node.coords[0] !== coords[0]) {
                const left = this.createNode(node.extra, node.coords[0], coords[0]);
                result.push(left);
            }

            if (node.coords[1] !== coords[1]) {
                const right = this.createNode(node.extra, coords[1], node.coords[1]);
                result.push(right);
            }

            return result;
        }

        // node is shifted right
        if (coords[0] < node.coords[0]) {
            const inside = this.createNode(node.extra, node.coords[0], coords[1]);
            const other = this.createNode(node.extra, coords[1], node.coords[1]);
            return [inside, other];
        }

        // node is shifted left
        const inside = this.createNode(node.extra, coords[0], node.coords[1]);
        const other = this.createNode(node.extra, node.coords[0], coords[0]);
        return [inside, other];
    }

    private addChild(to: ChunkNode<E>, child: ChunkNode<E>): void {
        to.children.push(child);
        to.children.sort((a, b) => a.coords[0] - b.coords[0]);
    }

    private fitsBeneath(parent: Coords, child: Coords): boolean {
        return parent[0] <= child[0] && parent[1] >= child[1];
    }

    private isIntersecting(a: Coords, b: Coords): boolean {
        return a[1] > b[0] && a[0] < b[1];
    }

    private insertNode(node: ChunkNode<E>): void {
        // this.root implied that it will always intersect
        let curr: ChunkNode<E> = this.root;

        for (let i = 0; i < curr.children.length; i++) {
            const parent = curr.children[i]!;

            if (this.isIntersecting(parent.coords, node.coords)) {
                curr = parent;
                // next continue will make this 0
                i = -1;
                continue;
            }
        }

        if (curr === this.root || this.fitsBeneath(curr.coords, node.coords)) {
            this.addChild(curr, node);
            return;
        }

        const [inside, ...others] = this.splitNode(node, curr.coords);

        console.log({ inside, others }, { node, curr });

        this.addChild(curr, inside);
        others.forEach(other => this.insertNode(other));
    }

    addNode(extra: E, from: number, to: number): void {
        this.insertNode(this.createNode(extra, from, to));
    }

    constructor(extra: E, from: number, to: number) {
        this.root = this.createNode(extra, from, to);
    }
}
