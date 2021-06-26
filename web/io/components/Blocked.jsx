import React from 'react';

import Row from '../../core/components/Row.jsx';
import Columns from '../../core/components/Columns.jsx';

const blocked = () => {
  return (
        <div className="centered-message">
            <Row>
                <Columns className="small-6 small-offset-3 end columns">
                    <h2>Hey there!</h2>
                    <p>
                        Unfortunately, we are not accepting any new user sign up at this time.
                        Please make sure to check that back soon.
                    </p>
                    <p>Kindly,</p>
                    <p>The quatroLABS Team.</p>
                </Columns>
            </Row>
        </div>
  );
};

export default blocked;
