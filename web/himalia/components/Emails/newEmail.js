import React from 'react'

import { extractDataForm } from '@utils/forms'

const newEmail = ({ addEmail }) => {
  return (
    <div className="emails__form">
      <form
        className="form-common"
        action="."
        method="post"
        data-form-type="other"
        onSubmit={e => {
          e.preventDefault()
          const data = extractDataForm(e.target, ['address'])
          addEmail(data)
        }}
      >
        <div className="globals__siblings emails__form-fields">
          <div className="globals__children">
            <div className="globals__input-wrapper">
              <label htmlFor="new-email__address">Add email address</label>
              <input
                tabIndex="1"
                required
                autoComplete="off"
                id="new-email__address"
                name="address"
                inputMode="email"
                type="email"
                placeholder="Email address"
              />
            </div>
          </div>
          <div className="globals__children">
            <div className="globals__input-wrapper">
              <input
                tabIndex="2"
                type="submit"
                className="button-anchor emails__button"
                value="Add"
              />
            </div>
          </div>
        </div>
      </form>
    </div>
  )
}

export default newEmail
