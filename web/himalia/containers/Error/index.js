import React from 'react'

import './style.css'

const error = ({ type, children }) => {
  let imgSrc = '/public/imgs/illustration.server_error.svg'

  if (type == 'bug') {
    imgSrc = '/public/imgs/illustration.bug.svg'
  }

  return (
    <div className="error-layout">
      <div className="error-illustration">
        <img width="300" src={imgSrc} />
      </div>
      <div className="globals__error-message">
        <h2>Oh, crap</h2>
        {children}
      </div>
    </div>
  )
}

export default error
