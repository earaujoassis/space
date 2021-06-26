import React, { useEffect, useState } from 'react';

import Row from '../../core/components/Row.jsx';
import Columns from '../../core/components/Columns.jsx';
import PasswordInput from '../../core/components/PasswordInput.jsx';
import SuccessBox from '../../core/components/SuccessBox.jsx';

import { extractFormData } from '../../core/utils/forms';

import UserStore from '../stores/users';
import UsersActions from '../actions/users';

const attemptPassworUpdate = (target) => {
  const data = extractFormData(target, ['new_password', 'password_confirmation']);
  data.append('_', UserStore.getActionToken());
  UsersActions.updatePassword(data);
};

const UpdatePassword = ({ onSuccess }) => {
  useEffect(() => {
    const updateLocalStoreState = () => {
      if (UserStore.success()) {
        onSuccess();
      }
    };

    UserStore.loadData();
    UserStore.addChangeListener(updateLocalStoreState);

    return function cleanup () {
      UserStore.removeChangeListener(updateLocalStoreState);
    };
  }, []);

  return (
        <div className="middle-box plain resource-owner-password">
            <Row>
                <Columns className="small-12">
                    <form
                        className="form-common"
                        action="."
                        method="patch"
                        onSubmit={(e) => {
                          e.preventDefault();
                          attemptPassworUpdate(e.target);
                        }}>
                        <p>Update your password with the required fields below</p>
                        <PasswordInput placeholder="New password" name="new_password" />
                        <PasswordInput placeholder="Confirm password" name="password_confirmation" />
                        <button type="submit"
                            className="button expand"
                            disabled={false}>Update password</button>
                    </form>
                </Columns>
            </Row>
        </div>
  );
};

const ownerUpdate = () => {
  const [hasUpdated, setHasUpdated] = useState(false);

  return (
        <div className="resource-owner-update">
            <Row>
                <Columns className="small-12">
                    {hasUpdated === true
                      ? (
                        <SuccessBox>
                            <p>Password updated sucessfully!</p>
                            <p>Get <a href="/">back to the application</a>.</p>
                        </SuccessBox>
                        )
                      : (
                        <UpdatePassword onSuccess={() => setHasUpdated(true)} />
                        )}
                </Columns>
            </Row>
        </div>
  );
};

export default ownerUpdate;
