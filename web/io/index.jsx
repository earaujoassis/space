import React from 'react'
import ReactDOM from 'react-dom'

import FeaturesStore from '../core/stores/features'

import Root from './components/Root.jsx'

FeaturesStore.loadData()

ReactDOM.render(<Root />, document.getElementById('application-context'))
