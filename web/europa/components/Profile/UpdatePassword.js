import React, { useState } from 'react';

export const UpdatePassword = ({ requestUserUpdate = () => {}, user = {} }) => {
  const [updatePasswordRequested, setUpdatePasswordRequested] = useState(false);

  if (updatePasswordRequested) {
    return (
      <p>You should receive an e-mail message in the next few minutes.</p>
    );
  }

  return (
    <p>
      <button
        onClick={(e) => {
          e.preventDefault();
          const data = new FormData();
          data.append('request_type', 'password');
          data.append('holder', user.username);
          requestUserUpdate(data);
          setUpdatePasswordRequested(true);
        }}
        className="button-anchor"
      >
        Update password
      </button>
    </p>
  );
};
