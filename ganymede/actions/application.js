import { ActionTypes } from '../constants';
import ActionCreator from '../../core/actions/base';

class ApplicationActionFactory {
    setupApp() {
        let action = new ActionCreator();
        action.setUUID();
        action.dispatch({type: ActionTypes.SETUP_APP});
        return action.actionID();
    }
};

const ApplicationActionCreator = new ApplicationActionFactory();

export default ApplicationActionCreator;
