import React from 'react';
import PropTypes from 'prop-types';

const columns = ({ className, children }) => {
  const finalClassName = `columns ${className}`;

  return (
    <div className={finalClassName}>{children}</div>
  );
};

columns.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node
};

columns.defaultProps = {
  className: ''
};

export default columns;
