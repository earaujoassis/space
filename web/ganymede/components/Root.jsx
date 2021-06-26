import React, { useState } from 'react';

import SignIn from './SignIn.jsx';
import UserRequest from './UserRequest.jsx';

const PASSWORD_STRATEGY = 'password';
const MAGIC_LINK_STRATEGY = 'magic-link';
const REQUEST_PASSWORD = 'request-password';

const strategyComponent = (current) => {
  if (current === PASSWORD_STRATEGY) {
    return <SignIn />;
  } else if (current === MAGIC_LINK_STRATEGY) {
    return <UserRequest type="magic" />;
  } else {
    return <UserRequest type="password" />;
  }
};

const strategyCallToAction = (current, setStrategy) => {
  if (current === PASSWORD_STRATEGY) {
    return <button
            className="button-anchor"
            onClick={() => setStrategy(MAGIC_LINK_STRATEGY)}>Request magic link</button>;
  } else if (current === MAGIC_LINK_STRATEGY) {
    return <button
            className="button-anchor"
            onClick={() => setStrategy(PASSWORD_STRATEGY)}>Use password and TOTP</button>;
  } else {
    return null;
  }
};

const passwordRequestButtonStates = (current, setStrategy) => {
  if (current === REQUEST_PASSWORD) {
    return <button
            className="button-anchor"
            onClick={() => setStrategy(PASSWORD_STRATEGY)}>I remember now!</button>;
  } else {
    return <button
            className="button-anchor"
            onClick={() => setStrategy(REQUEST_PASSWORD)}>I forgot my password</button>;
  }
};

const root = () => {
  const [currentStrategy, setSignInStrategy] = useState(PASSWORD_STRATEGY);

  return (
        <div className="ganymede-root">
            {strategyComponent(currentStrategy)}
            <div className="ganymede-strategies">
                <div className="strategy-sibling">
                    {strategyCallToAction(currentStrategy, setSignInStrategy)}
                </div>
                <div className="strategy-sibling">
                    {passwordRequestButtonStates(currentStrategy, setSignInStrategy)}
                </div>
            </div>
        </div>
  );
};

export default root;
