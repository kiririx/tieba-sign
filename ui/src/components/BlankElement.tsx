import React from "react";

interface BlankElementProps {
    width: number
    height: number
}

export default class BlankElement extends React.Component<BlankElementProps, any> {

    static defaultProps = {
        width: 10,
        height: 10
    }

    constructor(props: BlankElementProps) {
        super(props);
    }

    render() {
        return (
            <div style={{display: "inline-block", padding: `${this.props.height}px ${this.props.width}px`}}/>
        )
    }
}