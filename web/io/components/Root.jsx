import React from 'react';

import FeaturesStore from '../../core/stores/features';
import UserStore from '../stores/users';

import Blocked from './Blocked.jsx';
import SignUp from './SignUp.jsx';
import Success from './Success.jsx';

export default class Root extends React.Component {
  constructor () {
    super();
    this.state = UserStore.getState().payload || {};
    this._updateFromStore = this._updateFromStore.bind(this);
  }

  componentDidMount () {
    UserStore.addChangeListener(this._updateFromStore);
  }

  componentWillUnmount () {
    UserStore.removeChangeListener(this._updateFromStore);
  }

  render () {
    if (!FeaturesStore.isFeatureActive('user.create')) {
      return (<Blocked />);
    }
    if (!!this.state.recover_secret && !!this.state.code_secret_image) {
      return (<Success codeSecretImage={this.state.code_secret_image}
                recoverSecret={this.state.recover_secret} />);
    } else {
      return (<SignUp validationFailed={this.state.validationFailed} />);
    }
  }

  _updateFromStore () {
    if (UserStore.success()) {
      this.setState(UserStore.getState().payload || {});
    } else {
      const error = UserStore.getState().payload;
      if (error.user) {
        this.setState({ validationFailed: true });
      }
    }
  }
}
