import React from 'react'
import { icon } from '@fortawesome/fontawesome-svg-core/import.macro'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

const radioGroup = ({ label, leftOption, rightOption }) => {
    return (
        <div className="globals__siblings globals__no-gap">
            <div className="globals__input-wrapper globals__center">
                <span className="globals__radio-label">{label}</span>
            </div>
            <div className="globals__input-wrapper globals__center">
                <span className="globals__radio-wrapper">
                    <span className="globals__radio-icon-wrapper">
                        <FontAwesomeIcon
                            className="globals__radio-icon"
                            icon={icon({name: 'circle', style: 'regular'})} />
                    </span>
                    <span className="globals__radio-copy">{leftOption}</span>
                </span>
            </div>
            <div className="globals__input-wrapper">
                <span className="globals__radio-wrapper selected">
                    <span className="globals__radio-icon-wrapper">
                        <FontAwesomeIcon
                            className="globals__radio-icon"
                            icon={icon({name: 'circle-check', style: 'regular'})} />
                    </span>
                    <span className="globals__radio-copy">{rightOption}</span>
                </span>
            </div>
        </div>
    )
}

export default radioGroup
