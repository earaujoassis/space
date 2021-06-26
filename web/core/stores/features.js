import Store from './base';

let _setupData = {};

class FeaturesStoreBase extends Store {
  constructor () {
    super();
  }

  isFeatureActive (key) {
    return _setupData['feature.gates'] && _setupData['feature.gates'][key];
  }

  loadData () {
    if (document.getElementById('data')) {
      _setupData = JSON.parse(document.getElementById('data').innerHTML);
    }
  }
}

const FeaturesStore = new FeaturesStoreBase();

export default FeaturesStore;
