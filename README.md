# Brighter, Blacker

Brighter, Blacker (BB) is a photo portfolio web application, meant to be used in conjunction with [bbservice](https://github.com/object88/bbservice).

## Architecture

On the client side, BB uses React to manage component rendering, and Relay to manage state & data fetching.

## Technical goals

* Use Google services to authenticate a user via SSO
* Turn this code into a bundle (or bundles) that can be served up from a non-Node / non-JS web server that natively serves over HTTP/2.
* Transform relay's strange object shape into something more friendly towards react proptypes.

## Note
This application requires the schema file from `bbservice` to be copied to `/schema/schema.json`.

#### Front end
* ES6 via [Babel](https://babeljs.io/), [Node](https://nodejs.org/en/)
* [Express](https://expressjs.com/)
* [React](https://facebook.github.io/react/)
* [Relay](https://facebook.github.io/relay/)

#### Quality
* [ESLint](http://eslint.org/)
* [Flow](https://flowtype.org/)
* [Webpack](https://webpack.js.org/)

#### Package management
* [Yarn](https://yarnpkg.com/)
