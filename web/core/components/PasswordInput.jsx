import React, { useState } from 'react';

const hideImg = '/public/imgs/eye-open.png';
const displayImg = '/public/imgs/eye-blocked.png';

const passwordInput = ({ name = 'password', placeholder = 'Password' }) => {
  const [showPassword, setShowPassword] = useState(true);

  return (
    <div className="password-visibility">
      <input type={showPassword ? 'password' : 'text'} name={name} placeholder={placeholder} required />
      <button
        className="visibility-toggle"
        onClick={(e) => {
          e.preventDefault();
          setShowPassword(!showPassword);
        }}
      >
        <img
          src={showPassword ? displayImg : hideImg}
          width="20"
          title="Toggle password visibility"
          alt=""
        />
      </button>
    </div>
  );
};

export default passwordInput;
