export const Routes = [
    {path: '/', view: 'Settings'},
    {path: '/applications', view: 'Settings', subview: 'Applications'},
    {path: '/profile', view: 'Settings', subview: 'Profile'},
    {path: '/account-log', view: 'Settings', subview: 'AccountLog'},
    {path: '/settings', view: 'Settings', subview: 'SpaceSettings'}
];

export const ActionTypes = [
    'SUCCESS', 'ERROR', 'SETUP_APP'
].reduce(function(obj, str){ obj[str] = str; return obj; }, {});
