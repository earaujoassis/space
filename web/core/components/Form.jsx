import React from 'react'

import Row from './Row.jsx'
import Columns from './Columns.jsx'

export class Entry extends React.Component {
    static get propTypes() {
        return {
            field: React.PropTypes.string.isRequired,
            value: React.PropTypes.string.isRequired,
            editable: React.PropTypes.bool
        }
    }

    static get defaultProps() {
        return {
            editable: true
        }
    }

    render() {
        return (
            <Row className="profile-entry">
                <Columns className="columns small-5 small-offset-1 field">{this.props.field}</Columns>
                <Columns className="small-6 value">{this.props.value}</Columns>
            </Row>
        )
    }
}
