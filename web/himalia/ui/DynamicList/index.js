import React, { useState, useEffect } from 'react'

import { prependUrlWithHttps } from '@utils/forms'

import './style.css'

const convertListTopMap = (list) => {
    const map = new Map()
    list.forEach((value, index) => {
        map.set(index, value)
    })
    return map
}

const dynamicList = ({ defaultList, label, labelPlural, removeTitle, id, tabIndex, onChange = null }) => {
    const [inputValue, setInputValue] = useState('')
    const [counter, setCounter] = useState(defaultList.length + 1)
    const [localList, setLocalList] = useState(convertListTopMap(defaultList))

    useEffect(() => {
        setLocalList(convertListTopMap(defaultList))
        setCounter(defaultList.length + 1)
    }, [defaultList])

    const addEntry = (inputValue) => {
        if (inputValue.length === 0) {
            return
        }

        localList.set(counter + 1, inputValue)
        setCounter(counter + 1)
        setLocalList(new Map(localList))
        setInputValue('')
        if (onChange) {
            onChange(Array.from(localList.values()))
        }
    }

    const removeEntry = (key) => {
        return (e) => {
            e.preventDefault()
            localList.delete(key)
            setLocalList(new Map(localList))
            if (onChange) {
                onChange(Array.from(localList.values()))
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
            e.preventDefault()
            e.stopPropagation()
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
                        tabIndex={tabIndex}
                        value={inputValue}
                        onKeyDown={handleKeypress}
                        onChange={(e) => setInputValue(e.target.value)}
                        onBlurCapture={(e) => prependUrlWithHttps(e)}
                        id={`dynamic-list__${id}`}
                        inputMode="url"
                        type="text" />
                </div>
                <div className="dynamic-list__shared-wrapper__action">
                    <button
                        type="button"
                        onClick={() => addEntry(inputValue)}
                        className="button-anchor">Add entry</button>
                </div>
            </div>
        </div>
    )
}

export default dynamicList
