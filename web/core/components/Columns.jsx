import React from 'react';
import PropTypes from 'prop-types'

export default class Columns extends React.Component {
    static get propTypes() {
        return {
            className: PropTypes.string
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
