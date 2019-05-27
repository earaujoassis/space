# Space Features

Following are described the feature-gates for the whole application. It relies on a memory store
for that and it works on a toggle-based behaviour: when the key-field exists, the feature is available;
when the key-field is not available at the memory store, the feature is not available.

## `user.create`

It turns on the sign-up option throughout the entire application.

## `user.adminify`

It turns on the option to make a given user (post-sign-in) to turn herself an admin,
given the provided application-key for that.

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
