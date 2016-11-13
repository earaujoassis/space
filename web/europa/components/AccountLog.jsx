import React from 'react'

import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'

export default class AccountLog extends React.Component {
    render() {
        return (
            <Row>
                <Columns className="large-12">
                    <table>
                        <thead>
                            <tr>
                                <th scope="col">IP</th>
                                <th scope="col">Domain</th>
                                <th scope="col">URI (path)</th>
                                <th scope="col">Date & Time (UTC)</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>?</td>
                                <td>?</td>
                                <td>?</td>
                                <td>?</td>
                            </tr>
                        </tbody>
                    </table>
                </Columns>
            </Row>
        )
    }
}
