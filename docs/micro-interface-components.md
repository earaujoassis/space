# Micro-interface components

The `web` folder is organized through micro-interface components:

- *Amalthea* is the JavaScript (React + Flux) application responsible for updating a user's
password. It provides interfaces (HTML + CSS + JS) to ask for the resource owner's authorization
over a client request of access to update a user's password.

- *Callisto* is the JavaScript (React + Flux) application responsible for user authorization.
It provides interfaces (HTML + CSS + JS) to ask for the resource owner's authorization over a
client request of access.

- *Europa* is the JavaScript (React + Flux) application responsible for user management. It
provides interfaces (HTML + CSS + JS) to manage user data (profile), client applications
and sessions.

- *Ganymede* is the JavaScript (React + Flux) application responsible for user authentication
(or sign in). There is three steps for user authentication: (1) provide access holder (username
or email); (2) provide password / passphrase; (3) provide a Time-Based One Time Password.

- *Io* is the JavaScript (React + Flux) application responsible for user sign up.

The *`core`* subproject provides reusable JavaScript code (React + Flux) in order to build satellite
applications (like *Io*, *Ganymede*, and *Europa*).

Since the project was originally named *Jupiter* (instead of `space`), the micro-interfaces were
named after Jupiter's moons.

We could also call *micro-interface components* as *microapps* or even *micro-frontends*, however
since they're deployed within the same application and only differs from each other considering
their usage and context (if the user is signed-in etc.), *micro-interfaces components* sounds like
a good definition for what we're doing here.

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
