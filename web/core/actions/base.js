import * as node_uuid from 'node-uuid'

import { ActionTypes } from '../constants'
import dispatcher from '../dispatcher'

export function processResponse(response) {
    if (response.status >= 200 && response.status < 300) {
        return Promise.resolve(response)
    } else {
        return Promise.resolve(response)
    }
}

export function processData(response) {
    let data

    if (response.status !== 204) {
        try {
            return response.json()
        } catch(e) { /*...*/ }
    }
    return {}
}

export function processHandler(data) {
    let action

    if (data.error) {
        action = new ActionCreator()
        action.setUUID()
        action.dispatch({type: ActionTypes.ERROR, payload: data})
        return action
    }

    action = new ActionCreator()
    action.setUUID()
    action.dispatch({type: ActionTypes.SUCCESS, payload: data})
    return action
}

export class ActionCreator {
    constructor() {
        this._actionUUID = null
        this.customPayload = null
    }

    setUUID(uuid) {
        this._actionUUID = uuid || node_uuid.v1()
        return this
    }

    dispatch(doc) {
        if (!doc.actionUUID && this._actionUUID)
            doc.actionUUID = this._actionUUID
        if (!this.customPayload)
            this.customPayload = {}
        doc.payload = Object.assign({}, this.customPayload, doc.payload)
        dispatcher.dispatch(doc)
        this.customPayload = undefined
        return this
    }

    setPayload(payload) {
        this.customPayload = payload
        return this
    }

    actionID() {
        return this._actionUUID
    }
}
