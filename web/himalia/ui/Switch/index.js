import React from 'react'

const switchUi = ({ id, name, label, checked, disabled, onChange }) => {
  return (
    <div className="switch tiny">
      <input
        className="switch-input"
        onChange={onChange}
        disabled={disabled}
        checked={checked}
        id={id}
        type="checkbox"
        name={name}
      />
      <label className="switch-paddle" htmlFor={name}>
        <span className="show-for-sr">{label}</span>
      </label>
    </div>
  )
}

export default switchUi
