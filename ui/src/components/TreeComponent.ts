export namespace TreeComponent {
     export interface DataNode {
        title: string;
        key: string;
        isLeaf?: boolean;
        children?: DataNode[];
    }
}

