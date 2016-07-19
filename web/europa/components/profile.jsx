import React from 'react'

import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'
import { Entry } from '../../core/components/form.jsx'

export default class Profile extends React.Component {
    render() {
        return (
            <Row>
                <Columns className="large-offset-1 large-10 end">
                    <Row className="profile">
                        <Columns className="large-12">
                            <Entry field="Name" value="John Doe" />
                            <Entry field="Email" value="email@example.com" />
                            <Entry field="Phone" value="+1 123 4567-890" />
                        </Columns>
                    </Row>
                </Columns>
            </Row>
        )
    }
}
