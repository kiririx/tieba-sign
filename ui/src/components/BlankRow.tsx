import React from "react";

class _Props {
    height = 10;
}

export default class BlankRow extends React.Component<_Props, any> {

    static defaultProps = new _Props();

    constructor(props: _Props) {
        super(props);
    }

    render() {
        return (
            <div style={{padding: `${this.props.height}px 0`}}/>
        );
    }
}