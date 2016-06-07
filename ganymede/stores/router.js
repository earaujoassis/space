import page from 'page';

import { Routes, ActionTypes } from '../constants';
import dispatcher from '../dispatcher';
import Store from './base';

const getInitialState = function() {
    return {
        routeData: null,
        context: null
    }
};

const onPageNavigation = function(pattern, context) {
    _state.routeData = Routes.find(route => route.path == pattern);
    _state.context = Object.assign({}, context, { pattern: pattern });
    RouterStore.emitChange();
};

const handleNotFound = function(context) {
    console.debug(`Unhandled path: ${ context.path }`);
};

var _state = getInitialState();

class RouterStoreBase extends Store {
    constructor() {
        super();
        this.dispatchToken = dispatcher.register(function(action) {
            switch (action.type) {
                case ActionTypes.SETUP_APP:
                    RouterStore.setup();
                    break;
            }
        });
    }

    setup() {
        _state = getInitialState();
        var paths = Routes.map(value => value['path']);
        paths.forEach(function(path) {
            page(path, onPageNavigation.bind({}, path));
        });
        page('*', handleNotFound);
        page();
    }

    getState() {
        return _state;
    }

    getCurrentContext() {
        return _state.context;
    }

    getRouteData() {
        return _state.routeData;
    }

    goToRoot() {
        page('/');
    }
};

const RouterStore = new RouterStoreBase();

export default RouterStore;
