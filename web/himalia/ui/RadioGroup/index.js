import React, { useState, useEffect } from 'react'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

const radioGroup = ({
  label,
  defaultOption,
  leftOption,
  rightOption,
  onChange = null,
}) => {
  const [selectedRadio, setSelectedRadio] = useState(defaultOption)

  useEffect(() => {
    setSelectedRadio(defaultOption)
  }, [defaultOption])

  const setRadio = value => {
    return () => {
      setSelectedRadio(value)
      if (onChange) {
        onChange(value)
      }
    }
  }

  const isSelected = value => {
    return selectedRadio === value ? 'selected' : ''
  }

  const setIcon = value => {
    if (selectedRadio === value) {
      return (
        <FontAwesomeIcon
          className="globals__radio-icon"
          icon={icon({ name: 'circle-check', style: 'regular' })}
        />
      )
    } else {
      return (
        <FontAwesomeIcon
          className="globals__radio-icon"
          icon={icon({ name: 'circle', style: 'regular' })}
        />
      )
    }
  }

  return (
    <div className="globals__siblings globals__no-gap">
      <div className="globals__input-wrapper globals__center">
        <span className="globals__radio-label">{label}</span>
      </div>
      <div className="globals__input-wrapper globals__center">
        <span
          onClick={setRadio(leftOption)}
          className={`globals__radio-wrapper ${isSelected(leftOption)}`}
        >
          <span className="globals__radio-icon-wrapper">
            {setIcon(leftOption)}
          </span>
          <span className="globals__radio-copy">{leftOption}</span>
        </span>
      </div>
      <div className="globals__input-wrapper">
        <span
          onClick={setRadio(rightOption)}
          className={`globals__radio-wrapper ${isSelected(rightOption)}`}
        >
          <span className="globals__radio-icon-wrapper">
            {setIcon(rightOption)}
          </span>
          <span className="globals__radio-copy">{rightOption}</span>
        </span>
      </div>
    </div>
  )
}

export default radioGroup
