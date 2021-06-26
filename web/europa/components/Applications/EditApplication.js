import React from 'react';

import Row from '@core/components/Row.jsx';
import Columns from '@core/components/Columns.jsx';
import { extractFormData } from '@core/utils/forms';

const editApplication = ({ updateClient, cancelUpdate, client, selectedClientId }) => {
  if (client.id === selectedClientId) {
    return (
      <form
        className="form-common internal"
        onSubmit={(e) => {
          e.preventDefault();
          const attrs = ['canonical_uri', 'redirect_uri', 'scopes'];
          const data = extractFormData(e.target, attrs);
          updateClient(client.id, data, () => cancelUpdate());
        }}
      >
        <Row className="new-application">
          <Columns className="small-5">
            <label htmlFor="canonical_uri">Canonical URI</label>
            <textarea
              type="url"
              id="canonical_uri"
              name="canonical_uri"
              placeholder="Canonical URI"
              pattern="((https?://.*)(\n)?)+"
              defaultValue={client.uri}
              required
            >
            </textarea>
          </Columns>
          <Columns className="small-5">
            <label htmlFor="redirect_uri">Redirect URI</label>
            <textarea
              type="url"
              id="redirect_uri"
              name="redirect_uri"
              placeholder="Redirect URI"
              pattern="((https?://.*)(\n)?)+"
              defaultValue={client.redirect}
              required
            >
            </textarea>
          </Columns>
          <Columns className="small-2 end">
            <div className="button-wrapper">
              <button className="button-anchor" type="submit">Save</button>
            </div>
            <div>
              <button
                onClick={(e) => {
                  e.preventDefault();
                  cancelUpdate();
                }}
                className="button-anchor"
                type="button"
              >
                Cancel
              </button>
            </div>
          </Columns>
        </Row>
        <Row>
          <Columns className="small-12">
            <p className="form-warning">
              Use one line for each redirect URI. Each redirect URI must be included as a canonical URI.
            </p>
          </Columns>
        </Row>
        <Row>
          <Columns className="small-12">
            <label htmlFor="scopes">Scope</label>
            <span className="select-wrapper">
              <select defaultValue={''} name="scopes">
                <option value="" disabled>Select new scope</option>
                <option value="public">Public</option>
                <option value="read">Read</option>
              </select>
            </span>
            <p className="form-warning">By default, all applications are created with &quot;Public&quot; scope</p>
          </Columns>
        </Row>
      </form>
    );
  }

  return null;
};

export default editApplication;
