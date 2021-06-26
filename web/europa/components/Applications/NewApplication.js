import React from 'react';

import Row from '@core/components/Row.jsx';
import Columns from '@core/components/Columns.jsx';

import { extractFormData } from '@core/utils/forms';

const newApplication = ({
  createClient,
  postSubmit = () => {}
}) => {
  return (
    <Row className="new-application">
      <Columns className="small-6">
        <h2 className="title">Create a new client application</h2>
        <p className="description">
          By clicking &quot;Create Application&quot;, you agree to our <a href="//quatrolabs.com/terms-of-service">terms
          of service</a> and <a href="//quatrolabs.com/privacy-policy">privacy policy</a>. Also, you guarantee that the corresponding
          client application will adhere to those terms and policites, while handling user data.
        </p>
      </Columns>
      <Columns className="small-6">
        <form
          className="form-common"
          action="."
          method="post"
          onSubmit={(e) => {
            e.persist();
            e.preventDefault();
            const attrs = ['name', 'description', 'canonical_uri', 'redirect_uri'];
            createClient(extractFormData(e.target, attrs), () => {
              e.target.reset();
              postSubmit();
            });
          }}
        >
          <input type="hidden" name="action_token" value="" />
          <input type="text" name="name" placeholder="Name" required />
          <input type="text" name="description" placeholder="Description" required />
          <input type="url" name="canonical_uri" placeholder="Canonical URI" pattern="https?://.*" required />
          <input type="url" name="redirect_uri" placeholder="Redirect URI" pattern="https?://.*" required />
          <button type="submit" className="button expand">Create Application</button>
        </form>
      </Columns>
    </Row>
  );
};

export default newApplication;
