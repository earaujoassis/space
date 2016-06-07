import React from 'react';

import Row from '../reusable/row.jsx';
import Columns from '../reusable/columns.jsx';

export default class Applications extends React.Component {
    render() {
        return (
            <Row>
                <Columns className="large-offset-1 large-10 end">
                    <Row className="applications">
                        <Columns className="medium-6">
                            <div className="application-card">
                                <div className="title">Application (application.example.com)</div>
                            </div>
                        </Columns>
                        <Columns className="medium-6"></Columns>
                    </Row>
                </Columns>
            </Row>
        );
    }
};
