import React, { useState, useEffect } from 'react'

import SessionsActions from '../actions/sessions'
import Row from '../../core/components/Row.jsx'
import Columns from '../../core/components/Columns.jsx'

import { getParameterByName } from '../../core/utils/url'

const magicLink = () => {
    const [lockedForm, setLockedForm] = useState(false)
    const [holder, setHolderValue] = useState('')
    const [magicLinkRequested, setMagicLinkRequested] = useState(false)

    useEffect(() => {
        let securityTimeoutID = setTimeout(() => {
            setLockedForm(true)
            clearTimeout(securityTimeoutID)
        }, 60000)
        return function cleanup() {
            clearTimeout(securityTimeoutID)
        }
    })

    return (
        <div className="signin-content">
            <Row>
                <Columns className="small-12">
                    <div className="user-avatar magic"></div>
                    {magicLinkRequested === true ? (
                        <div className="requested">
                            <p>If the account holder is valid and active, you should receive an e-mail message in the next few minutes.</p>
                        </div>
                    ) : null}
                    <form action="." method="post">
                        <input type="text"
                            name="holder"
                            placeholder="Access holder"
                            value={holder}
                            onChange={(e) => setHolderValue(e.target.value)}
                            required
                            disabled={lockedForm} />
                        <button type="submit"
                            className="button expand"
                            onClick={(e) => {
                                if (e) e.preventDefault()
                                setMagicLinkRequested(true)
                                setLockedForm(true)
                                let next = getParameterByName('_')
                                let formData = new FormData()
                                formData.append('holder', holder)
                                if (next && next) {
                                    formData.append('next', next)
                                }
                                console.log(next)
                                SessionsActions.requestMagicLink(formData)
                            }}
                            disabled={lockedForm}>Request Magic Link</button>
                    </form>
                    <p className="upper-box">1<sub>min</sub> to request a magic link</p>
                </Columns>
            </Row>
        </div>
    )
}

export default magicLink
