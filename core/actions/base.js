import * as node_uuid from 'node-uuid';

import dispatcher from '../dispatcher';

export default class ActionCreator {
    constructor() {
        this._actionUUID = null;
        this.customPayload = null;
    }

    setUUID(uuid) {
        this._actionUUID = uuid || node_uuid.v1();
        return this;
    }

    dispatch(payload) {
        if (!payload.actionUUID && this._actionUUID)
            payload.actionUUID = this._actionUUID;
        if (!this.customPayload)
            this.customPayload = {};
        let customPayload = Object.assign({}, this.customPayload)
        payload.payload = customPayload
        dispatcher.dispatch(payload);
        this.customPayload = null;
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
