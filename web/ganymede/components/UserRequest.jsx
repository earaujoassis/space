import React, { useState, useEffect } from 'react';

import SessionsActions from '../actions/sessions';
import UsersActions from '../actions/users';
import Row from '../../core/components/Row.jsx';
import Columns from '../../core/components/Columns.jsx';
import SuccessBox from '../../core/components/SuccessBox.jsx';

import { getParameterByName } from '../../core/utils/url';

// TODO Handle errors from Stores
const RequestForm = ({ type, onRequest }) => {
  const [lockedForm, setLockedForm] = useState(false);
  const [holder, setHolderValue] = useState('');

  useEffect(() => {
    const securityTimeoutID = setTimeout(() => {
      setLockedForm(true);
      clearTimeout(securityTimeoutID);
    }, 60000);
    return function cleanup () {
      clearTimeout(securityTimeoutID);
    };
  });

  return (
        <div className="middle-box signin-content">
            <Row>
                <Columns className="small-12">
                    <div className={`user-avatar ${type}`}></div>
                    <form action="." method="post">
                        <input type="text"
                            name="holder"
                            placeholder="Access holder"
                            value={holder}
                            onChange={(e) => setHolderValue(e.target.value)}
                            required
                            disabled={lockedForm} />
                        <button type="submit"
                            className="button expand"
                            onClick={(e) => {
                              if (e) e.preventDefault();

                              const formData = new FormData();
                              formData.append('holder', holder);
                              if (type === 'magic') {
                                const next = getParameterByName('_');
                                if (next && next) {
                                  formData.append('next', next);
                                }
                                SessionsActions.requestMagicLink(formData);
                              } else {
                                formData.append('request_type', 'password');
                                UsersActions.requestUpdate(formData);
                              }

                              setLockedForm(true);
                              onRequest(true);
                            }}
                            disabled={lockedForm}>
                            {type === 'magic' ? 'Request Magic Link' : 'Request to update password'}
                        </button>
                    </form>
                    <p className="upper-box">1<sub>min</sub> to make your request</p>
                </Columns>
            </Row>
        </div>
  );
};

const userRequest = ({ type }) => {
  const [requestedFulfilled, setRequestFulfilment] = useState(false);

  return (
        <div>
            {requestedFulfilled === true
              ? (
                <SuccessBox>
                    <p>If the account holder is valid and active, you should receive an e-mail message in the next few minutes.</p>
                </SuccessBox>
                )
              : (
                <RequestForm
                    type={type}
                    onRequest={setRequestFulfilment} />
                )}

        </div>
  );
};

export default userRequest;
