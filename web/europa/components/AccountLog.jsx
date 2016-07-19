import React from 'react'

import Row from '../../core/components/row.jsx'
import Columns from '../../core/components/columns.jsx'

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
                                <td>123.456.789.0</td>
                                <td>application.example.com</td>
                                <td>/api/1/oauth/authorize</td>
                                <td>2016-04-16 15:50:40.211005</td>
                            </tr>
                            <tr>
                                <td>123.456.789.0</td>
                                <td>application.example.com</td>
                                <td>/api/1/oauth/authorize</td>
                                <td>2016-04-16 15:52:10.783773</td>
                            </tr>
                            <tr>
                                <td>123.456.789.0</td>
                                <td>application.example.com</td>
                                <td>/api/1/oauth/authorize</td>
                                <td>2016-04-16 15:52:10.783773</td>
                            </tr>
                        </tbody>
                    </table>
                </Columns>
            </Row>
        )
    }
}
