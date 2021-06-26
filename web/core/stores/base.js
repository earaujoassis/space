import EventEmitter from 'events';

const CHANGE_EVENT = 'change';

export default class Store extends EventEmitter {
  constructor () {
    super();
  }

  emitChange () {
    this.emit(CHANGE_EVENT);
  }

  addChangeListener (callback) {
    this.on(CHANGE_EVENT, callback);
  }

  removeChangeListener (callback) {
    this.removeListener(CHANGE_EVENT, callback);
  }

  isCurrentActionType (type) {
    return this.getState().type === type;
  }

  isCurrentAction (type, uuid) {
    return this.getState().type === type && uuid === this.getState().action_uuid;
  }

  isRequestedAction (requestType) {
    return this.getState().request_type === requestType;
  }

  success () {
    return this.getState().type === 'SUCCESS';
  }
}
