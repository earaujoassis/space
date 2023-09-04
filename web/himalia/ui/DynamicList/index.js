import React, { useState } from 'react'

import './style.css'

const dynamicList = ({ label, labelPlural, removeTitle, id, onChange = null }) => {
    const [inputValue, setInputValue] = useState('')
    const [counter, setCounter] = useState(0)
    const [localList, setLocalList] = useState(new Map())

    const addEntry = (inputValue) => {
        if (inputValue.length === 0) {
            return
        }

        localList.set(counter + 1, inputValue)
        setCounter(counter + 1)
        setLocalList(new Map(localList))
        setInputValue('')
        if (onChange) {
            onChange(localList)
        }
    }

    const removeEntry = (key) => {
        return (e) => {
            e.preventDefault()
            localList.delete(key)
            setLocalList(new Map(localList))
            if (onChange) {
                onChange(localList)
            }
        }
    }

    const listElements = () => {
        if (localList.size === 0) {
            return null
        }

        return Array.from(localList.entries()).map(([key, value]) => {
            return (
                <li key={`dynamic-list-${id}-entry-${key}`}>
                    <div className="dynamic-list__list-entry">
                        <button
                            onClick={removeEntry(key)}
                            className="dynamic-list__list-remove-entry"
                            title={removeTitle}>&times;</button>
                        <span>{value}</span>
                    </div>
                </li>
            )
        })
    }

    const handleKeypress = (e) => {
        if (e.key === 'Enter') {
            addEntry(inputValue)
        }
    }

    return (
        <div className="globals__siblings">
            <div className="globals__input-wrapper">
                <h4 className="dynamic-list__list-header">{labelPlural}</h4>
                <ul className="dynamic-list__list-urls">
                    {listElements()}
                </ul>
            </div>
            <div className="globals__input-wrapper dynamic-list__shared-wrapper">
                <div className="dynamic-list__shared-wrapper__field">
                    <label htmlFor={`dynamic-list__${id}`}>{label}</label>
                    <input
                        value={inputValue}
                        onKeyDown={handleKeypress}
                        onChange={(e) => setInputValue(e.target.value)}
                        id={`dynamic-list__${id}`}
                        inputMode="url"
                        type="text" />
                </div>
                <div className="dynamic-list__shared-wrapper__action">
                    <button
                        onClick={() => addEntry(inputValue)}
                        className="button-anchor">Add entry</button>
                </div>
            </div>
        </div>
    )
}

export default dynamicList
