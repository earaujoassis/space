import * as node_uuid from 'node-uuid';

import dispatcher from '../dispatcher';

export function errorHandler(error) {
    let action = new ActionCreator()
    action.setUUID()
    action.dispatch({type: 'ERROR', payload: error})
    return action
}

export function processResponse(response) {
    if (response.status >= 200 && response.status < 300) {
        return Promise.resolve(response)
    } else {
        var error = new Error(response.statusText);
        error.response = response;
        return Promise.reject(error);
    }
}

export function processData(response) {
    if (response.status !== 204) {
        return response.json()
    } else {
        return {}
    }
}

export function successHandler(data) {
    let action = new ActionCreator()
    action.setUUID()
    action.dispatch({type: 'SUCCESS', payload: data})
    return action
}

export class ActionCreator {
    constructor() {
        this._actionUUID = null;
        this.customPayload = null;
    }

    setUUID(uuid) {
        this._actionUUID = uuid || node_uuid.v1();
        return this;
    }

    dispatch(doc) {
        if (!doc.actionUUID && this._actionUUID)
            doc.actionUUID = this._actionUUID;
        if (!this.customPayload)
            this.customPayload = {};
        doc.payload = Object.assign({}, this.customPayload, doc.payload)
        dispatcher.dispatch(doc);
        this.customPayload = undefined;
        return this;
    }

    setPayload(payload) {
        this.customPayload = payload;
        return this;
    }

    actionID() {
        return this._actionUUID;
    }
};
