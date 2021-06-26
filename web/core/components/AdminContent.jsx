import React from 'react';
import PropTypes from 'prop-types';

const adminContent = ({ user = {}, children }) => {
  if (user.is_admin) {
    return (
      <div>
        {children}
      </div>
    );
  }

  return null;
};

adminContent.propTypes = {
  user: PropTypes.object,
  children: PropTypes.node
};

adminContent.defaultProps = {
  user: {}
};

export default adminContent;
