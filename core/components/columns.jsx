import React from 'react';

export default class Columns extends React.Component {
    static get propTypes() {
        return {
            className: React.PropTypes.string
        }
    }

    static get defaultProps() {
        return {
            className: ""
        }
    }

    render() {
        let className = `columns ${this.props.className}`;
        return (
            <div className={className}>{this.props.children}</div>
        );
    }
};
