import React, { useEffect, useState } from 'react'

import Row from '@core/components/Row.jsx'
import Columns from '@core/components/Columns.jsx'

const SCOPE_WEIGHTS = {
    'public': 0,
    'openid': 1,
    'read': 2,
    'profile': 2,
    'write': 3
}

const sortScopesByWeight = (scopeString) => {
    if (!scopeString || scopeString.trim() === '') {
        return []
    }

    return scopeString
        .trim()
        .split(/\s+/)
        .filter(scope => scope in SCOPE_WEIGHTS)
        .sort((a, b) => SCOPE_WEIGHTS[a] - SCOPE_WEIGHTS[b])
}

const getHighestWeightScope = (scopeString) => {
    const sortedScopes = sortScopesByWeight(scopeString)
    return sortedScopes.length > 0 ? sortedScopes[sortedScopes.length - 1] : null
}

const Authorize = () => {
    const [serverData, setServerData] = useState(null)
    let keyNumber = 0

    const requestedData = (scope) => {
        switch(scope) {
        case 'openid':
            return [
                'Authentication data for that given application'
            ]
        case 'public':
            return [
                'Authentication data for that given application'
            ]
        case 'profile':
            return [
                'Authentication data for that given application',
                'Your profile data (including first and last names)'
            ]
        case 'read':
            return [
                'Authentication data for that given application',
                'Your profile data (including first and last names)'
            ]
        case 'write':
            return [
                'Authentication data for that given application',
                'Your profile data (including first and last names)',
                'Update your profile data (including first and last names)'
            ]
        }
    }

    useEffect(() => {
        if (document.getElementById('data')) {
            const data = JSON.parse(document.getElementById('data').innerHTML)
            setServerData(data)
        }
    }, [])

    if (serverData === null) {
        return <></>
    }

    return (
        <div className="authorize">
            <Row className="separator">
                <Columns className="small-offset-2 small-4 user">
                    <h2>{`${serverData.first_name} ${serverData.last_name}`}</h2>
                </Columns>
                <Columns className="small-4 end client">
                    <h2>{serverData.client_name}</h2>
                </Columns>
            </Row>
            <Row className="separator">
                <Columns className="small-offset-2 small-8 end">
                    <p>The application ({serverData.client_name}) is requesting access to the following information:</p>
                    <ul className="">
                        {
                            requestedData(getHighestWeightScope(serverData.requested_scope)).map((message) => {
                                return (<li key={keyNumber++}>{message}</li>)
                            })
                        }
                    </ul>
                    <p>{serverData.client_name} (the application) and Space (the current website) will use this information
                    in accordance with their respective terms of service and privacy policies.</p>
                </Columns>
            </Row>
            <Row className="separator">
                <Columns className="small-12 text-center">
                    <Row>
                        <Columns className="small-offset-2 small-4">
                            <form action="" method="post">
                                <input type="hidden" name="access_denied" value="true" />
                                <button className="button expand secondary" type="submit">Cancel</button>
                            </form>
                        </Columns>
                        <Columns className="small-4 end">
                            <form action="" method="post">
                                <button className="button expand" type="submit">Accept</button>
                            </form>
                        </Columns>
                    </Row>
                </Columns>
            </Row>
        </div>
    )
}

export default Authorize
