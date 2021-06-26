import React from 'react';

import Row from '@core/components/Row.jsx';
import Columns from '@core/components/Columns.jsx';
import { Entry } from '@core/components/Forms.jsx';

import { extractFormData } from '@core/utils/forms';

export const Adminify = ({ adminify, user = {}, features = [] }) => {
  if (!user.is_admin && features.includes('user.adminify')) {
    return (
      <Row className="profile-entry">
        <Columns className="columns small-11 small-offset-1 field">
          <form
            onSubmit={(e) => {
              e.preventDefault();
              const data = new FormData();
              data.append('application_key', extractFormData(e.target, ['application_key'].application_key));
              data.append('user_id', user.id);
              adminify(data);
            }}>
            <input className="thin-input"
              type="text"
              name="application_key"
              placeholder="Application Key"
              required />
            <button
              className="button-anchor"
            >
              Make me an admin
            </button>
          </form>
        </Columns>
      </Row>
    );
  } else if (user.is_admin) {
    return (
      <Entry field="Role" value="Administrator" />
    );
  }

  return null;
};
