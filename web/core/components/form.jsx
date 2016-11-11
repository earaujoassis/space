import React from 'react'

import Row from './row.jsx'
import Columns from './columns.jsx'

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
                <Columns className="columns medium-5 medium-offset-1 field">{this.props.field}</Columns>
                <Columns className="medium-6 value">{this.props.value}</Columns>
            </Row>
        )
    }
}
