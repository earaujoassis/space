import React from 'react';

import Row from './Row.jsx';
import Columns from './Columns.jsx';

const successBox = ({ children }) => {
  return (
    <div className="middle-box success-box">
      <Row className="success-box-row">
        <Columns className="small-12">
          <div className="success-box-content">{children}</div>
        </Columns>
      </Row>
    </div>
  );
};

export default successBox;
