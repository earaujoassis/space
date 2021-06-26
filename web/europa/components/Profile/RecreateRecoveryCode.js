import React, { useState } from 'react';

export const RecreateRecoveryCode = ({ requestUserUpdate = () => {}, user = {} }) => {
  const [secretCodesRequested, setSecretCodesRequested] = useState(false);

  if (secretCodesRequested) {
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
          data.append('request_type', 'secrets');
          data.append('holder', user.username);
          requestUserUpdate(data);
          setSecretCodesRequested(true);
        }}
        className="button-anchor"
      >
        Recreate recovery code and secret code generator
      </button>
    </p>
  );
};
