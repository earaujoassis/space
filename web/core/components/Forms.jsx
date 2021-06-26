import React from 'react';
import PropTypes from 'prop-types';

import Row from './Row.jsx';
import Columns from './Columns.jsx';

// eslint-disable-next-line no-unused-vars
const Entry = ({ field, value, editable }) => {
  return (
    <Row className="profile-entry">
      <Columns className="columns small-5 small-offset-1 field">{field}</Columns>
      <Columns className="small-6 value">{value}</Columns>
    </Row>
  );
};

Entry.propTypes = {
  field: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  editable: PropTypes.bool
};

Entry.defaultProps = {
  editable: false
};

export {
  Entry
};
