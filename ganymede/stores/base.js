import EventEmitter from 'events';

const CHANGE_EVENT = 'change';

export default class Store extends EventEmitter {
    constructor() {
        super();
    }

    emitChange() {
        this.emit(CHANGE_EVENT);
    }

    addChangeListener(callback) {
        this.on(CHANGE_EVENT, callback);
    }

    removeChangeListener(callback) {
        this.removeListener(CHANGE_EVENT, callback);
    }

    isCurrentActionType(type) {
        this.getState().type === type;
    }

    isCurrentAction(type, uuid) {
        this.getState().type === type && uuid === this.getState().action_uuid;
    }

    isRequestedAction(request_type) {
        this.getState().request_type === request_type;
    }

    success() {
        this.getState().type === 'SUCCESS';
    }
};
