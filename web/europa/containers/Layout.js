import React from 'react';
import { withRouter } from 'react-router-dom';

import Settings from '@europa/components/Settings';

const layout = ({ children, location }) => {
  return (
    <Settings location={location}>
      {children}
    </Settings>
  );
};

export default withRouter(layout);
