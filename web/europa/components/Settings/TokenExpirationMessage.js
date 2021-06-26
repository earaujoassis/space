import React from 'react';

export const TokenExpirationMessage = ({ tokenExpired }) => {
  if (tokenExpired) {
    return (
      <div className="token-error">Your action token is expired. Please refresh your page.</div>
    );
  }

  return null;
};
