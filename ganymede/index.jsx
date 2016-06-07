import React from 'react';
import ReactDOM from 'react-dom';

import Root from './components/root.jsx';
import ApplicationActionCreator from './actions/application';

new ApplicationActionCreator.setupApp();

ReactDOM.render(
    <Root />,
    document.getElementById('application-context')
);
