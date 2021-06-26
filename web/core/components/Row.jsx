import React from 'react';
import PropTypes from 'prop-types';

const row = ({ className, children }) => {
  const finalClassName = `row ${className}`;

  return (
    <div className={finalClassName}>{children}</div>
  );
};

row.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
};

row.defaultProps = {
  className: ''
};

export default row;
